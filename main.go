package monobot

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/monofuel/monobot/handlers"
	"github.com/monofuel/monobullet"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Start() {
	fmt.Println("hello world")
	err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s\n", err.Error())
	}

	handlers.SetHandler("info", func(args []string) string {
		info, err := getInfo()
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("Host: %s\nBotName: %s\nOwner: %s\nPID: %d",
			info.Hostname, info.BotName, info.Owner, info.PID)
	})

	startIrcBot()
	startDiscordBot()
	startPushbulletBot()
}

func startPushbulletBot() error {
	if Settings.PushbulletAPIKey == "" {
		return nil
	}
	info, err := getInfo()
	if err != nil {
		return err
	}
	config := &monobullet.Config{
		ApiKey:     Settings.PushbulletAPIKey,
		Realtime:   true,
		DeviceName: info.BotName,
	}

	monobullet.Configuration(config)
	monobullet.Start()
	return nil
}

func startIrcBot() {
	for _, server := range Settings.IrcServers {
		bot, err := connectToIrc(server, []string{"#bots"})
		if err != nil {
			fmt.Printf("Error connecting to irc server: %s\n", err.Error())
		}
		bot.CommandCallback = handlers.HandleCommand

		go bot.Loop()
	}
}

func startDiscordBot() {

	discordBot, err := connectToDiscord()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	discordBot.CommandCallback = handlers.HandleCommand
}

//BotInfo is a struct detailing into about the runtime
type BotInfo struct {
	Hostname string
	PID      int
	BotName  string
	Owner    string
}

func getInfo() (*BotInfo, error) {
	info := new(BotInfo)
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	info.Hostname = hostname
	info.PID = os.Getpid()
	info.BotName = fmt.Sprintf("monobot-%s", info.Hostname)
	info.Owner = Settings.Owner
	return info, nil
}
