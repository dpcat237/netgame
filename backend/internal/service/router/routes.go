package router

import (
	"net/http"

	"github.com/netgame/backend/internal/controller"
	"github.com/netgame/backend/internal/model"
)

// GetV1Routes sets version API version 1 routes
func GetV1Routes(ctrs controller.Collector) []model.Route {
	return []model.Route{
		newRoute(http.MethodOptions, "/game/sign", meta),
		newRoute(http.MethodPost, "/game/sign", ctrs.Game.SignUser),
		newRoute(http.MethodGet, "/game/leaderboard", ctrs.Game.Leaderboard),
		newRoute(http.MethodGet, "/game/ranking", ctrs.Game.Ranking),
	}
}

func getSysRoutes(hnd http.HandlerFunc) []model.Route {
	return []model.Route{
		newRoute(http.MethodGet, "/services/health", hnd),
	}
}

func newRoute(m, p string, h http.HandlerFunc) model.Route {
	return model.Route{Method: m, Pattern: p, HandlerFunc: h}
}

func meta(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}
