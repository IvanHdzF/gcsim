package tartaglia

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

var (
	burstMeleeFrames  []int
	burstRangedFrames []int
)

const (
	burstMeleeHitmark  = 69
	burstRangedHitmark = 70
)

func init() {
	// burst (melee) -> x
	burstMeleeFrames = frames.InitAbilSlice(102)
	burstMeleeFrames[action.ActionAttack] = 101
	burstMeleeFrames[action.ActionSkill] = 102
	burstMeleeFrames[action.ActionDash] = 102
	burstMeleeFrames[action.ActionJump] = 102
	burstMeleeFrames[action.ActionSwap] = 100

	// burst (ranged) -> x
	burstRangedFrames = frames.InitAbilSlice(55)
	burstRangedFrames[action.ActionAttack] = 53
	burstRangedFrames[action.ActionDash] = 54
	burstRangedFrames[action.ActionJump] = 53
	burstRangedFrames[action.ActionSwap] = 52
}

// Performs a different attack depending on the stance in which it is cast.
// Ranged Stance: dealing AoE Hydro DMG. Apply Riptide status to enemies hit. Returns 20 Energy after use
// Melee Stance: dealing AoE Hydro DMG. Triggers Riptide Blast (clear riptide after triggering riptide blast)
func (c *char) Burst(p map[string]int) action.ActionInfo {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Ranged Stance: Flash of Havoc",
		AttackTag:  combat.AttackTagElementalBurst,
		ICDTag:     combat.ICDTagNone,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypeDefault,
		Element:    attributes.Hydro,
		Durability: 50,
		Mult:       burst[c.TalentLvlBurst()],
	}

	cancels := burstRangedFrames
	hitmark := burstRangedHitmark
	cb := c.rangedBurstApplyRiptide

	if c.StatusIsActive(meleeKey) {
		ai.Abil = "Melee Stance: Light of Obliteration"
		ai.Mult = meleeBurst[c.TalentLvlBurst()]
		cancels = burstMeleeFrames
		hitmark = burstMeleeHitmark
		cb = c.rtBlastCallback
		if c.Base.Cons >= 6 {
			c.mlBurstUsed = true
		}
	} else {
		c.Core.Tasks.Add(func() {
			c.AddEnergy("tartaglia-ranged-burst-refund", 20)
		}, 4)
	}

	c.Core.QueueAttack(ai, combat.NewCircleHit(c.Core.Combat.Player(), 5, false, combat.TargettableEnemy), hitmark, hitmark, cb)

	if c.StatusIsActive(meleeKey) {
		c.ConsumeEnergy(71)
		c.SetCDWithDelay(action.ActionBurst, 900, 66)
	} else {
		c.ConsumeEnergy(3)
		c.SetCDWithDelay(action.ActionBurst, 900, 0)
	}

	return action.ActionInfo{
		Frames:          frames.NewAbilFunc(cancels),
		AnimationLength: cancels[action.InvalidAction],
		CanQueueAfter:   cancels[action.ActionDash], // earliest cancel
		State:           action.BurstState,
	}
}
