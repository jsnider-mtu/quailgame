package items

import "github.com/jsnider-mtu/quailgame/player"

type Greataxe struct {
}

func (g Greataxe) Slot() string {
    return "BothHands"
}

func (g Greataxe) Use(p *player.Player) {
    // must be equipped to use
}

func (g Greataxe) Save() string {
    return "Greataxe"
}

func (g Greataxe) PrettyPrint() string {
    return "Greataxe"
}

func (g Greataxe) Function() string {
    return "melee"
}

func (g Greataxe) Damage() (int, int, string) {
    return 1, 12, "slashing"
}
