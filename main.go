package main

import (
	"disrabot/utils"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!build") {
		builds := utils.LoadBuilds()
		randomNum := utils.GetRandomUpTo(150)
		fmt.Println(m.MessageReference)

		go func() {
			_, err := s.ChannelMessageSend(m.ChannelID, builds.GetBuild(randomNum))
			if err != nil {
				fmt.Println("Cant send message")
			}
		}()
	} else if strings.HasPrefix(m.Content, "!map") {
		randomNum := utils.GetRandomUpTo(6)
		go func() {
			_, err := s.ChannelMessageSend(m.ChannelID, utils.GetMap(randomNum))
			if err != nil {
				fmt.Println("Cant get map")
			}
		}()
	} else {
		//
	}
}
