package main

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"

	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/config"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/constants"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/discordCommandInterface"
)

func main() {
	// create context
	ctx := context.Background()

	// get all the environment variables
	var configValues config.EnvironmentConfig
	if err := envconfig.Process(ctx, &configValues); err != nil {
		log.Fatal(err)
	}

	// setting the token value to a context
	ctx = context.WithValue(ctx, constants.BotTokenHeader, configValues.DiscordToken)

	// calling the initialize server for discord
	s, err := discordCommandInterface.InitializeDiscordServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := discordCommandInterface.ReadWriteMethod(ctx, s); err != nil {
		log.Fatal(err)
	}
}
