package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/bwmarrin/discordgo"
)

var botId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == botId {
		return
	}

	if m.Content == "/flagz" || m.Content == "!flagz" {
		images, err := os.ReadDir("./images")
		rand.Shuffle(len(images), func(i, j int) { images[i], images[j] = images[j], images[i] })
		fmt.Println(images)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

/*	pick a random file from the slice, display it in the chat, add the multiple choice menu, somehow add right and wrong answers,
	let user select and enter ... msg in the chat if it's right or wrong, sleep for a few seconds, remove old image from slice,
	repeat from the start */
