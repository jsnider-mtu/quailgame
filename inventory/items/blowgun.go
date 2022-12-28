package items

import "github.com/jsnider-mtu/quailgame/player"

type Blowgun struct {
}

func (b Blowgun) Slot() string {
    return "RightHand"
}

func (b Blowgun) Use(p *player.Player) {
    // must be equipped to use
}

func (b Blowgun) Save() string {
    return "Blowgun"
}

func (b Blowgun) PrettyPrint() string {
    return "Blowgun"
}

func (b Blowgun) Function() string {
    return "range"
}

func (b Blowgun) Damage() (int, int, string) {
    return 1, 1, "piercing"
}
