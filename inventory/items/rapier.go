package items

import "github.com/jsnider-mtu/quailgame/player"

type Rapier struct {
}

func (r Rapier) Slot() string {
    return "RightHand"
}

func (r Rapier) Use(p *player.Player) {
    // must be equipped to use
}

func (r Rapier) Save() string {
    return "Rapier"
}

func (r Rapier) PrettyPrint() string {
    return "Rapier"
}

func (r Rapier) Function() string {
    return "melee"
}

func (r Rapier) Damage() (int, int, string) {
    return 1, 8, "piercing"
}
