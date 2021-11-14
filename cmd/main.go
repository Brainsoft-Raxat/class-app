package main

import (
	"class-app/pkg/handler"
	"class-app/pkg/service"
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"

	"class-app/pkg/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	e := echo.New()
	databaseUrl := viper.GetString("dbUrl")
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		//logrus.Fatalf("unable to connect to database: %v\n", err)
		e.Logger.Fatalf("unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	repos := repository.NewRepository(dbPool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(e)


	go func() {
		if err := e.Start(viper.GetString("address")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}