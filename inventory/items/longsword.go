package items

import "github.com/jsnider-mtu/quailgame/player"

type Longsword struct {
}

func (l Longsword) Slot() string {
    return "RightHand"
}

func (l Longsword) Use(p *player.Player) {
    // must be equipped to use
}

func (l Longsword) Save() string {
    return "Longsword"
}

func (l Longsword) PrettyPrint() string {
    return "Longsword"
}

func (l Longsword) Function() string {
    return "melee"
}

func (l Longsword) Damage() (int, int, string) {
    return 1, 8, "slashing"
}
