package monobot

import (
	"context"
	"fmt"
	"strings"

	"log"

	"github.com/monofuel/monobullet"
)

type PushbulletBot struct {
	Device          *monobullet.Device
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
		Debug:      false,
	}
	monobullet.Configuration(pushConfig)
	device, err := monobullet.AddOwnDevice()
	if err != nil {
		log.Fatal(err)
	}
	bot.Device = device
	go func(bot *PushbulletBot) {
		monobullet.Start(bot.ctx)
	}(bot)
	go func() {
		for {
			note := <-monobullet.PushChannel
			if (note.TargetDeviceIden == bot.Device.Iden ||
				note.TargetDeviceIden == "") &&
				bot.CommandCallback != nil &&
				note.SourceDeviceIden != bot.Device.Iden {
				fmt.Printf("note was to bot, running command\n")
				args := strings.Split(note.Body, " ")
				resp := bot.CommandCallback(args)
				if resp == "no such command" {
					continue
				}
				if resp != "" {
					monobullet.SendPush(&monobullet.Push{
						Type:             "note",
						Title:            args[0],
						Body:             resp,
						TargetDeviceIden: note.SourceDeviceIden,
						SourceDeviceIden: bot.Device.Iden,
					})
				}
			}
		}
	}()
	/*_, err = monobullet.SendPush(&monobullet.Push{
		Type:  "note",
		Title: fmt.Sprintf("%v started", info.BotName),
	})
	if err != nil {
		log.Fatal(err)
	}*/
	return bot, nil
}
