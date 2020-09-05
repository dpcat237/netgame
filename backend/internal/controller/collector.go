package controller

import (
	"github.com/netgame/backend/internal/handler"
	"github.com/netgame/backend/internal/interfaces"
	"github.com/netgame/backend/internal/logger"
)

type Collector struct {
	Game interfaces.GameController
}

func NewCollector(lgr logger.Logger, hnds handler.Collector) Collector {
	return Collector{
		Game: newGame(hnds.Game, lgr),
	}
}
