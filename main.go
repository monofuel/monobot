package monobot

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/monofuel/monobot/handlers"
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
	handlers.SetHandler("interfaces", func(args []string) string {
		interfaces, err := getInterfaces()
		if err != nil {
			return fmt.Sprintf("error getting interfaces: %v", err)
		}
		return fmt.Sprintf("interfaces: %v\n", interfaces)
	})

	startIrcBot()
	startDiscordBot()
	startPushbulletBot()
}

func startPushbulletBot() {
	if Settings.PushbulletAPIKey == "" {
		return
	}
	pushBot, err := connectToPushbullet()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pushBot.CommandCallback = handlers.HandleCommand
	return
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
	if Settings.DiscordToken == "" {
		return
	}

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
	if Settings.BotName == "" {
		info.BotName = fmt.Sprintf("monobot-%s", info.Hostname)
	} else {
		info.BotName = Settings.BotName
	}
	info.Owner = Settings.Owner

	return info, nil
}

func getInterfaces() ([]string, error) {
	var Interfaces []string
	ifaces, err := net.Interfaces()
	if err != nil {
		return Interfaces, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("error listing addresses: %v\n", err)
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			Interfaces = append(Interfaces, ip.String())
		}
	}

	return Interfaces, nil
}
