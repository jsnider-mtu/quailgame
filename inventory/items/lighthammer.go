package items

import "github.com/jsnider-mtu/quailgame/player"

type LightHammer struct {
}

func (l LightHammer) Slot() string {
    return "RightHand"
}

func (l LightHammer) Use(p *player.Player) {
    // must be equipped to use
}

func (l LightHammer) Save() string {
    return "LightHammer"
}

func (l LightHammer) PrettyPrint() string {
    return "Light Hammer"
}

func (l LightHammer) Function() string {
    return "melee"
}

func (l LightHammer) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}
