package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

//DiscordBot wrapper around all the bot work for a single discord server
type DiscordBot struct {
	*discordgo.Session
	CommandCallback func([]string) string
}

func connectToDiscord() (*DiscordBot, error) {
	/*info, err := getInfo()
	if err != nil {
		return nil, err
	}*/
	if Settings.DiscordToken != "" {
		fmt.Println("Picking up discord token")
	} else {
		return nil, fmt.Errorf("discord not configured")
	}
	discord, err := discordgo.New("", "", fmt.Sprintf("Bot %s", Settings.DiscordToken))
	if err != nil {
		return nil, err
	}
	bot := &DiscordBot{discord, nil}

	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

		if bot.CommandCallback != nil {
			args := strings.Fields(m.Content)
			fmt.Println(args)
			if args[0] == Settings.IrcPrefix {
				args = args[1:]
				fmt.Println("running command: ", args)
				response := bot.CommandCallback(args)
				if response != "" {
					fmt.Println("responding: ", response)
					_, err := s.ChannelMessageSend(m.ChannelID, response)
					if err != nil {
						fmt.Println("error: ", err.Error())
					}
				} else {
					fmt.Println("not responding")
				}
			}
		}
	})

	err = bot.Open()
	if err != nil {
		return nil, err
	}

	return bot, nil
}
