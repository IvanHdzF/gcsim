package combat

import (

	//characters
	_ "github.com/genshinsim/gsim/internal/characters/albedo"
	_ "github.com/genshinsim/gsim/internal/characters/amber"
	_ "github.com/genshinsim/gsim/internal/characters/ayaka"
	_ "github.com/genshinsim/gsim/internal/characters/beidou"
	_ "github.com/genshinsim/gsim/internal/characters/bennett"
	_ "github.com/genshinsim/gsim/internal/characters/chongyun"
	_ "github.com/genshinsim/gsim/internal/characters/diluc"
	_ "github.com/genshinsim/gsim/internal/characters/diona"
	_ "github.com/genshinsim/gsim/internal/characters/eula"
	_ "github.com/genshinsim/gsim/internal/characters/fischl"
	_ "github.com/genshinsim/gsim/internal/characters/ganyu"
	_ "github.com/genshinsim/gsim/internal/characters/hutao"
	_ "github.com/genshinsim/gsim/internal/characters/jean"
	_ "github.com/genshinsim/gsim/internal/characters/kaeya"
	_ "github.com/genshinsim/gsim/internal/characters/kazuha"
	_ "github.com/genshinsim/gsim/internal/characters/keqing"
	_ "github.com/genshinsim/gsim/internal/characters/klee"
	_ "github.com/genshinsim/gsim/internal/characters/lisa"
	_ "github.com/genshinsim/gsim/internal/characters/ningguang"
	_ "github.com/genshinsim/gsim/internal/characters/noelle"
	_ "github.com/genshinsim/gsim/internal/characters/raiden"
	_ "github.com/genshinsim/gsim/internal/characters/sucrose"
	_ "github.com/genshinsim/gsim/internal/characters/xiangling"
	_ "github.com/genshinsim/gsim/internal/characters/xingqiu"
	_ "github.com/genshinsim/gsim/internal/characters/yoimiya"

	//artifacts
	_ "github.com/genshinsim/gsim/internal/artifacts/archaic"
	_ "github.com/genshinsim/gsim/internal/artifacts/blizzard"
	_ "github.com/genshinsim/gsim/internal/artifacts/bloodstained"
	_ "github.com/genshinsim/gsim/internal/artifacts/bolide"
	_ "github.com/genshinsim/gsim/internal/artifacts/crimson"
	_ "github.com/genshinsim/gsim/internal/artifacts/gladiator"
	_ "github.com/genshinsim/gsim/internal/artifacts/heartofdepth"
	_ "github.com/genshinsim/gsim/internal/artifacts/lavawalker"
	_ "github.com/genshinsim/gsim/internal/artifacts/maiden"
	_ "github.com/genshinsim/gsim/internal/artifacts/noblesse"
	_ "github.com/genshinsim/gsim/internal/artifacts/paleflame"
	_ "github.com/genshinsim/gsim/internal/artifacts/reminiscence"
	_ "github.com/genshinsim/gsim/internal/artifacts/seal"
	_ "github.com/genshinsim/gsim/internal/artifacts/tenacity"
	_ "github.com/genshinsim/gsim/internal/artifacts/thunderingfury"
	_ "github.com/genshinsim/gsim/internal/artifacts/viridescent"
	_ "github.com/genshinsim/gsim/internal/artifacts/wanderer"

	//weapons
	_ "github.com/genshinsim/gsim/internal/weapons/common/blackcliff"
	_ "github.com/genshinsim/gsim/internal/weapons/common/favonius"
	_ "github.com/genshinsim/gsim/internal/weapons/common/generic"
	_ "github.com/genshinsim/gsim/internal/weapons/common/lithic"
	_ "github.com/genshinsim/gsim/internal/weapons/common/royal"
	_ "github.com/genshinsim/gsim/internal/weapons/common/sacrificial"

	_ "github.com/genshinsim/gsim/internal/weapons/bow/alley"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/amos"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/elegy"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/iram"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/prototype"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/rust"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/sharpshooter"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/skyward"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/stringless"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/thundering"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/viridescent"
	_ "github.com/genshinsim/gsim/internal/weapons/bow/windblume"

	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/dodoco"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/frostbearer"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/magicguide"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/mappa"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/memory"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/perception"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/prayer"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/prototype"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/skyward"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/solar"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/thrilling"
	_ "github.com/genshinsim/gsim/internal/weapons/catalyst/widsith"

	_ "github.com/genshinsim/gsim/internal/weapons/claymore/bell"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/pines"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/prototype"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/rainslasher"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/skyrider"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/skyward"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/spine"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/starsilver"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/unforged"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/whiteblind"
	_ "github.com/genshinsim/gsim/internal/weapons/claymore/wolf"

	_ "github.com/genshinsim/gsim/internal/weapons/spear/catch"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/crescent"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/deathmatch"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/dragonbane"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/dragonspine"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/grasscutter"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/homa"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/kitain"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/primordial"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/prototype"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/skyward"
	_ "github.com/genshinsim/gsim/internal/weapons/spear/vortex"

	_ "github.com/genshinsim/gsim/internal/weapons/sword/alley"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/amenoma"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/aquila"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/black"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/festering"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/freedom"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/harbinger"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/ironsting"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/lion"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/mistsplitter"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/primordial"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/prototype"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/skyrider"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/skyward"
	_ "github.com/genshinsim/gsim/internal/weapons/sword/summit"
)
