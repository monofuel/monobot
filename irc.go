package main

import (
	"fmt"
	"strings"

	irc "github.com/thoj/go-ircevent"
)

//IrcBot wrapper around all the bot work for a single irc server
type IrcBot struct {
	*irc.Connection
	Channels        []string
	Server          string
	CommandCallback func([]string) string
}

func connectToIrc(server string, channels []string) (*IrcBot, error) {
	fmt.Println("connecting to ", server)
	info, err := getInfo()
	if err != nil {
		return nil, err
	}

	bot := &IrcBot{irc.IRC(info.BotName, info.Owner), channels, server, nil}

	bot.AddCallback("PRIVMSG", func(e *irc.Event) {
		channel := e.Arguments[0]
		fmt.Printf("%s|%s|%s\n", channel, e.Source, e.Message())
		if bot.CommandCallback != nil {
			args := strings.Fields(e.Message())
			if args[0] == Settings.IrcPrefix {
				args = args[1:]
				fmt.Println("running command: ", args)
				response := bot.CommandCallback(args)
				if response != "" {
					lines := strings.Split(response, "\n")
					bot.Privmsg(channel, strings.Join(lines, "\t"))
				}
			}
		}

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

	return bot, nil
}
