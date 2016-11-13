package main

import "github.com/monofuel/monobot"

func main() {
	c := make(chan int)
	monobot.Start()
	<-c //never quit (for now)
}
