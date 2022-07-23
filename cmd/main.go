package main

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"

	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/config"
)

func main() {
	// create context
	ctx := context.Background()

	// get all the environment variables
	var c config.EnvironmentConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

}
