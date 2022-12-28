package items

import "github.com/jsnider-mtu/quailgame/player"

type HeavyCrossbow struct {
}

func (h HeavyCrossbow) Slot() string {
    return "BothHands"
}

func (h HeavyCrossbow) Use(p *player.Player) {
    // must be equipped to use
}

func (h HeavyCrossbow) Save() string {
    return "HeavyCrossbow"
}

func (h HeavyCrossbow) PrettyPrint() string {
    return "Heavy Crossbow"
}

func (h HeavyCrossbow) Function() string {
    return "range"
}

func (h HeavyCrossbow) Damage() (int, int, string) {
    return 1, 10, "piercing"
}
