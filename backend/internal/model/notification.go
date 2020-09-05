package model

type RoundNotification struct {
	Finished           bool             `json:"finished"`
	FinishedPlayground bool             `json:"finished_playground"`
	Players            []PlayerFrontend `json:"players"`
	Round              uint8            `json:"round"`
	Winner             PlayerFrontend   `json:"winner"`
}
