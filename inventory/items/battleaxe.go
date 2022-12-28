package items

import "github.com/jsnider-mtu/quailgame/player"

type Battleaxe struct {
}

func (b Battleaxe) Slot() string {
    return "BothHands"
}

func (b Battleaxe) Use(p *player.Player) {
    // must be equipped to use
}

func (b Battleaxe) Save() string {
    return "Battleaxe"
}

func (b Battleaxe) PrettyPrint() string {
    return "Battleaxe"
}

func (b Battleaxe) Function() string {
    return "melee"
}

func (b Battleaxe) Damage() (int, int, string) {
    return 1, 8, "slashing"
}
