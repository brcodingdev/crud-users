package main

import (
	"fmt"
	"github.com/brcodingdev/go-crud-users/internal/adapters/repository"
	"github.com/brcodingdev/go-crud-users/internal/core/service"
	"github.com/brcodingdev/go-crud-users/internal/ports/api"
	"github.com/brcodingdev/go-crud-users/internal/server"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"
)

// @title GO CRUD Users API Docs
// @version 1.0.0
// @contact.name Cleberson Henrique
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8010
// @BasePath /
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(
		logger,
		"ts",
		log.DefaultTimestampUTC,
		"loc",
		log.DefaultCaller,
	)

	err := level.Info(logger).Log("users service starting...")
	if err != nil {
		return err
	}

	// setup database and migrate model
	db, err := repository.Connect()

	if err != nil {
		return err
	}

	// migrate DB with table models
	err = repository.MigrateDB()

	if err != nil {
		return err
	}

	level.Info(logger).Log("Database", "migrated")

	userRepository := repository.NewUserUserRepository(db)
	userService := service.NewUserService(userRepository)
	r := api.RegisterRoute()
	userAPI := api.NewUserAPI(r, userService)
	userAPI.RegisterRoutes()

	port := os.Getenv("PORT")
	appServer := server.NewServer(port, r)

	// handle graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		fmt.Println("received interrupt signal. shutting down gracefully...")
		// close the database connection
		db, err := db.DB()
		if err == nil {
			db.Close()
		}

		fmt.Println("graceful shutdown completed")
		os.Exit(0)
	}()

	// start server
	level.Info(logger).Log("Server", "starting", "port", port)
	defer appServer.Shutdown()
	err = appServer.Serve()
	return err
}
