package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gobot/commands"
	"gobot/error"
	"gobot/models"
	"gobot/utils/fs"

	"github.com/bwmarrin/discordgo"
)

func main() {
	const prefix = "!"
	var owners = make([]string, 0)
	var token = fs.ReadFileWhole("token.txt")
	var commandList = make(map[string]models.Command)

	session, err := discordgo.New("Bot " + token)
	error.ErrorCheckPanic(err)

	var Bot = models.LanaBot{Session: session, Token: token, Owners: owners, Prefix: prefix, Commands: commandList}
	Bot.AddCommands(commands.AllCommands)
	session.Identify.Intents = models.Intents

	session.AddHandler(Bot.ProcessMessage)

	fmt.Println("Starting bot...")

	err = session.Open()
	error.ErrorCheckPanic(err)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	fmt.Println("Bot closing...")
	session.Close()
	fmt.Println("Bot closed.")

}
