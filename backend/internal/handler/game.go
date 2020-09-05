package handler

import (
	"net/http"
	"time"

	"github.com/netgame/backend/internal/interfaces"
	"github.com/netgame/backend/internal/model"
)

const (
	maxGames  = 10
	pauseTime = time.Second * 10
	roundTime = time.Second
)

type Game struct {
	leaderboard model.Leaderboard
	notify      interfaces.Notification
	playground  model.Game
	ranking     model.Ranking
}

func newGame(notify interfaces.Notification) *Game {
	return &Game{
		notify:     notify,
		playground: model.NewGame(),
	}
}

func (hnd *Game) ListLeaderboard() model.Leaderboard {
	return hnd.leaderboard
}

func (hnd *Game) ListRanking() model.Ranking {
	return hnd.ranking
}

func (hnd *Game) SignUser(us model.Player) model.Output {
	if !hnd.playground.IsUserNameAllowed(us.Name) {
		return model.ErrorString("This name is taken. Change name and try again", http.StatusConflict)
	}

	hnd.playground.AddUser(us)
	if hnd.playground.IsReady() {
		go hnd.playGame()
	}

	return model.ErrorNil()
}

func (hnd *Game) playGame() {
	hnd.playground.Start()
	var cg uint8
	var maxGms bool

	for {
		hnd.playground.Round()
		hnd.notify.RoundNotification(hnd.playground.ToRoundNotification())
		time.Sleep(roundTime)
		if !hnd.playground.IsFinished() {
			continue
		}

		cg++
		if cg == maxGames {
			maxGms = true
		}
		hnd.playground.SearchCandidates()
		hnd.playground.DefineWinner()

		hnd.processResult(hnd.playground, maxGms)
		hnd.playground.Reset()
		if maxGms {
			break
		}

		time.Sleep(pauseTime)
		hnd.playground.Start()
	}
}

func (hnd *Game) processResult(gm model.Game, gmsEnd bool) {
	hnd.leaderboard.AddPlayer(gm.GetWinner())
	hnd.ranking.AddGame(gm)
	ntf := gm.ToResultsNotification()
	ntf.FinishedPlayground = gmsEnd
	hnd.notify.RoundNotification(ntf)
}
