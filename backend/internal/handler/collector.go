package handler

import "github.com/netgame/backend/internal/interfaces"

type Collector struct {
	Game interfaces.GameHandler
}

func Init(ntf interfaces.Notification) Collector {
	return Collector{
		Game: newGame(ntf),
	}
}
