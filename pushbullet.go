package monobot

import (
	"context"
	"fmt"

	"log"

	"github.com/monofuel/monobullet"
)

type PushbulletBot struct {
	ctx             context.Context
	CommandCallback func([]string) string
}

func connectToPushbullet() (*PushbulletBot, error) {
	bot := &PushbulletBot{
		ctx: context.Background(),
	}
	info, err := getInfo()
	if err != nil {
		log.Fatal(err)
	}
	pushConfig := &monobullet.Config{
		APIKey:     Settings.PushbulletAPIKey,
		DeviceName: info.BotName,
		Debug:      true,
	}
	monobullet.Configuration(pushConfig)
	go func(bot *PushbulletBot) {
		monobullet.Start(bot.ctx)
		go func() {
			for {
				note := <-monobullet.PushChannel
				if note.Direction == "self" {
					fmt.Printf("got note %v\n", note)
				}
			}
		}()
	}(bot)
	_, err = monobullet.SendPush(&monobullet.Push{
		Type:  "note",
		Title: fmt.Sprintf("%v started", info.BotName),
	})
	if err != nil {
		log.Fatal(err)
	}
	return bot, nil
}
