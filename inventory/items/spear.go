package items

import "github.com/jsnider-mtu/quailgame/player"

type Spear struct {
}

func (s Spear) Slot() string {
    return "BothHands"
}

func (s Spear) Use(p *player.Player) {
    // must be equipped to use
}

func (s Spear) Save() string {
    return "Spear"
}

func (s Spear) PrettyPrint() string {
    return "Spear"
}

func (s Spear) Function() string {
    return "melee"
}

func (s Spear) Damage() (int, int, string) {
    return 1, 6, "piercing"
}
