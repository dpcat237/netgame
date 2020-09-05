package model

import (
	"math/rand"
	"sort"
	"time"
)

const (
	gameStatusBeginning = GameStatus(iota)
	gameStatusActive
	gameStatusFinished
)

const (
	minRandom              = 1
	maxRandom              = 10
	maxRounds              = 30
	winnerPoints           = 21
	requiredPlayersToStart = 2
)

type GameStatus uint8

type Game struct {
	players    []Player
	waiting    []Player
	candidates []Player
	winner     Player
	names      map[string]bool
	status     GameStatus
	round      uint8
}

type GameResults struct {
	Number  uint32           `json:"number"`
	Players []PlayerFrontend `json:"players"`
}

func NewGame() Game {
	rand.Seed(time.Now().UTC().UnixNano())

	return Game{
		names: make(map[string]bool),
	}
}

func (gm *Game) AddUser(us Player) {
	if gm.status == gameStatusBeginning || gm.status == gameStatusFinished {
		gm.players = append(gm.players, us)
		gm.names[us.Name] = true
		return
	}
	gm.waiting = append(gm.waiting, us)
	gm.names[us.Name] = true
}

func (gm *Game) DefineWinner() {
	if len(gm.candidates) == 1 {
		gm.winner = gm.candidates[0]
		return
	}

	lastK := len(gm.candidates) - 1
	penultK := len(gm.candidates) - 2

	sort.Sort(playersByUpperBound(gm.candidates))
	if gm.candidates[lastK].upperBound > gm.candidates[penultK].upperBound {
		gm.winner = gm.candidates[lastK]
		return
	}

	sort.Sort(playersByLowerBound(gm.candidates))
	if gm.candidates[lastK].lowerBound > gm.candidates[penultK].lowerBound {
		gm.winner = gm.candidates[lastK]
		return
	}

	sort.Sort(playersByName(gm.candidates))
	gm.winner = gm.candidates[0]
}

func (gm Game) GetWinner() Player {
	return gm.winner
}

func (gm Game) IsReady() bool {
	return gm.status == gameStatusBeginning && len(gm.players) >= requiredPlayersToStart
}

func (gm Game) IsUserNameAllowed(nm string) bool {
	if _, ok := gm.names[nm]; ok {
		return false
	}
	return true
}

func (gm Game) IsFinished() bool {
	return gm.round == maxRounds || len(gm.candidates) > 0
}

func (gm *Game) Reset() {
	gm.status = gameStatusFinished
	gm.candidates = []Player{}
	gm.winner = Player{}
	gm.round = 0
	for k := range gm.players {
		gm.players[k].points = 0
	}
}

func (gm *Game) Round() {
	num := gm.randNumber()
	for k := range gm.players {
		gm.players[k].round(num)
		if gm.players[k].points == winnerPoints {
			gm.candidates = append(gm.candidates, gm.players[k])
		}
	}
	gm.round++
}

func (gm *Game) SearchCandidates() {
	if len(gm.candidates) > 0 {
		return
	}

	hgScr := int16(-1 << 10)
	var bffPlrs []Player
	for _, plr := range gm.players {
		if plr.points > hgScr {
			bffPlrs = []Player{}
			bffPlrs = append(bffPlrs, plr)
			hgScr = plr.points
			continue
		}
		if plr.points == hgScr {
			bffPlrs = append(bffPlrs, plr)
		}
	}
	gm.candidates = append(gm.candidates, bffPlrs...)
}

func (gm *Game) Start() {
	gm.status = gameStatusActive
	if len(gm.waiting) > 0 {
		gm.players = append(gm.players, gm.waiting...)
		gm.waiting = []Player{}
	}
}

func (gm Game) ToResults() GameResults {
	var plrs []PlayerFrontend
	for _, plr := range gm.players {
		plrs = append(plrs, plr.toFrontend())
	}
	return GameResults{
		Players: plrs,
	}
}

func (gm Game) ToResultsNotification() RoundNotification {
	return RoundNotification{
		Finished: true,
		Round:    gm.round,
		Winner:   gm.winner.toFrontend(),
	}
}

func (gm Game) ToRoundNotification() RoundNotification {
	var ntf RoundNotification
	ntf.Round = gm.round
	for _, plr := range gm.players {
		ntf.Players = append(ntf.Players, plr.toFrontend())
	}
	return ntf
}

func (gm Game) randNumber() uint8 {
	return minRandom + uint8(rand.Intn(maxRandom-minRandom))
}
