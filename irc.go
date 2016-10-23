package main

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

var handlers []func(string)

//IrcBot wrapper around all the bot work for a single irc server
type IrcBot struct {
	irc.Connection
	Channels []string
	server   string
}

func connectToIrc(server string, channels []string) (*IrcBot, error) {
	fmt.Println("connecting to ", server)
	info, err := getInfo()
	if err != nil {
		return nil, err
	}

	bot := irc.IRC(info.BotName, info.Owner)

	bot.AddCallback("PRIVMSG", func(e *irc.Event) {
		fmt.Println(e)
	})
	bot.AddCallback("001", func(e *irc.Event) {
		for _, channel := range channels {
			fmt.Println("joining", channel)
			bot.Join(channel)
		}
	})
	bot.AddCallback("JOIN", func(e *irc.Event) {
		fmt.Println(e)
		bot.Privmsg(e.Arguments[0], "hello world")
	})

	if err := bot.Connect(server); err != nil {
		return nil, err
	}

	return &IrcBot{*bot, channels, server}, nil
}
