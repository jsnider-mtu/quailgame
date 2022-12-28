package items

import "github.com/jsnider-mtu/quailgame/player"

type Club struct {
}

func (c Club) Slot() string {
    return "RightHand"
}

func (c Club) Use(p *player.Player) {
    // must be equipped to use
}

func (c Club) Save() string {
    return "Club"
}

func (c Club) PrettyPrint() string {
    return "Club"
}

func (c Club) Function() string {
    return "melee"
}

func (c Club) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}
