package interfaces

import (
	"net/http"

	"github.com/netgame/backend/internal/model"
)

type GameController interface {
	Leaderboard(w http.ResponseWriter, r *http.Request)
	Ranking(w http.ResponseWriter, r *http.Request)
	SignUser(w http.ResponseWriter, r *http.Request)
}

type GameHandler interface {
	ListLeaderboard() model.Leaderboard
	ListRanking() model.Ranking
	SignUser(us model.Player) model.Output
}
