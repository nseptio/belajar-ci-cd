package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Ditjen-Dikti-Kemdikbud-RI/iss-be/app"
)

func main() {
	application := app.New(*app.NewConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := application.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
