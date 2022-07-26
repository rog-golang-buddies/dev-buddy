package utils

import (
	"context"
	"log"

	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/config"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/constants"
	"github.com/sethvargo/go-envconfig"
)

func SetContext() (context.Context, error) {
	// create context
	ctx := context.Background()

	// get all the environment variables
	var configValues config.EnvironmentConfig
	if err := envconfig.Process(ctx, &configValues); err != nil {
		return ctx, err
	}

	log.Print(configValues.GithubPAT)
	log.Print(configValues.DiscordToken)

	// setting the discord token value to context
	ctx = context.WithValue(ctx, constants.BotTokenHeader, configValues.DiscordToken)

	// setting the personal access token value to context
	ctx = context.WithValue(ctx, constants.GHPATHeader, configValues.GithubPAT)

	// setting the name of the organization
	ctx = context.WithValue(ctx, constants.OwnerKey, configValues.OwnerName)

	return ctx, nil
}
