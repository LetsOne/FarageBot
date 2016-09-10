package main

import (
	"os"
	"strings"

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
	            file, err := os.Open("emotes/" + Emotes[j] + ".png")
	            if err != nil {
    				log.Fatal(err)
    				return 
    			}
	            s.ChannelFileSend(channel.ID ,Emotes[j] + ".png", file)
	            log.Info("Sending Emote " + Emotes[j] )
	            file.Close()	       
	           	return   
	        } 
	    }  	  
	}
}