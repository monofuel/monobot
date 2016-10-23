package main

import "fmt"

func init() {

}

func main() {
	c := make(chan int)
	fmt.Println("hello world")
	err := connectToIrc("japura.net:6667", []string{"#monobot"})
	if err != nil {
		fmt.Printf("Error connecting to irc server: %s\n", err.Error())
	}
	<-c
}
