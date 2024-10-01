package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	router      *echo.Echo
	mongoClient *mongo.Client
	mongoDB     *mongo.Database
	config      Config
}

func New(config Config) *App {
	fmtDatabaseUri := "mongodb://%s:%s@%s:%d/%s"
	databaseUri := fmt.Sprintf(fmtDatabaseUri, config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	app := &App{
		mongoClient: client,
		mongoDB:     client.Database(config.DB.DBName),
		config:      config,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", a.config.Server.Port),
		Handler:           a.router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := a.mongoClient.Ping(pingCtx, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	defer func() {
		if err = a.mongoClient.Disconnect(ctx); err != nil {
			log.Panicf("failed to disconnect from MongoDB: %v", err)
		}
		log.Println("Close MongoDB connection")
	}()

	log.Printf("Starting server on port %d", a.config.Server.Port)

	// Graceful Shutdown
	ch := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
