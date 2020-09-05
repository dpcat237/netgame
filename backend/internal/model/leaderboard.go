package model

import "sort"

type Leaderboard struct {
	Winners []PlayerFrontend `json:"winners"`
}

func (lds *Leaderboard) AddPlayer(plr Player) {
	lds.Winners = append(lds.Winners, plr.toFrontend())
	sort.Sort(playersByPointsDesc(lds.Winners))
}
