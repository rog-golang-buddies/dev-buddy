package main

import (
	"log"

	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/discordCommandInterface"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/utils"
)

func main() {
	ctx, err := utils.SetContext()
	if err != nil {
		log.Fatal(err)
	}
	// test
	// githubClient, _ := githubBroker.CreateBroker(ctx)
	// fmt.Print(githubBroker.CreateOrganizationInvite(ctx, githubClient, "SupornoChaudhury"))

	// calling the initialize server for discord
	s, err := discordCommandInterface.InitializeDiscordServer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := discordCommandInterface.StartListening(s); err != nil {
		log.Fatal(err)
	}
}
