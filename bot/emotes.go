package main

import (
	"os"
	"strings"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

//List of all the Emotes
var (

Emotes = [...]string{

"Cuck",
"Jakka",
"JakesGame",
"SUL",
"Bush",

}
)

func CheckforEmote(s *discordgo.Session, m *discordgo.MessageCreate){

	msg := strings.Replace(m.ContentWithMentionsReplaced(), s.State.Ready.User.Username, "username", 1)
	parts := strings.Split(msg, " ")

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

	for i := range parts {
	    for j := range Emotes {
	        if parts[i] == Emotes[j] {  
	        	gopath := os.Getenv("GOPATH")
	        	log.Info(gopath+"/bin/emotes/" + Emotes[j] + ".png")
	            file, err := os.Open(filepath.FromSlash(gopath+"/bin/emotes/" + Emotes[j] + ".png"))
	            if err != nil {
    				log.Fatal(err)
    				return 
    			}
	            s.ChannelFileSend(channel.ID ,Emotes[j] + ".png", file)
	            u := m.Author
	            s.ChannelMessageSend("203630579617366016", (u.Username + " sent Emote:" + Emotes[j]))
	            log.Info("Sending Emote " + Emotes[j] )
	            file.Close()
	            if i == 0 {
	            	deleteID := m.ID
					s.ChannelMessageDelete(channel.ID, deleteID)
					log.Info(deleteID + " has been deleted")
	            }	       
	           	return   
	        } 
	    }  	  
	}
}