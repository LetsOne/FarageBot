package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"strconv"
	"time"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
)

var (

BName = [...]string{

"Per",

}

BDate = [...]string{

"8/9",

}

archie = "97099676871823360"
mark = "110110924102205440"

)




func CommandsAndSound(u *discordgo.User, msg string, parts []string,partsunchanged []string,channel *discordgo.Channel, guild *discordgo.Guild, s *discordgo.Session, m *discordgo.MessageCreate) {

	s.ChannelMessageSend("203630579617366016",u.Username + " sent " + msg)

	switch parts[0]{

	case "!addchamp":

		if m.Author.ID == mark {
			log.Info("!addchamp has been recieved")

			removedspaces := strings.SplitN(msg," ",2)
	    	addchamp := removedspaces[1] 

	        f, err := os.OpenFile("champ.txt", os.O_APPEND|os.O_WRONLY, 0600)
	        if err != nil {
	            panic(err)
	        }

	        defer f.Close()

	        if _, err = f.WriteString(addchamp+"\n\n"); err != nil {
	            panic(err)
	        }
	    }
    
    case "!champ":

    	log.Info("!champ has been recieved")
        file, err  := os.Open(filepath.FromSlash(gopath+"/bin/champ.txt"))
	    if err != nil {
	        panic(err)
	    }
        data, _  := ioutil.ReadAll(file)
        stringdata := fmt.Sprintf("%s", data)
        discord.ChannelMessageSend(channel.ID,stringdata)
        file.Close()


	case "!help":

		log.Info("!help has been recieved")

		dm, _ := s.UserChannelCreate(u.ID)
		s.ChannelMessageSend(dm.ID, "Commands: http://pastebin.com/9xN5MxfT")

    
	case "!stop":

		if m.Author.ID == archie {
			s.Close()
			os.Exit(0)
		}

	case "!birthday":

		log.Info("!birthday has been recieved")

		for i := range BDate {
			log.Info(i)
        	if BDate[i] == (strconv.Itoa(now.Day())  + "/" + strconv.Itoa(int(now.Month()))) {  
				log.Info("Happy Birthday "+ BName[i])
				coll := BIRTHDAYSOUND
				var sound *Sound
				go enqueuePlay(m.Author, guild, coll, sound)
				time.Sleep(200 * time.Millisecond)
				lcaseBName := strings.ToLower(BName[i])
				coll = NAME
				for _, sound = range coll.Sounds {
					if sound.Name == lcaseBName {
					go enqueuePlay(m.Author, guild, coll, sound)
					}

				}
           	} 
        }

    case "!sr":

    	if len(parts) > 1 {

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
			})}
	case "!uptime":
        uptimeDur := time.Since(initialTime)
        uptimeDur = Round(uptimeDur, time.Second)
        uptimeDurString := uptimeDur.String()
        discord.ChannelMessageSend(channel.ID,"I have been spreading memes for : " + uptimeDurString)
	case "!reloademotes":

		EmoteLookUp()

	default:			
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
					}
				}
				go enqueuePlay(m.Author, guild, coll, sound)
			}
		}
	}
	

    //Removes commands after they have been executed to reduce spam
	deleteID := m.ID
	s.ChannelMessageDelete(channel.ID, deleteID)
	log.Info(deleteID + " has been deleted")
	return
    
}
func Round(d, r time.Duration) time.Duration {
	if r <= 0 {
		return d
	}
	neg := d < 0
	if neg {
		d = -d
	}
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	if neg {
		return -d
	}
	return d
}