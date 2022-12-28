package items

import "github.com/jsnider-mtu/quailgame/player"

type Greatsword struct {
}

func (g Greatsword) Slot() string {
    return "BothHands"
}

func (g Greatsword) Use(p *player.Player) {
    // must be equipped to use
}

func (g Greatsword) Save() string {
    return "Greatsword"
}

func (g Greatsword) PrettyPrint() string {
    return "Greatsword"
}

func (g Greatsword) Function() string {
    return "melee"
}

func (g Greatsword) Damage() (int, int, string) {
    return 2, 6, "slashing"
}
