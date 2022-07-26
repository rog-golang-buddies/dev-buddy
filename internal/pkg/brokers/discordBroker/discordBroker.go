package discordBroker

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rog-golang-buddies/dev-buddy/internal/pkg/models"
)

func CreateMeeting(ctx context.Context, startDate string, endDate string) (models.DiscordResponses, error) {
	startDateTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return models.DiscordResponses{}, fmt.Errorf("failed converting startDate: %v", err)
	}

	endDateTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return models.DiscordResponses{}, fmt.Errorf("failed converting endDate: %v", err)
	}

	days := endDateTime.Sub(startDateTime).Hours() / 24
	daysInt := int(math.Round(days))
	var responseMessages []discordgo.Message
	for i := 0; i < daysInt; i++ {
		displayDateTime := startDateTime.AddDate(0, 0, i)
		dayName := displayDateTime.Weekday()
		fmt.Println(dayName)

		messageReturnFirstHalf := discordgo.Message{
			Content: "Please fill in your available time slots for first half of" + dayName.String() + ", date: " + displayDateTime.String(),
			// TODO: start from here
			//Reactions: &[]discordgo.MessageReaction{},
		}
		messageReturnSecondHalf := discordgo.Message{
			Content: "Please fill in your available time slots for second half of" + dayName.String() + ", date: " + displayDateTime.String(),
		}

		responseMessages[2*i] = messageReturnFirstHalf
		secondIndex := (2 * i) + 1
		responseMessages[secondIndex] = messageReturnSecondHalf
	}
	response := models.DiscordResponses{
		Responses: responseMessages,
	}
	return response, nil
}
