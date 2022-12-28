package items

import "github.com/jsnider-mtu/quailgame/player"

type Shortsword struct {
}

func (s Shortsword) Slot() string {
    return "RightHand"
}

func (s Shortsword) Use(p *player.Player) {
    // must be equipped to use
}

func (s Shortsword) Save() string {
    return "Shortsword"
}

func (s Shortsword) PrettyPrint() string {
    return "Shortsword"
}

func (s Shortsword) Function() string {
    return "melee"
}

func (s Shortsword) Damage() (int, int, string) {
    return 1, 6, "piercing"
}
