package handlers

import "math/rand"

var randomLinks = []string{
	"https://youtu.be/-w-58hQ9dLk",
	"https://youtu.be/w4WrnNRkYE8",
}

func randLink(args []string) string {
	return randomLinks[rand.Intn(len(randomLinks))]
}
