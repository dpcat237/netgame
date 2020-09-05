package interfaces

import "github.com/netgame/backend/internal/model"

type Notification interface {
	RoundNotification(ntf model.RoundNotification)
}
