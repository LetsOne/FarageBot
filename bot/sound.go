package main

import (

)

//SoundCollection represents
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
		createSound("cavalry", 50, 250),
		createSound("cheers", 50, 250),
		createSound("deja", 50, 250),
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
		createSound("gg", 50, 250),
		createSound("leet", 50, 250),
		createSound("lol", 50, 250),
		createSound("ree", 50, 250),
		createSound("winky", 50, 250),
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
		createSound("abaddon", 50, 250),
		createSound("bristle", 50, 250),
		createSound("centaur", 50, 250),
		createSound("earthspirit", 50, 250),
		createSound("eldertitan", 50, 250),
		createSound("jakiro", 50, 250),
		createSound("legion", 50, 250),
		createSound("magnus", 50, 250),
		createSound("medusa", 50, 250),
		createSound("meepo", 50, 250),
		createSound("nyx", 50, 250),
		createSound("omniknight", 50, 250),
		createSound("oracle", 50, 250),
		createSound("orge", 50, 250),
		createSound("phoenix", 50, 250),
		createSound("skeletonking", 50, 250),
		createSound("skywrath", 10, 250),
		createSound("spiritbreaker", 50, 250),
		createSound("timbersaw", 50, 250),
		createSound("treant", 50, 250),
		createSound("wisp", 50, 250),

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

var KEEM *SoundCollection = &SoundCollection{
	Prefix: "keem",
	Commands: []string{
		"!keem",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var JASON *SoundCollection = &SoundCollection{
	Prefix: "jason",
	Commands: []string{
		"!jason",
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
		createSound("absolutely", 50, 250),
		createSound("barrel", 50, 250),
		createSound("blinkdagger", 50, 250),
		createSound("bread", 1, 250),
		createSound("deal", 50, 250),
		createSound("diplomatic", 50, 250),
		createSound("dirtybitch", 50, 250),
		createSound("dota", 50, 250),
		createSound("gandalf", 50, 250),
		createSound("gyroshit", 50, 250),
		createSound("hazard", 50, 250),
		createSound("hell", 50, 250),
		createSound("hittheroad", 50, 250),
		createSound("hlhf", 50, 250),
		createSound("illuminate", 50, 250),
		createSound("jizz", 50, 250),
		createSound("lawn", 50, 250),
		createSound("mint", 50, 250),
		createSound("mother", 50, 250),
		createSound("no", 50, 250),
		createSound("ohnoyou", 50, 250),
		createSound("ouch", 1, 250),
		createSound("pirate", 50, 250),
		createSound("pureskill", 50, 250),
		createSound("redundant", 50, 250),
		createSound("reslove", 50, 250),
		createSound("skillshot", 50, 250),
		createSound("smarter", 50, 250),
		createSound("thereitgoes", 50, 250),
		createSound("ugly", 50, 250),
		createSound("vipershit", 50, 250),
		createSound("whynot", 50, 250),
		createSound("worse", 50, 250),
	},
}

var TWELVE *SoundCollection = &SoundCollection{
	Prefix: "twelve",
	Commands: []string{
		"!twelve",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var NOOT *SoundCollection = &SoundCollection{
	Prefix: "noot",
	Commands: []string{
		"!noot",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},
}

var PINK *SoundCollection = &SoundCollection{
	Prefix: "pink",
	Commands: []string{
		"!pink",
	},
	Sounds: []*Sound{
		createSound("stupid", 50, 250),
		createSound("cunt", 50, 250),
		createSound("stop", 50, 250),
		createSound("shut", 50, 250),
		createSound("nobody", 50, 250),
	},
}

var BIRTHDAYSOUND *SoundCollection = &SoundCollection{
	Prefix: "birthday",
	Commands: []string{
		"",
	},
	Sounds: []*Sound{
		createSound("one", 50, 250),
	},

}

var NAME *SoundCollection = &SoundCollection{
	Prefix: "name",
	Commands: []string{
		"",
	},
	Sounds: []*Sound{
		createSound("per", 50, 250),
	},
}



var COLLECTIONS []*SoundCollection = []*SoundCollection{
	BAG,
	BIRTHDAYSOUND,
	DISASTER,
	DONETHIS,	
	DOTA,
	DVA,
	HEY,
	JASON,
	KEEM,
	KHALED,
	KILLYOURSELF,
	MCREE,
	MEI,
	NAME,
	NOBALLS,
	NOOT,
	PINK,
	PROFANITY,
	SUH,
	TORB,
	TRACER,
	TRY,
	TWELVE,
	WAKE,
	WHATCHA,
	WOAW,
	WOW,

}
