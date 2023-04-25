package main

import (
	"context"
	"fmt"
	author2 "github.com/Salavei/golang_advanced_restapi/internal/author"
	author "github.com/Salavei/golang_advanced_restapi/internal/author/db"
	book2 "github.com/Salavei/golang_advanced_restapi/internal/book"
	book "github.com/Salavei/golang_advanced_restapi/internal/book/db"
	"github.com/Salavei/golang_advanced_restapi/internal/config"
	"github.com/Salavei/golang_advanced_restapi/internal/user"
	"github.com/Salavei/golang_advanced_restapi/pkg/client/postgresql"
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {

	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.Background(), cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	repository := author.NewRepository(postgreSQLClient, logger)

	handler := author2.NewHandler(repository, logger)
	handler.Register(router)

	repository1 := book.NewRepository(postgreSQLClient, logger)

	handler1 := book2.NewHandler(repository1, logger)
	handler1.Register(router)

	//repository := author.NewRepository(postgreSQLClient, logger)
	//author1 := author.Author{Name: "Andrew"}
	//_ = storage.Create(context.Background(), &author1)

	//_ = repository.Delete(context.Background(), "Andrew")

	//cfgMongo := cfg.MongoDB
	//NewClient, err := mongodb.NewClient(context.Background(),
	//cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)

	handler = user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {

		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
