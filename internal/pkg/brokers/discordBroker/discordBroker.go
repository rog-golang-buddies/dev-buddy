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
	startDateTime, err := time.Parse("2006-01-02", startDate) //YYYY-MM-DD
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
	initialMessage := discordgo.Message{
		Content: "Please react with your availability during the following days. Try to react with as many emojis as possible.",
	}
	responseMessages = append(responseMessages, initialMessage)
	for i := 0; i < daysInt; i++ {
		displayDateTime := startDateTime.AddDate(0, 0, i)
		dayName := displayDateTime.Weekday()

		formattedDisplayDate := displayDateTime.Format("2006-01-02")

		firstReactions := createFirstHalfReactionsList()
		messageReturnFirstHalf := discordgo.Message{
			Content: "First half of " + dayName.String() + "\ndate(YYYY-MM-DD): " + formattedDisplayDate + "\nAll times are in GMT",
			// TODO: start from here
			Reactions: firstReactions,
		}

		secondReactions := createSecondHalfReactionsList()
		messageReturnSecondHalf := discordgo.Message{
			Content:   "Second half of " + dayName.String() + "\ndate(YYYY-MM-DD): " + formattedDisplayDate + "\nAll times are in GMT",
			Reactions: secondReactions,
		}

		responseMessages = append(responseMessages, messageReturnFirstHalf)
		responseMessages = append(responseMessages, messageReturnSecondHalf)
	}
	response := models.DiscordResponses{
		Responses: responseMessages,
	}
	return response, nil
}

func createFirstHalfReactionsList() []*discordgo.MessageReactions {
	emojiList := []string{":12to1am:1001391717376860211", ":1to2am:1001391662674739271", ":2to3am:1001391667061989376", ":3to4am:1001391671914811422", ":4to5am:1001391676180414644", ":5to6am:1001391680747995246",
		":6to7am:1001391686104133633", ":7to8am:1001391690822729739", "8to9am:1001391696279502859", ":9to10am:1001391700972929034",
		":10to11am:1001391706522009611", ":11to12noon:1001391714826735686"}
	var response []*discordgo.MessageReactions
	for _, emojiName := range emojiList {
		discordEmoji := discordgo.Emoji{
			Name: emojiName,
		}

		discordReaction := discordgo.MessageReactions{
			Emoji: &discordEmoji,
		}

		response = append(response, &discordReaction)
	}
	return response
}

func createSecondHalfReactionsList() []*discordgo.MessageReactions {
	emojiList := []string{":1to2pm:1001391664805462046", ":2to3pm:1001391669574373386", ":3to4pm:1001391673982599189",
		":4to5pm:1001391678688600144", ":5to6pm:1001391683054862386", ":6to7pm:1001391688641675364", ":7to8pm:1001391693322530847",
		":8to9pm:1001391698594762774", ":9to10pm:1001391703632125952", ":10to11pm:1001391709130862694", ":11to12mid:1001391712108822549"}
	var response []*discordgo.MessageReactions
	for _, emojiName := range emojiList {
		discordEmoji := discordgo.Emoji{
			Name: emojiName,
		}

		discordReaction := discordgo.MessageReactions{
			Emoji: &discordEmoji,
		}

		response = append(response, &discordReaction)
	}
	return response
}
