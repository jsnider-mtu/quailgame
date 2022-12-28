package items

import "github.com/jsnider-mtu/quailgame/player"

type Morningstar struct {
}

func (m Morningstar) Slot() string {
    return "RightHand"
}

func (m Morningstar) Use(p *player.Player) {
    // must be equipped to use
}

func (m Morningstar) Save() string {
    return "Morningstar"
}

func (m Morningstar) PrettyPrint() string {
    return "Morningstar"
}

func (m Morningstar) Function() string {
    return "melee"
}

func (m Morningstar) Damage() (int, int, string) {
    return 1, 8, "piercing"
}
