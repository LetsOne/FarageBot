package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"strconv"
	"time"
	"path/filepath"
	"bufio"

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


BattleTag = make([]string, 0)
SkillRank = make([]string, 0)
lines = make([]string, 0)


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

			fileHandle, _ := os.Open(filepath.FromSlash(gopath+"/bin/skillrank.txt"))
			defer fileHandle.Close()
			fileScanner := bufio.NewScanner(fileHandle)

			for fileScanner.Scan() {
				fmt.Println(fileScanner.Text())
				line := strings.Split(fileScanner.Text(), " ")
				BattleTag = append(BattleTag, line[0])
				SkillRank = append(SkillRank, line[1])
				fmt.Println(BattleTag)
				fmt.Println(SkillRank)
			}

		    doc, err := goquery.NewDocument("https://playoverwatch.com/en-us/career/pc/eu/" + partsunchanged[1]) 
		    if err != nil {
		        log.Fatal(err)
		    }		
		    doc.Find(".masthead-player-progression:nth-child(3) > div:nth-child(2) > div:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		        rank := s.Text()

			    for i := range BattleTag {
			    	fmt.Println(i)
			    	fmt.Println(BattleTag[i])
			    	if partsunchanged[1] == BattleTag[i] {
			    		if rank == SkillRank[i] {
			    			fmt.Println(SkillRank[i])
			    			discord.ChannelMessageSend(channel.ID, partsunchanged[1] + " is Skill Rank " + rank)
			    			return
			    		} else {
			    			newrankint, _ := strconv.Atoi(rank)
			    			oldrankint, _ := strconv.Atoi(SkillRank[i])
			    			rankdiff := newrankint - oldrankint 
			    			fmt.Println("diff:",rankdiff)
			    			if rankdiff > 0 {
			    				discord.ChannelMessageSend(channel.ID, partsunchanged[1] + " is Skill Rank " + rank + " which is " + strconv.Itoa(rankdiff) + " more than last time you checked." )
			    			} else {
			    				rankdiff = -rankdiff
			    				discord.ChannelMessageSend(channel.ID, partsunchanged[1] + " is Skill Rank " + rank + " which is " + strconv.Itoa(rankdiff) + " less than last time you checked." )			    				
			    			}
			    			SkillRank[i] = rank
			    			for j := range BattleTag {
			    				skillranklines := BattleTag[j] + " " + SkillRank[j]
			    				lines = append(lines,skillranklines)
			    			}
			    			output := strings.Join(lines, "\n")
					        err = ioutil.WriteFile((filepath.FromSlash(gopath+"/bin/skillrank.txt")), []byte(output), 0644)
					        if err != nil {
					                log.Fatalln(err)
					        }	    			
			    			return
						}	
					}

				}
		        appendrank := "\n" + partsunchanged[1] + " "  +  rank 

		        f, err := os.OpenFile(filepath.FromSlash(gopath+"/bin/skillrank.txt"), os.O_APPEND|os.O_WRONLY, 0600)
		        if err != nil {
		            panic(err)
		        }

		        defer f.Close()

		        if _, err = f.WriteString(appendrank); err != nil {
		            panic(err)
		        }

			})	
}
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