package ningguang

import (
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/construct"
)

type screen struct {
	src    int
	expiry int
	char   *char
}

func (c *char) newScreen(dur int) *screen {
	return &screen{
		src:    c.Core.F,
		expiry: c.Core.F + dur,
		char:   c,
	}
}

func (c *screen) OnDestruct() {
	if c.char.Base.Cons < 2 {
		return
	}
	//make sure last reset is more than 6 seconds ago
	if c.char.c2reset <= c.char.Core.F-360 && c.char.Cooldown(action.ActionSkill) > 0 {
		//reset cd
		c.char.c2reset = c.char.Core.F
		c.char.ResetActionCooldown(action.ActionSkill)
	}
}

func (c *screen) Key() int                         { return c.src }
func (c *screen) Type() construct.GeoConstructType { return construct.GeoConstructNingSkill }
func (c *screen) Expiry() int                      { return c.expiry }
func (c *screen) IsLimited() bool                  { return true }
func (c *screen) Count() int                       { return 1 }
