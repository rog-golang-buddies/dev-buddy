package discordCommandInterface

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/commandHandlers"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/constants"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/utils"
)

func InitializeDiscordServer(c context.Context) (*discordgo.Session, error) {
	botToken := c.Value(constants.BotTokenHeader)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + fmt.Sprint(botToken))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
		return nil, err
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageRead)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	return dg, nil
}

func StartListening(dg *discordgo.Session) error {
	// Open a websocket connection to Discord and begin listening.
	err := dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
		return err
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

	return nil
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageRead(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	brokerContext, err := utils.SetContext()
	if err != nil {
		log.Fatal(err)
	}
	response, err := commandHandlers.CommandTranslation(m.Content, brokerContext)
	if err != nil {
		log.Fatal(err)
	}
	for _, resp := range response.Responses {
		sentMessage, _ := s.ChannelMessageSend(m.ChannelID, resp.Content)
		if len(resp.Reactions) > 0 {
			for _, reaction := range resp.Reactions {
				fmt.Print(reaction.Emoji.Name)
				err = s.MessageReactionAdd(m.ChannelID, sentMessage.ID, reaction.Emoji.Name)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	// If the message is "ping" reply with "Pong!"
	// if m.Content == "ping" {
	// 	s.ChannelMessageSend(m.ChannelID, "Pong!")
	// }

	// If the message is "pong" reply with "Ping!"
	// if m.Content == "pong" {
	// 	s.ChannelMessageSend(m.ChannelID, "Ping!")
	// }
}
