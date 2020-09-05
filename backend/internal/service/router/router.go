package router

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"github.com/netgame/backend/internal/logger"
	"github.com/netgame/backend/internal/model"
)

const (
	V1API = "v1"

	requestsPerSecond        = 5
	serviceReadTimeout       = 2 * time.Second
	serviceHandlerTimeout    = 5 * time.Second
	serviceReadHeaderTimeout = 2 * time.Second
	serviceIdleTimeout       = 60 * time.Second
	successfulResponse       = `{ "message": "Success" }`
	timeoutResponse          = `{ "message": "Service temporarily unavailable" }`
)

type Manager struct {
	lgr    logger.Logger
	rtr    *mux.Router
	routes map[string][]model.Route
	srv    *http.Server
}

func NewManager(lgr logger.Logger) Manager {
	return Manager{
		lgr:    lgr,
		routes: make(map[string][]model.Route),
	}
}

func (mng *Manager) AddRoutesGroup(pre string, rts []model.Route) {
	mng.routes[pre] = rts
}

func (mng *Manager) Start(port string) {
	if mng.rtr == nil {
		mng.createRouter()
	}

	// CORS definition
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})

	// Contains:
	// - Logger of each request
	// - Panic recovery
	// - Rate limiter
	// - Timeout for handler
	bsdHnd := mng.panicRecovery(tollbooth.LimitHandler(mng.createRateLimiter(), http.TimeoutHandler(mng.rtr, serviceHandlerTimeout, timeoutResponse)))
	mng.srv = &http.Server{
		Addr:              ":" + port,
		Handler:           mng.logging()(handlers.CORS(headersOk, originsOk, methodsOk)(bsdHnd)),
		ReadTimeout:       serviceReadTimeout,
		ReadHeaderTimeout: serviceReadHeaderTimeout,
		IdleTimeout:       serviceIdleTimeout,
	}

	go func() {
		if err := mng.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mng.lgr.Fatal("Error starting HTTP service: " + err.Error())
		}
	}()
}

func (mng *Manager) Shutdown(ctx context.Context) {
	if err := mng.srv.Shutdown(ctx); err != nil {
		mng.lgr.Warn("Error stopping HTTP router " + err.Error())
	}
}

func (mng *Manager) addRoutes(r *mux.Router, routes []model.Route) {
	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}
}

func (mng Manager) createRateLimiter() *limiter.Limiter {
	return tollbooth.NewLimiter(requestsPerSecond, nil)
}

func (mng *Manager) createRouter() {
	mng.rtr = mux.NewRouter().StrictSlash(true)
	mng.addRoutes(mng.rtr, getSysRoutes(mng.healthCheck))
	mng.rtr.Handle("/debug/vars", http.DefaultServeMux)

	for pre, rts := range mng.routes {
		mng.addRoutes(mng.rtr.PathPrefix("/"+pre).Subrouter(), rts)
	}
}

func (mng Manager) healthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(successfulResponse)); err != nil {
		mng.lgr.Warn("Error returning health check " + err.Error())
	}
}

func (mng Manager) logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				mng.lgr.Infow("HTTP request",
					// Structured context as loosely typed key-value pairs
					"method", r.Method,
					"path", r.URL.Path,
					"remote_addr", r.RemoteAddr,
					"user_agent", r.UserAgent(),
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func (mng Manager) panicRecovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				mng.lgr.Infow("Panic recovered",
					// Structured context as loosely typed key-value pairs
					"method", r.Method,
					"path", r.URL.Path,
					"remote_addr", r.RemoteAddr,
					"user_agent", r.UserAgent(),
					"error", err,
				)

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
