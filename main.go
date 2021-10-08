package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	eventHandlers "kallelganewk/discordbot.go/events"
)

func main() {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	session.AddHandler(eventHandlers.Ready)
	session.AddHandler(eventHandlers.MessageCreate)

	session.Identify.Intents = discordgo.IntentsGuildMessages
	err = session.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("User: " + session.State.User.String())
	fmt.Println("Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	fmt.Println("Exitting...")
	session.Close()
	fmt.Println("Bye!")
}
