package websocket

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/netgame/backend/internal/logger"
	"github.com/netgame/backend/internal/model"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Manager struct {
	lgr       logger.Logger
	observers map[*websocket.Conn]bool
	wbs       *http.Server
}

func NewManager(lgr logger.Logger, wbsPort string) *Manager {
	mng := &Manager{
		lgr:       lgr,
		observers: make(map[*websocket.Conn]bool),
	}
	rtr := mux.NewRouter()
	rtr.HandleFunc("/subscribe", mng.subscribe)
	mng.wbs = &http.Server{
		Addr:    ":" + wbsPort,
		Handler: rtr,
	}

	return mng
}

func (mng *Manager) RoundNotification(ntf model.RoundNotification) {
	mng.sendToObservers(ntf)
}

func (mng *Manager) Shutdown(ctx context.Context) {
	if err := mng.wbs.Shutdown(ctx); err != nil {
		mng.lgr.Fatal("Failed to shutdown websocket: " + err.Error())
	}
	mng.lgr.Info("Stopped websocket service")
}

func (mng *Manager) Start() {
	go func() {
		if err := mng.wbs.ListenAndServe(); err != nil {
			mng.lgr.Fatal("Failed to start websocket: " + err.Error())
		}
	}()
}

func (mng *Manager) sendToObservers(v interface{}) {
	for obr := range mng.observers {
		msg, err := json.Marshal(v)
		if err != nil {
			mng.lgr.Fatal(err.Error())
			return
		}

		if err := obr.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(mng.observers, obr)
		}
	}
}

func (mng *Manager) subscribe(w http.ResponseWriter, r *http.Request) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		mng.lgr.Warn(err.Error())
		return
	}
	mng.observers[con] = true
}
