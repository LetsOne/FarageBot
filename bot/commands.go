package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/PuerkitoBio/goquery"
)


func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Content) <= 0 || m.Content[0] != '!'  {
		return
	}

	u := m.Author
	s.ChannelMessageSend("203630579617366016", (u.Username + " sent " + m.Content))
	
	msg := strings.Replace(m.ContentWithMentionsReplaced(), s.State.Ready.User.Username, "username", 1)
	parts := strings.Split(strings.ToLower(msg), " ")
	partsunchanged := strings.Split(msg, " ")

	log.Info(u.Username + " sent " + m.Content)
	log.Info("Which splits into:")
	log.Info(parts)

	channel, _ := discord.State.Channel(m.ChannelID)
	if channel == nil {
		log.WithFields(log.Fields{
			"channel": m.ChannelID,
			"message": m.ID,
		}).Warning("Failed to grab channel")
		return
	}

	guild, _ := discord.State.Guild(channel.GuildID)
	if guild == nil {
		log.WithFields(log.Fields{
			"guild":   channel.GuildID,
			"channel": channel,
			"message": m.ID,
		}).Warning("Failed to grab guild")
		return
	}


	// Champ tracking
	if parts[0] == "!addchamp" && m.Author.ID == "110110924102205440"  {

		log.Info("!addchamp has been recieved")

		removedspaces := strings.SplitN(msg," ",2)
    	addchamp := removedspaces[1] 

        f, err := os.OpenFile("champ.txt", os.O_APPEND|os.O_WRONLY, 0600)
        if err != nil {
            panic(err)
        }

        defer f.Close()

        if _, err = f.WriteString(addchamp+"\n"); err != nil {
            panic(err)
        }
    }


    if m.Content == "!champ" {

    	log.Info("!champ has been recieved")

        file, _  := os.Open("champ.txt")
        data, _  := ioutil.ReadAll(file)
        stringdata := fmt.Sprintf("%s", data)
        discord.ChannelMessageSend(channel.ID,stringdata)
        file.Close()

    }
    
    //Sends a direct message with the list of possible commands
	if m.Content == "!help" {

		log.Info("!help has been recieved")

		dm, _ := s.UserChannelCreate(u.ID)
		s.ChannelMessageSend(dm.ID, "Commands: http://pastebin.com/9xN5MxfT")
	}
    
    if m.Content == "!stop" && m.Author.ID == "97099676871823360"{

    	os.Exit(0)

    }

	if m.Content == "!birthday" {

		log.Info("!birthday has been recieved")

		for i := range BDate {
			log.Info(i)
        	if BDate[i] == (strconv.Itoa(now.Day())  + "/" + strconv.Itoa(int(now.Month()))) {  
				log.Info("Happy Birthday "+ BName[i])
				discord.ChannelMessageSend("203630579617366016", "!birthdaysound")
				discord.ChannelMessageSend("203630579617366016", "!name per")
				
           	} 
        }
	}    

	if parts[0] == "!sr" {
		log.Info("!sr has been recieved")

		website := ("https://playoverwatch.com/en-us/career/pc/eu/" + partsunchanged[1])

		log.Info("Checking " + website + " for Skill Ranking" )
		doc, err := goquery.NewDocument(website) 
		if err != nil {
			log.Fatal(err)
		}
	  	//Thanks to my boy Jake for working this one out.
	  	doc.Find(".masthead-player-progression:nth-child(3) > div:nth-child(2) > div:nth-child(2)").Each(func(i int, s *goquery.Selection) {
	    	rank := s.Text()

	    	discord.ChannelMessageSend(channel.ID, partsunchanged[1] + " is Skill Rank " + rank )
		})
	}


    //Removes commands after they have been executed to reduce spam
	deleteID := m.ID
	s.ChannelMessageDelete(channel.ID, deleteID)
	log.Info(deleteID + " has been deleted")
    
	// Find the collection for the command we got
	for _, coll := range COLLECTIONS {
		if scontains(parts[0], coll.Commands...) {
			// If they passed a specific sound effect, find and select that (otherwise play nothing)
			var sound *Sound
			if len(parts) > 1 {
				for _, s := range coll.Sounds {
					if parts[1] == s.Name {
						sound = s
					}
				}
				if sound == nil {
					return
				}
			}
			go enqueuePlay(m.Author, guild, coll, sound)
			return
		}
	}
}
