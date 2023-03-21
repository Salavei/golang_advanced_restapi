package main

import (
	"github.com/Salavei/golang_advanced_restapi/internal/user"
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {

	logging.Init()
	logger := logging.GetLogger()

	logger.Info("create router")

	router := httprouter.New()
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening 0.0.0.0:3000")
	logger.Fatal(server.Serve(listener))
}
