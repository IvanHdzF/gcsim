package mona

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/enemy"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

//When a Normal Attack hits, there is a 20% chance that it will be automatically followed by a Charged Attack.
//This effect can only occur once every 5s.
func (c *char) c2(a combat.AttackCB) {
	if c.Base.Cons < 2 {
		return
	}
	if c.Core.Rand.Float64() > .2 {
		return
	}
	if c.c2icd > c.Core.F {
		return
	}
	c.c2icd = c.Core.F + 300 //every 5 seconds
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Charge Attack",
		AttackTag:  combat.AttackTagExtra,
		ICDTag:     combat.ICDTagNone,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeDefault,
		Element:    attributes.Hydro,
		Durability: 25,
		Mult:       charge[c.TalentLvlAttack()],
	}

	c.Core.QueueAttack(ai, combat.NewCircleHit(c.Core.Combat.Player(), 0.3, false, combat.TargettableEnemy), 0, 0)
}

//When any party member attacks an opponent affected by an Omen, their CRIT Rate is increased by 15%.
func (c *char) c4() {
	m := make([]float64, attributes.EndStatType)
	m[attributes.CR] = 0.15

	for _, char := range c.Core.Player.Chars() {
		char.AddAttackMod(character.AttackMod{
			Base: modifier.NewBase("mona-c4", -1),
			Amount: func(_ *combat.AttackEvent, t combat.Target) ([]float64, bool) {
				x, ok := t.(*enemy.Enemy)
				if !ok {
					return nil, false
				}
				//ok only if either bubble or omen is present
				if x.StatusIsActive(bubbleKey) || x.StatusIsActive(omenKey) {
					return m, true
				}
				return nil, false
			},
		})
	}
}
