package controller

import (
	"net/http"
	"time"

	"github.com/netgame/backend/internal/interfaces"
	"github.com/netgame/backend/internal/logger"
	"github.com/netgame/backend/internal/model"
)

type Game struct {
	gmHnd interfaces.GameHandler
	lgr   logger.Logger
}

func newGame(gmHnd interfaces.GameHandler, lgr logger.Logger) *Game {
	return &Game{gmHnd: gmHnd, lgr: lgr}
}

func (ctr *Game) Leaderboard(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	errMsg := ""
	defer ctr.lgr.RequestEnd("game.leaderboard", time.Now(), &status, &errMsg)

	returnJson(w, ctr.gmHnd.ListLeaderboard())
}

func (ctr *Game) Ranking(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	errMsg := ""
	defer ctr.lgr.RequestEnd("game.ranking", time.Now(), &status, &errMsg)

	returnJson(w, ctr.gmHnd.ListRanking())
}

func (ctr *Game) SignUser(w http.ResponseWriter, r *http.Request) {
	status := http.StatusAccepted
	errMsg := ""
	defer ctr.lgr.RequestEnd("game.sign_user", time.Now(), &status, &errMsg)

	var us model.Player
	if err := getBodyContent(r, &us); err != nil {
		status = http.StatusPreconditionFailed
		errMsg = err.Error()
		returnError(w, model.ErrorString(errMsg, status))
		return
	}

	if err := us.Validate(); err != nil {
		status = http.StatusPreconditionFailed
		errMsg = err.Error()
		returnError(w, model.ErrorString(errMsg, status))
		return
	}

	if out := ctr.gmHnd.SignUser(us); out.IsError() {
		status = out.Status
		errMsg = out.Message
		returnError(w, out)
		return
	}
	w.WriteHeader(status)
}
