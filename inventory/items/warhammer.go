package items

import "github.com/jsnider-mtu/quailgame/player"

type Warhammer struct {
}

func (w Warhammer) Slot() string {
    return "BothHands"
}

func (w Warhammer) Use(p *player.Player) {
    // must be equipped to use
}

func (w Warhammer) Save() string {
    return "Warhammer"
}

func (w Warhammer) PrettyPrint() string {
    return "Warhammer"
}

func (w Warhammer) Function() string {
    return "melee"
}

func (w Warhammer) Damage() (int, int, string) {
    return 1, 8, "bludgeoning"
}
