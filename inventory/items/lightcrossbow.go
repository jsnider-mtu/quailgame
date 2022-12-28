package items

import "github.com/jsnider-mtu/quailgame/player"

type LightCrossbow struct {
}

func (l LightCrossbow) Slot() string {
    return "BothHands"
}

func (l LightCrossbow) Use(p *player.Player) {
    // must be equipped to use
}

func (l LightCrossbow) Save() string {
    return "LightCrossbow"
}

func (l LightCrossbow) PrettyPrint() string {
    return "Light Crossbow"
}

func (l LightCrossbow) Function() string {
    return "range"
}

func (l LightCrossbow) Damage() (int, int, string) {
    return 1, 8, "piercing"
}
