package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	botId string
	s     *discordgo.Session
)

func Start() {
	s, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := s.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	botId = u.ID

	err = s.Open()

	s.AddHandler(messageHandler)

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

	if m.Content == "/flagz" {
		initialRandomNum := pickRandomNr()
		fmt.Println(initialRandomNum)
	}
}

/* TODO?

1) Get a copy of the complete, original map from the cache
2) If the initialRandomNum is e.g. 197, pick the respective value (in this case: images\z\zimbabwe.png) and post this image in the chat (os.Open? os.Read? How to do it with Discordgo?)
3) Somehow create a select menu fitting to this posted national flag --> A: South Africa, B: Zimbabwe, C: Iran, D: USA (for example).
   Where to store these options for 197 flags and how to include them each time? Because for other flags (e.g. 1:images\a\afghanistan.png), I want other select menu options.
4) Inform the player if their answer was wrong or right. Or if better, just post the right answer without checking via the select menu.
5) Delete the already used image from the map and sleep a few seconds (so the user can read the result)
6) Pick a new random number (... create another function: pick random number from current len() of the map copy)
7) Edit the message, show new image, select menu etc.
8) Repeat until map is empty

*/
