package items

import "github.com/jsnider-mtu/quailgame/player"

type Javelin struct {
}

func (j Javelin) Slot() string {
    return "RightHand"
}

func (j Javelin) Use(p *player.Player) {
    // must be equipped to use
}

func (j Javelin) Save() string {
    return "Javelin"
}

func (j Javelin) PrettyPrint() string {
    return "Javelin"
}

func (j Javelin) Function() string {
    return "melee"
}

func (j Javelin) Damage() (int, int, string) {
    return 1, 6, "piercing"
}
