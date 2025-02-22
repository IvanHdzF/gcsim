package yunjin

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/core/player/character/profile"
)

func init() {
	core.RegisterCharFunc(keys.Yunjin, NewChar)
}

type char struct {
	*tmpl.Character
	partyElementalTypes int
	c4bonus             []float64
	c6bonus             []float64
}

func NewChar(s *core.Core, w *character.CharWrapper, _ profile.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = 60
	c.NormalHitNum = normalHitNum
	c.BurstCon = 3
	c.SkillCon = 5

	c.partyElementalTypes = 0

	w.Character = &c

	return nil
}

// Occurs after all characters are loaded, so getPartyElementalTypeCounts works properly
func (c *char) Init() error {
	c.getPartyElementalTypeCounts()
	c.burstProc()
	if c.Base.Cons >= 4 {
		c.c4()
	}
	if c.Base.Cons >= 6 {
		c.c6bonus = make([]float64, attributes.EndStatType)
		c.c6bonus[attributes.AtkSpd] = .12
	}

	return nil
}

// Adds event to get the number of elemental types in the party for Yunjin A4
func (c *char) getPartyElementalTypeCounts() {
	partyElementalTypes := make(map[attributes.Element]int)
	for _, char := range c.Core.Player.Chars() {
		partyElementalTypes[char.Base.Element]++
	}
	for range partyElementalTypes {
		c.partyElementalTypes += 1
	}
	c.Core.Log.NewEvent("Yun Jin Party Elemental Types (A4)", glog.LogCharacterEvent, c.Index).
		Write("party_elements", c.partyElementalTypes)
}
