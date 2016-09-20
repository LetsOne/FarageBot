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

func CheckforEmote(u *discordgo.User, partsunchanged []string,channel *discordgo.Channel, s *discordgo.Session, m *discordgo.MessageCreate){

	for i := range partsunchanged {
		log.Info(partsunchanged[i])
	    for j := range EmotesName {
	    	log.Info(EmotesName[j])
	        if partsunchanged[i] == EmotesName[j] {  
	            file, err := os.Open(filepath.FromSlash(gopath+"/bin/emotes/" + EmotesExt[j]))
	            if err != nil {
    				log.Fatal(err)
    				return 
    			}
	            s.ChannelFileSend(channel.ID ,EmotesExt[j], file)
	            s.ChannelMessageSend(botlogs, (u.Username + " sent " + EmotesExt[j]))
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