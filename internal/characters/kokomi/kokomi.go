package kokomi

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/core/player/character/profile"
)

func init() {
	core.RegisterCharFunc(keys.Kokomi, NewChar)
}

type char struct {
	*tmpl.Character
	skillFlatDmg  float64
	skillLastUsed int
	swapEarlyF    int
	c4ICDExpiry   int
}

func NewChar(s *core.Core, w *character.CharWrapper, _ profile.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = 70
	c.NormalHitNum = normalHitNum
	c.BurstCon = 3
	c.SkillCon = 5

	c.skillFlatDmg = 0
	c.skillLastUsed = 0
	c.swapEarlyF = 0
	c.c4ICDExpiry = 0

	w.Character = &c

	return nil
}

func (c *char) Init() error {
	c.a4()
	c.passive()
	c.onExitField()
	c.burstActiveHook()
	return nil
}
