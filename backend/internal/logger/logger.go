package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/netgame/backend/config"
)

type Logger interface {
	Fatal(msg string)
	Info(msg string)
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	RequestEnd(act string, startAt time.Time, status *int, errMsg *string)
	Warn(msg string)
}

// Logger wraps zap.Logger
type logger struct {
	client *zap.Logger
}

func (lgr logger) Fatal(msg string) {
	lgr.client.Fatal(msg)
}

func (lgr logger) Info(msg string) {
	lgr.client.Info(msg)
}

func (lgr logger) Infof(template string, args ...interface{}) {
	lgr.client.Sugar().Infof(template, args...)
}

func (lgr logger) Infow(msg string, keysAndValues ...interface{}) {
	lgr.client.Sugar().Infow(msg, keysAndValues...)
}

// New creates a new preconfigured zap.Logger
func New(md config.Mode) (*logger, error) {
	var lgr logger
	var zapLg *zap.Logger
	var err error

	if md.IsProduction() {
		zapLg, err = zap.NewProduction()
	} else {
		zapLg, err = zap.NewDevelopment()
	}
	if err != nil {
		return &lgr, err
	}

	if md.IsProduction() {
		if err := zapLg.Sync(); err != nil {
			return &lgr, err
		}
	}

	lgr.client = zapLg

	return &lgr, nil
}

// RequestEnd log end of request process
func (lgr logger) RequestEnd(act string, startAt time.Time, status *int, errMsg *string) {
	lgr.client.Sugar().Infow("http_request",
		"action", act,
		"status_code", status,
		"error_message", errMsg,
		"created_at", startAt.Unix(),
		"response_time", fmt.Sprintf("%.4f", time.Since(startAt).Seconds()))
}

func (lgr logger) Warn(msg string) {
	lgr.client.Warn(msg)
}
