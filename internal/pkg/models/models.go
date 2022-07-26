package models

import "github.com/bwmarrin/discordgo"

type DiscordResponses struct {
	Responses []discordgo.Message
}

type DiscordReactions struct {
	Reactions []discordgo.MessageReaction
}

type DiscordEmojis struct {
	Emojis []discordgo.Emoji
}
