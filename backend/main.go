package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/netgame/backend/config"
	"github.com/netgame/backend/internal/controller"
	"github.com/netgame/backend/internal/handler"
	"github.com/netgame/backend/internal/logger"
	"github.com/netgame/backend/internal/service/router"
	"github.com/netgame/backend/internal/service/websocket"
)

const shutdownTimeout = 5 * time.Second

func main() {
	cfg := config.LoadData()
	lgr, err := logger.New(cfg.Mode)
	if err != nil {
		panic("Error initializing logger: " + err.Error())
	}

	wscMng := websocket.NewManager(lgr, cfg.WebsocketPort)

	// Init handlers
	hndsCll := handler.Init(wscMng)
	// Init router controllers
	ctrs := controller.NewCollector(lgr, hndsCll)

	// Create router manager
	rtrMng := router.NewManager(lgr)
	// Add routes
	rtrMng.AddRoutesGroup(router.V1API, router.GetV1Routes(ctrs))

	wscMng.Start()
	lgr.Infof("Websocket started on port %s", cfg.WebsocketPort)

	rtrMng.Start(cfg.HTTPport)
	lgr.Infof("HTTP r	outer started on port %s", cfg.HTTPport)

	gracefulShutdown(lgr, rtrMng, wscMng)
}

// gracefulShutdown stops router after receiving system notification
func gracefulShutdown(lgr logger.Logger, rtrMng router.Manager, wscMng *websocket.Manager) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	close(c)

	timeout, cancelFunc := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelFunc()

	lgr.Info("Service stopping")
	wscMng.Shutdown(timeout)
	rtrMng.Shutdown(timeout)
	lgr.Info("Service stopped")
}
