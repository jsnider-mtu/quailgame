package items

import "github.com/jsnider-mtu/quailgame/player"

type Shortbow struct {
}

func (s Shortbow) Slot() string {
    return "BothHands"
}

func (s Shortbow) Use(p *player.Player) {
    // must be equipped to use
}

func (s Shortbow) Save() string {
    return "Shortbow"
}

func (s Shortbow) PrettyPrint() string {
    return "Shortbow"
}

func (s Shortbow) Function() string {
    return "range"
}

func (s Shortbow) Damage() (int, int, string) {
    return 1, 6, "piercing"
}
