package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/mmvergara/golang-gala/backend/application"
)

func main(){
	godotenv.Load(".env")

	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(),os.Interrupt)
	defer cancel()


	err := app.Start(ctx)

	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
	
}