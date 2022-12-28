package items

import "github.com/jsnider-mtu/quailgame/player"

type Longbow struct {
}

func (l Longbow) Slot() string {
    return "BothHands"
}

func (l Longbow) Use(p *player.Player) {
    // must be equipped to use
}

func (l Longbow) Save() string {
    return "Longbow"
}

func (l Longbow) PrettyPrint() string {
    return "Longbow"
}

func (l Longbow) Function() string {
    return "range"
}

func (l Longbow) Damage() (int, int, string) {
    return 1, 8, "piercing"
}
