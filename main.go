package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	client "github.com/bwmarrin/discordgo"
)

func main() {
	session, err := client.New("Bot " + "BOT TOKEN HERE") // change to your bots token

	if err != nil {
		fmt.Println(err)
		return
	}
	session.AddHandler(message)

	fmt.Print("Bot Is Online")
	defer session.Close()

	if err = session.Open(); err != nil {
		fmt.Println(err)
		return
	}

	scall := make(chan os.Signal, 1)
	signal.Notify(scall, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-scall
}

func message(bot *client.Session, message *client.MessageCreate) {
	if message.Author.Bot { // we don't want our bot to execute it's own commands
		return
	}

	switch { // add more cases as you add commands

	case strings.HasPrefix(message.Content, "&hello"):
		bot.ChannelMessageSend(message.ChannelID, "Well Hello")
	}

}
