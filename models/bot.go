package models

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	Intents = discordgo.IntentsDirectMessages |
		discordgo.IntentsGuildBans |
		discordgo.IntentsGuildEmojis |
		discordgo.IntentsGuildIntegrations |
		discordgo.IntentsGuildInvites |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuildMessageReactions |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildPresences |
		discordgo.IntentsGuildVoiceStates |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildVoiceStates
)

type LanaBot struct {
	Session  *discordgo.Session
	Token    string
	Owners   []int
	Prefix   string
	Commands map[string]Command
}

func (bot LanaBot) ProcessMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	fmt.Println("Incoming message")

	if strings.HasPrefix(message.Content, bot.Prefix) {
		var listOfWords = strings.Fields(message.Content)
		fmt.Println(listOfWords)
		var possibleCommandString, ok = strings.CutPrefix(listOfWords[0], bot.Prefix)
		fmt.Println(possibleCommandString)
		if ok {
			command, exists := bot.Commands[possibleCommandString]
			if exists {
				var _, argsraw, _ = strings.Cut(message.Content, possibleCommandString)
				var args = strings.Fields(argsraw) // Create the context
				ctx := &Context{bot, message.Author, message.ChannelID, message.GuildID}

				// Execute all the checks
				for i := 0; i < len(command.Checks); i++ {
					err := command.Checks[i](ctx)
					if err != nil {
						ctx.Send(err.Error())
						fmt.Println("Check not passed")
						return
					}
				}

				command.Callback(ctx, args)
			}
		}
	}
}

func (bot LanaBot) AddCommands(commands map[string]Command) {
	for name, command := range commands {
		fmt.Println("Adding command: " + name)
		bot.Commands[name] = command
	}
}

//
//func (bot LanaBot) AddIntents(session *discordgo.Session) {
//	session.Identify.Intents = discordgo.IntentsGuildMessages
//}
