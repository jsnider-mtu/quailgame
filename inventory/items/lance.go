package items

import "github.com/jsnider-mtu/quailgame/player"

type Lance struct {
}

func (l Lance) Slot() string {
    return "RightHand"
}

func (l Lance) Use(p *player.Player) {
    // must be equipped to use
}

func (l Lance) Save() string {
    return "Lance"
}

func (l Lance) PrettyPrint() string {
    return "Lance"
}

func (l Lance) Function() string {
    return "melee"
}

func (l Lance) Damage() (int, int, string) {
    return 1, 12, "piercing"
}
