package main

import (
	"./Utility"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)


var index int
var token string
var channels []string
var words [][2]string

func main() {

	token, index = Utility.ReadJson()
	words = Utility.GetCsv()
	fmt.Println("csv length",len(words))
	fmt.Println(token ,"  ", index)
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(guildCreate)
	dg.AddHandler(guildDelete)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.")
	waitForTime(dg)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return;
	}
	if m.Content == "++play" {
		sendWords(s)
	}
	return
	//if Utility.StartsWith(m.Content,"++") {
	//}

}

func guildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	channels = append(channels,g.SystemChannelID)
}

func guildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	channels = Utility.Remove(channels,Utility.GetIndex(g.SystemChannelID,channels))
}



func waitForTime( s *discordgo.Session) {
	for {
		time.Sleep(time.Minute)
		if time.Now().Hour() == 8 && time.Now().Minute() == 0 {
			sendWords(s)
			time.Sleep(time.Minute + time.Second)
		}
	}
}

func sendWords(s *discordgo.Session) {

	if len(words) == 0 {
		for x := 0; x< len(channels);x++ {
			_, err := s.ChannelMessageSend(channels[x], "the word list is empty lol")
			if err != nil {
				panic(err)
			}
		}
	}

	if index >= len(words) {
		index = 0
		fmt.Println("index has passed words length")
	}


	for x := 0; x < len(channels); x++ {
		_, err := s.ChannelMessageSend(channels[x], words[index][0])
		if err != nil {
			panic(err)
		}
		_, err = s.ChannelMessageSend(channels[x],words[index][1])
		if err != nil {
			panic(err)
		}
	}
	index += 1
	Utility.UpdateIndex(index,token)
}



