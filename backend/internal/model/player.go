package model

import (
	"errors"
)

type Player struct {
	Name      string `json:"name"`
	NumberOne int8   `json:"number_one"`
	NumberTwo int8   `json:"number_two"`

	lowerBound uint8
	upperBound uint8
	points     int16
}

type playersByLowerBound []Player

func (plrs playersByLowerBound) Len() int           { return len(plrs) }
func (plrs playersByLowerBound) Less(i, j int) bool { return plrs[i].lowerBound < plrs[j].lowerBound }
func (plrs playersByLowerBound) Swap(i, j int)      { plrs[i], plrs[j] = plrs[j], plrs[i] }

type playersByName []Player

func (plrs playersByName) Len() int           { return len(plrs) }
func (plrs playersByName) Less(i, j int) bool { return plrs[i].Name < plrs[j].Name }
func (plrs playersByName) Swap(i, j int)      { plrs[i], plrs[j] = plrs[j], plrs[i] }

type playersByUpperBound []Player

func (plrs playersByUpperBound) Len() int           { return len(plrs) }
func (plrs playersByUpperBound) Less(i, j int) bool { return plrs[i].upperBound < plrs[j].upperBound }
func (plrs playersByUpperBound) Swap(i, j int)      { plrs[i], plrs[j] = plrs[j], plrs[i] }

type PlayerFrontend struct {
	Name   string `json:"name"`
	Points int16  `json:"points"`
}

type playersByPointsDesc []PlayerFrontend

func (plrs playersByPointsDesc) Len() int           { return len(plrs) }
func (plrs playersByPointsDesc) Less(i, j int) bool { return plrs[i].Points > plrs[j].Points }
func (plrs playersByPointsDesc) Swap(i, j int)      { plrs[i], plrs[j] = plrs[j], plrs[i] }

func (plr *Player) Validate() error {
	if plr.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if plr.NumberOne < 1 || plr.NumberTwo < 1 {
		return errors.New("Number can not be lowerBound 1")
	}

	if plr.NumberOne > 10 || plr.NumberTwo > 10 {
		return errors.New("Number can not be upperBound than 10")
	}

	if plr.NumberOne >= plr.NumberTwo {
		plr.upperBound = uint8(plr.NumberOne)
		plr.lowerBound = uint8(plr.NumberTwo)
		return nil
	}

	plr.lowerBound = uint8(plr.NumberOne)
	plr.upperBound = uint8(plr.NumberTwo)
	return nil
}

func (plr Player) countRoundPoints(num uint8) int16 {
	if plr.upperBound == num || plr.lowerBound == num {
		return 5
	}

	if plr.upperBound > num && plr.lowerBound < num {
		return 5 - (int16(plr.upperBound) - int16(plr.lowerBound))
	}

	return -1
}

func (plr *Player) round(num uint8) {
	plr.points += plr.countRoundPoints(num)
}

func (plr Player) toFrontend() PlayerFrontend {
	return PlayerFrontend{
		Name:   plr.Name,
		Points: plr.points,
	}
}
