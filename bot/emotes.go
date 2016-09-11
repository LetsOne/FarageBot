package main

import (
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

//List of all the Emotes
var (

	EmotesExt = make([]string, 0)
	EmotesName = make([]string, 0)
)


func EmoteLookUp() {

	EmotesExt = make([]string, 0)
	EmotesName = make([]string, 0)

    files, _ := ioutil.ReadDir(filepath.FromSlash(gopath +"/bin/emotes/"))
    for _, f := range files {
    	EmotesExt = append(EmotesExt, f.Name())
    }

    for i := range EmotesExt{
    	split := strings.Split(EmotesExt[i], ".")
    	EmotesName = append(EmotesName,split[0])
    }

}

func CheckforEmote(s *discordgo.Session, m *discordgo.MessageCreate){

	u := m.Author

	msg := strings.Replace(m.ContentWithMentionsReplaced(), s.State.Ready.User.Username, "username", 1)
	parts := strings.Split(msg, " ")

	log.Info(u.Username + " sent " + m.Content)
	log.Info("Which splits into: ", parts)

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
	    for j := range EmotesName {
	        if parts[i] == EmotesName[j] {  
	            file, err := os.Open(filepath.FromSlash(gopath+"/bin/emotes/" + EmotesExt[j]))
	            if err != nil {
    				log.Fatal(err)
    				return 
    			}
	            s.ChannelFileSend(channel.ID ,EmotesExt[j], file)
	            u := m.Author
	            s.ChannelMessageSend("203630579617366016", (u.Username + " sent " + EmotesExt[j]))
	            log.Info("Sending Emote " + EmotesExt[j] )
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