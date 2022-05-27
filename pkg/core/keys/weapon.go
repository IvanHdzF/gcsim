package keys

import (
	"encoding/json"
	"errors"
	"strings"
)

type Weapon int

func (c *Weapon) MarshalJSON() ([]byte, error) {
	return json.Marshal(weaponNames[*c])
}

func (c *Weapon) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	s = strings.ToLower(s)
	for i, v := range weaponNames {
		if v == s {
			*c = Weapon(i)
			return nil
		}
	}
	return errors.New("unrecognized weapon key")
}

func (c Weapon) String() string {
	return weaponNames[c]
}

var weaponNames = []string{
	"",
	"akuoumaru",
	"alleyhunter",
	"amenomakageuchi",
	"amosbow",
	"apprenticesnotes",
	"aquilafavonia",
	"beginnersprotector",
	"blackcliffagate",
	"blackclifflongsword",
	"blackcliffpole",
	"blackcliffslasher",
	"blackcliffwarbow",
	"blacktassel",
	"bloodtaintedgreatsword",
	"calamityqueller",
	"cinnabarspindle",
	"compoundbow",
	"coolsteel",
	"crescentpike",
	"darkironsword",
	"deathmatch",
	"debateclub",
	"dodocotales",
	"dragonsbane",
	"dragonspinespear",
	"dullblade",
	"elegyfortheend",
	"emeraldorb",
	"engulfinglightning",
	"everlastingmoonglow",
	"eyeofperception",
	"favoniuscodex",
	"favoniusgreatsword",
	"favoniuslance",
	"favoniussword",
	"favoniuswarbow",
	"ferrousshadow",
	"festeringdesire",
	"filletblade",
	"freedomsworn",
	"frostbearer",
	"hakushinring",
	"halberd",
	"hamayumi",
	"harangeppakufutsu",
	"harbingerofdawn",
	"huntersbow",
	"ironpoint",
	"ironsting",
	"kagurasverity",
	"katsuragikirinagamasa",
	"kitaincrossspear",
	"lionsroar",
	"lithicblade",
	"lithicspear",
	"lostprayertothesacredwinds",
	"luxurioussealord",
	"magicguide",
	"mappamare",
	"memoryofdust",
	"messenger",
	"mistsplitterreforged",
	"mitternachtswaltz",
	"mouunsmoon",
	"oathsworneye",
	"oldmercspal",
	"otherworldlystory",
	"pocketgrimoire",
	"polarstar",
	"predator",
	"primordialjadecutter",
	"primordialjadewingedspear",
	"prototypeamber",
	"prototypearchaic",
	"prototypecrescent",
	"prototyperancour",
	"prototypestarglitter",
	"rainslasher",
	"ravenbow",
	"recurvebow",
	"redhornstonethresher",
	"royalbow",
	"royalgreatsword",
	"royalgrimoire",
	"royallongsword",
	"royalspear",
	"rust",
	"sacrificialbow",
	"sacrificialfragments",
	"sacrificialgreatsword",
	"sacrificialsword",
	"seasonedhuntersbow",
	"serpentspine",
	"sharpshootersoath",
	"silversword",
	"skyridergreatsword",
	"skyridersword",
	"skywardatlas",
	"skywardblade",
	"skywardharp",
	"skywardpride",
	"skywardspine",
	"slingshot",
	"snowtombedstarsilver",
	"solarpearl",
	"songofbrokenpines",
	"staffofhoma",
	"summitshaper",
	"swordofdescension",
	"thealleyflash",
	"thebell",
	"theblacksword",
	"thecatch",
	"theflute",
	"thestringless",
	"theunforged",
	"theviridescenthunt",
	"thewidsith",
	"thrillingtalesofdragonslayers",
	"thunderingpulse",
	"travelershandysword",
	"twinnephrite",
	"vortexvanquisher",
	"wastergreatsword",
	"wavebreakersfin",
	"whiteblind",
	"whiteirongreatsword",
	"whitetassel",
	"windblumeode",
	"wineandsong",
	"wolfsgravestone",
}

const (
	NoWeapon Weapon = iota
	Akuoumaru
	AlleyHunter
	AmenomaKageuchi
	AmosBow
	ApprenticesNotes
	AquilaFavonia
	BeginnersProtector
	BlackcliffAgate
	BlackcliffLongsword
	BlackcliffPole
	BlackcliffSlasher
	BlackcliffWarbow
	BlackTassel
	BloodtaintedGreatsword
	CalamityQueller
	CinnabarSpindle
	CompoundBow
	CoolSteel
	CrescentPike
	DarkIronSword
	Deathmatch
	DebateClub
	DodocoTales
	DragonsBane
	DragonspineSpear
	DullBlade
	ElegyForTheEnd
	EmeraldOrb
	EngulfingLightning
	EverlastingMoonglow
	EyeOfPerception
	FavoniusCodex
	FavoniusGreatsword
	FavoniusLance
	FavoniusSword
	FavoniusWarbow
	FerrousShadow
	FesteringDesire
	FilletBlade
	FreedomSworn
	Frostbearer
	HakushinRing
	Halberd
	Hamayumi
	HaranGeppakuFutsu
	HarbingerOfDawn
	HuntersBow
	IronPoint
	IronSting
	KagurasVerity
	KatsuragikiriNagamasa
	KitainCrossSpear
	LionsRoar
	LithicBlade
	LithicSpear
	LostPrayerToTheSacredWinds
	LuxuriousSeaLord
	MagicGuide
	MappaMare
	MemoryOfDust
	Messenger
	MistsplitterReforged
	MitternachtsWaltz
	MouunsMoon
	OathswornEye
	OldMercsPal
	OtherworldlyStory
	PocketGrimoire
	PolarStar
	Predator
	PrimordialJadeCutter
	PrimordialJadeWingedSpear
	PrototypeAmber
	PrototypeArchaic
	PrototypeCrescent
	PrototypeRancour
	PrototypeStarglitter
	Rainslasher
	RavenBow
	RecurveBow
	RedhornStonethresher
	RoyalBow
	RoyalGreatsword
	RoyalGrimoire
	RoyalLongsword
	RoyalSpear
	Rust
	SacrificialBow
	SacrificialFragments
	SacrificialGreatsword
	SacrificialSword
	SeasonedHuntersBow
	SerpentSpine
	SharpshootersOath
	SilverSword
	SkyriderGreatsword
	SkyriderSword
	SkywardAtlas
	SkywardBlade
	SkywardHarp
	SkywardPride
	SkywardSpine
	Slingshot
	SnowTombedStarsilver
	SolarPearl
	SongOfBrokenPines
	StaffOfHoma
	SummitShaper
	SwordOfDescension
	TheAlleyFlash
	TheBell
	TheBlackSword
	TheCatch
	TheFlute
	TheStringless
	TheUnforged
	TheViridescentHunt
	TheWidsith
	ThrillingTalesOfDragonSlayers
	ThunderingPulse
	TravelersHandySword
	TwinNephrite
	VortexVanquisher
	WasterGreatsword
	WavebreakersFin
	Whiteblind
	WhiteIronGreatsword
	WhiteTassel
	WindblumeOde
	WineAndSong
	WolfsGravestone
)
