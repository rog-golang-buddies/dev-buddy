package commandHandlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/brokers/discordBroker"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/brokers/githubBroker"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/models"
	"google.golang.org/appengine/log"
)

func CommandTranslation(command string, brokerContext context.Context) (models.DiscordResponses, error) {
	commandComponents := splitString(command)
	if commandComponents[0] == "db" {
		if commandComponents[1] == "get" {
			if commandComponents[2] == "issues" || commandComponents[2] == "issue" {
				githubClient, err := githubBroker.CreateBroker(brokerContext)
				if err != nil {
					log.Errorf(brokerContext, "the client did not get created.:%v", err)
					return models.DiscordResponses{}, err
				}
				respInitial, err := githubBroker.GetAllIssueNames(brokerContext, githubClient, commandComponents[3])
				if err != nil {
					log.Errorf(brokerContext, "issue not returned.:%v", err)
					return models.DiscordResponses{}, err
				}
				messageReturn := discordgo.Message{
					Content: respInitial,
				}
				messageReturnList := []discordgo.Message{messageReturn}
				resp := models.DiscordResponses{
					Responses: messageReturnList,
				}
				fmt.Print(resp)
				return resp, nil
			}
		} else if commandComponents[1] == "create" {
			if commandComponents[2] == "orginvite" || commandComponents[2] == "orginvites" {
				githubClient, err := githubBroker.CreateBroker(brokerContext)
				if err != nil {
					log.Errorf(brokerContext, "the client did not get created.:%v", err)
					return models.DiscordResponses{}, err
				}
				respInitial, err := githubBroker.CreateOrganizationInvite(brokerContext, githubClient, commandComponents[3])
				if err != nil {
					log.Errorf(brokerContext, "organization invite not created.:%v", err)
					return models.DiscordResponses{}, err
				}
				messageReturn := discordgo.Message{
					Content: respInitial,
				}
				messageReturnList := []discordgo.Message{messageReturn}
				resp := models.DiscordResponses{
					Responses: messageReturnList,
				}
				fmt.Print(resp)
				return resp, nil
			} else if commandComponents[2] == "meeting" {
				resp, err := discordBroker.CreateMeeting(brokerContext, commandComponents[3], commandComponents[4])
				if err != nil {
					log.Errorf(brokerContext, "meeting not created.:%v", err)
					return models.DiscordResponses{}, err
				}

				fmt.Print(resp)
				return resp, nil
			}
		}
	}
	return models.DiscordResponses{}, nil
}

func splitString(command string) []string {
	lowerCaseCommand := strings.ToLower(command)

	words := strings.Fields(lowerCaseCommand)

	return words
}
