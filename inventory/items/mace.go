package items

import "github.com/jsnider-mtu/quailgame/player"

type Mace struct {
}

func (m Mace) Slot() string {
    return "RightHand"
}

func (m Mace) Use(p *player.Player) {
    // must be equipped to use
}

func (m Mace) Save() string {
    return "Mace"
}

func (m Mace) PrettyPrint() string {
    return "Mace"
}

func (m Mace) Function() string {
    return "melee"
}

func (m Mace) Damage() (int, int, string) {
    return 1, 6, "bludgeoning"
}
