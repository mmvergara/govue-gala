package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router  http.Handler
	redisDb *redis.Client
}

func New() *App {
	app := &App{
		redisDb: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_DB_URL"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}),
	}
	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: a.router,
	}

	fmt.Println("Connecting to redis")
	err := a.redisDb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis %w", err)
	}

	defer func() {
		if err := a.redisDb.Close(); err != nil {
			fmt.Println("Failed to close redis", err)
		}
	}()

	fmt.Println("Starting Server")
	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		return server.Shutdown(timeout)
	}

}
