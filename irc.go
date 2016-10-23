package main

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

var handlers []func(string)

func connectToIrc(server string, channels []string) error {
	fmt.Println("connecting to ", server)
	bot := irc.IRC("monobot", "monofuel")

	bot.AddCallback("PRIVMSG", func(event *irc.Event) {
		fmt.Println("-----------------------")
		fmt.Println("MSG: ", event.Message)
		fmt.Println("NICK: ", event.Nick)
		fmt.Println("ARGS: ", event.Arguments)
	})

	if err := bot.Connect(server); err != nil {
		return err
	}

	for _, channel := range channels {
		fmt.Println("joining", channel)
		bot.Join(channel)
		bot.Privmsg(channel, "hello world")
	}

	return nil
}
