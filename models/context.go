package models

import (
	"gobot/error"

	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Client    LanaBot
	Author    *discordgo.User
	ChannelID string
	GuildID   string
}

func (ctx Context) Send(content string) *discordgo.Message {

	message, err := ctx.Client.Session.ChannelMessageSend(ctx.ChannelID, content)
	error.ErrorCheckPanic(err)
	return message
}
