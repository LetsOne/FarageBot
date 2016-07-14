package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
//	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

var (
	// discordgo session
	discord *discordgo.Session

	// Map of Guild id's to *Play channels, used for queuing and rate-limiting guilds
	queues map[string]chan *Play = make(map[string]chan *Play)

	// Sound encoding settings
	BITRATE        = 128
	MAX_QUEUE_SIZE = 3
)

// Play represents an individual use of the !airhorn command
type Play struct {
	GuildID   string
	ChannelID string
	UserID    string
	Sound     *Sound

	// The next play to occur after this, only used for chaining sounds like anotha
	Next *Play

	// If true, this was a forced play using a specific airhorn sound name
	Forced bool
}

type SoundCollection struct {
	Prefix    string
	Commands  []string
	Sounds    []*Sound
	ChainWith *SoundCollection

	soundRange int
}

// Sound represents a sound clip
type Sound struct {
	Name string

	// Weight adjust how likely it is this song will play, higher = more likely
	Weight int

	// Delay (in milliseconds) for the bot to wait before sending the disconnect request
	PartDelay int

	// Buffer to store encoded PCM packets
	buffer [][]byte
}

// Array of all the sounds we have

var KHALED *SoundCollection = &SoundCollection{
	Prefix:    "another",
	Commands: []string{
		"!anotha",
		"!anothaone",
	},
	Sounds: []*Sound{
		createSound("one_classic", 1, 250),
	},
}

var KILLYOURSELF *SoundCollection = &SoundCollection{
	Prefix: "kys",
	Commands: []string{
		"!kys",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var TRACER *SoundCollection = &SoundCollection{
	Prefix: "tracer",
	Commands: []string{
		"!tracer",
	},
	Sounds: []*Sound{
		createSound("deja", 50, 250),
		createSound("cheers", 50, 250),
		createSound("cavalry", 50, 250),
	},
}

var MEI *SoundCollection = &SoundCollection{
	Prefix: "mei",
	Commands: []string{
		"!mei",
	},
	Sounds: []*Sound{
		createSound("amazing", 50, 250),
	},
}

var TORB *SoundCollection = &SoundCollection{
	Prefix: "torb",
	Commands: []string{
		"!torb",
	},
	Sounds: []*Sound{
		createSound("turret", 50, 250),
	},
}

var MCREE *SoundCollection = &SoundCollection{
	Prefix: "mcree",
	Commands: []string{
		"!mcree",
	},
	Sounds: []*Sound{
		createSound("noon", 50, 250),
	},
}

var HEY *SoundCollection = &SoundCollection{
	Prefix: "hey",
	Commands: []string{
		"!hey",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var DVA *SoundCollection = &SoundCollection{
	Prefix: "dva",
	Commands: []string{
		"!dva",
	},
	Sounds: []*Sound{
		createSound("winky", 50, 250),
		createSound("lol", 50, 250),
		createSound("leet", 50, 250),
		createSound("ree", 50, 250),
		createSound("gg", 50, 250),
	},
}

var DONETHIS *SoundCollection = &SoundCollection{
	Prefix: "donethis",
	Commands: []string{
		"!donethis",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var BAG *SoundCollection = &SoundCollection{
	Prefix: "bag",
	Commands: []string{
		"!bag",
	},
	Sounds: []*Sound{
		createSound("skywrath", 10, 250),
		createSound("spiritbreaker", 50, 250),
		createSound("orge", 50, 250),
		createSound("phoenix", 50, 250),
		createSound("skeletonking", 50, 250),
		createSound("tree", 50, 250),
		createSound("timber", 50, 250),
		createSound("jakiro", 50, 250),
		createSound("legion", 50, 250),
		createSound("magnus", 50, 250),
		createSound("medusa", 50, 250),
		createSound("meepo", 50, 250),
		createSound("omniknight", 50, 250),
		createSound("nyx", 50, 250),
		createSound("oracle", 50, 250),
		createSound("io", 50, 250),
		createSound("abaddon", 50, 250),
		createSound("centaur", 50, 250),
		createSound("earthspirit", 50, 250),
		createSound("eldertitan", 50, 250),
		createSound("bristle", 50, 250),

	},
}

var DISASTER *SoundCollection = &SoundCollection{
	Prefix: "disaster",
	Commands: []string{
		"!disaster",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var PROFANITY *SoundCollection = &SoundCollection{
	Prefix: "profanity",
	Commands: []string{
		"!profanity",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var WOAW *SoundCollection = &SoundCollection{
	Prefix: "woaw",
	Commands: []string{
		"!woaw",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var WOW *SoundCollection = &SoundCollection{
	Prefix: "wow",
	Commands: []string{
		"!wow",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var TRY *SoundCollection = &SoundCollection{
	Prefix: "try",
	Commands: []string{
		"!try",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var WHATCHA *SoundCollection = &SoundCollection{
	Prefix: "whatcha",
	Commands: []string{
		"!whatcha",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var NOBALLS *SoundCollection = &SoundCollection{
	Prefix: "noballs",
	Commands: []string{
		"!noballs",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var SUH *SoundCollection = &SoundCollection{
	Prefix: "suh",
	Commands: []string{
		"!suh",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var WAKE *SoundCollection = &SoundCollection{
	Prefix: "wake",
	Commands: []string{
		"!wake",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var DOTA *SoundCollection = &SoundCollection{
	Prefix: "dota",
	Commands: []string{
		"!dota",
	},
	Sounds: []*Sound{
		createSound("pudgefirstblood", 50, 250),
		createSound("rikiredundant", 50, 250),
		createSound("sandkingno", 50, 250),
		createSound("tinkerhazard", 50, 250),
		createSound("viperholyshit", 50, 250),
		createSound("windcouldbeworse", 50, 250),
		createSound("eldertitangetoff", 50, 250),
		createSound("gyroholyshit", 50, 250),
		createSound("kotlmint", 50, 250),
		createSound("kotlthereitgoes", 50, 250),
		createSound("kotlgandalf", 50, 250),
		createSound("lionbeentohellandbacktohellandback", 50, 250),
		createSound("lunaabsolutelynot", 50, 250),
		createSound("lunapirate", 50, 250),
		createSound("lunailluminate", 50, 250),
		createSound("necrodirtybitch", 50, 250),
		createSound("orgepureskill", 50, 250),
		createSound("orgeskillshot", 50, 250),
		createSound("orgesmarter", 50, 250),
		createSound("orgehittheroad", 50, 250),
		createSound("qopouch", 1, 250),
		createSound("ababbdonohnoyoudidnt", 50, 250),
		createSound("blinkdaggerrrrrr", 50, 250),
		createSound("brewmasterbarrel", 50, 250),
		createSound("bristleugly", 50, 250),
		createSound("bastionfaceamother", 50, 250),
		createSound("spreslove", 50, 250),
		createSound("sphlhf", 50, 250),
		createSound("spdiplomatic", 50, 250),
		createSound("spbread", 1, 250),
		createSound("spdota", 50, 250),
	},
}


var COLLECTIONS []*SoundCollection = []*SoundCollection{
	KILLYOURSELF,
	KHALED,
	TRACER,
	MEI,
	TORB,
	MCREE,
	HEY,
	DVA,
	DONETHIS,	
	BAG,
	DISASTER,
	PROFANITY,
	WOAW,
	WOW,
	TRY,
	WHATCHA,
	NOBALLS,
	SUH,
	WAKE,
	DOTA,

}

// Create a Sound struct
func createSound(Name string, Weight int, PartDelay int) *Sound {
	return &Sound{
		Name:      Name,
		Weight:    Weight,
		PartDelay: PartDelay,
		buffer:    make([][]byte, 0),
	}
}

func (sc *SoundCollection) Load() {
	for _, sound := range sc.Sounds {
		sc.soundRange += sound.Weight
		sound.Load(sc)
	}
}

func (s *SoundCollection) Random() *Sound {
	var (
		i      int
		number int = randomRange(0, s.soundRange)
	)

	for _, sound := range s.Sounds {
		i += sound.Weight

		if number < i {
			return sound
		}
	}
	return nil
}

// Load attempts to load an encoded sound file from disk
// DCA files are pre-computed sound files that are easy to send to Discord.
// If you would like to create your own DCA files, please use:
// https://github.com/nstafie/dca-rs
// eg: dca-rs --raw -i <input wav file> > <output file>
func (s *Sound) Load(c *SoundCollection) error {
	path := fmt.Sprintf("audio/%v_%v.dca", c.Prefix, s.Name)

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// read opus frame length from dca file
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil
		}

		if err != nil {
			fmt.Println("error reading from dca file :", err)
			return err
		}

		// read encoded pcm from dca file
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("error reading from dca file :", err)
			return err
		}

		// append encoded pcm data to the buffer
		s.buffer = append(s.buffer, InBuf)
	}
}

// Plays this sound over the specified VoiceConnection
func (s *Sound) Play(vc *discordgo.VoiceConnection) {
	vc.Speaking(true)
	defer vc.Speaking(false)

	for _, buff := range s.buffer {
		vc.OpusSend <- buff
	}
}

// Attempts to find the current users voice channel inside a given guild
func getCurrentVoiceChannel(user *discordgo.User, guild *discordgo.Guild) *discordgo.Channel {
	for _, vs := range guild.VoiceStates {
		if vs.UserID == user.ID {
			channel, _ := discord.State.Channel(vs.ChannelID)
			return channel
		}
	}
	return nil
}

// Returns a random integer between min and max
func randomRange(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

// Prepares a play
func createPlay(user *discordgo.User, guild *discordgo.Guild, coll *SoundCollection, sound *Sound) *Play {
	// Grab the users voice channel
	channel := getCurrentVoiceChannel(user, guild)
	if channel == nil {
		log.WithFields(log.Fields{
			"user":  user.ID,
			"guild": guild.ID,
		}).Warning("Failed to find channel to play sound in")
		return nil
	}

	// Create the play
	play := &Play{
		GuildID:   guild.ID,
		ChannelID: channel.ID,
		UserID:    user.ID,
		Sound:     sound,
		Forced:    true,
	}

	// If we didn't get passed a manual sound, generate a random one
	if play.Sound == nil {
		play.Sound = coll.Random()
		play.Forced = false
	}

	// If the collection is a chained one, set the next sound
	if coll.ChainWith != nil {
		play.Next = &Play{
			GuildID:   play.GuildID,
			ChannelID: play.ChannelID,
			UserID:    play.UserID,
			Sound:     coll.ChainWith.Random(),
			Forced:    play.Forced,
		}
	}

	return play
}

// Prepares and enqueues a play into the ratelimit/buffer guild queue
func enqueuePlay(user *discordgo.User, guild *discordgo.Guild, coll *SoundCollection, sound *Sound) {
	play := createPlay(user, guild, coll, sound)
	if play == nil {
		return
	}

	// Check if we already have a connection to this guild
	//   yes, this isn't threadsafe, but its "OK" 99% of the time
	_, exists := queues[guild.ID]

	if exists {
		if len(queues[guild.ID]) < MAX_QUEUE_SIZE {
			queues[guild.ID] <- play
		}
	} else {
		queues[guild.ID] = make(chan *Play, MAX_QUEUE_SIZE)
		playSound(play, nil)
	}
}

//test
// Play a sound
func playSound(play *Play, vc *discordgo.VoiceConnection) (err error) {
	log.WithFields(log.Fields{
		"play": play,
	}).Info("Playing sound")

	if vc == nil {
		vc, err = discord.ChannelVoiceJoin(play.GuildID, play.ChannelID, false, false)
		// vc.Receive = false
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Failed to play sound")
			delete(queues, play.GuildID)
			return err
		}
	}

	// If we need to change channels, do that now
	if vc.ChannelID != play.ChannelID {
		vc.ChangeChannel(play.ChannelID, false, false)
		time.Sleep(time.Millisecond * 125)
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(time.Millisecond * 32)

	// Play the sound
	play.Sound.Play(vc)

	// If there is another song in the queue, recurse and play that
	if len(queues[play.GuildID]) > 0 {
		play := <-queues[play.GuildID]
		playSound(play, vc)
		return nil
	}

	// If the queue is empty, delete it
	time.Sleep(time.Millisecond * time.Duration(play.Sound.PartDelay))
	delete(queues, play.GuildID)
	return nil
}


func onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Info("Recieved READY payload")
	s.UpdateStatus(0, "Champions of Regnum")
}

func onGuildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable != nil {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			s.ChannelMessageSend(channel.ID, "**AIRHORN BOT READY FOR HORNING. TYPE `!AIRHORN` WHILE IN A VOICE CHANNEL TO ACTIVATE**")
			return
		}
	}
}

func scontains(key string, options ...string) bool {
	for _, item := range options {
		if item == key {
			return true
		}
	}
	return false
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Content) <= 0 || m.Content[0] != '!'  {
		return
	}

	msg := strings.Replace(m.ContentWithMentionsReplaced(), s.State.Ready.User.Username, "username", 1)
	parts := strings.Split(strings.ToLower(msg), " ")

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


	if m.Content == "!help" {
		s.ChannelMessageSend(channel.ID, "Commands: http://pastebin.com/9xN5MxfT")
	}

	deleteID := m.ID
	s.ChannelMessageDelete(channel.ID, deleteID)
	
	}


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

func main() {
	var (
		Token      = flag.String("t", "", "Discord Authentication Token")
		Owner      = flag.String("o", "", "Owner ID")
		err        error
	)
	flag.Parse()

	if *Owner != "" {
		OWNER = *Owner
	}

	// Preload all the sounds
	log.Info("Preloading sounds...")
	for _, coll := range COLLECTIONS {
		coll.Load()
	}

	// Create a discord session
	log.Info("Starting discord session...")
	discord, err = discordgo.New(*Token)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to create discord session")
		return
	}

	discord.AddHandler(onReady)
	discord.AddHandler(onGuildCreate)
	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to create discord websocket connection")
		return
	}

	// We're running!
	log.Info("FarageBot has is up!")

	// Wait for a signal to quit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

