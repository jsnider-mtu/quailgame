package items

import "github.com/jsnider-mtu/quailgame/player"

type Glaive struct {
}

func (g Glaive) Slot() string {
    return "BothHands"
}

func (g Glaive) Use(p *player.Player) {
    // must be equipped to use
}

func (g Glaive) Save() string {
    return "Glaive"
}

func (g Glaive) PrettyPrint() string {
    return "Glaive"
}

func (g Glaive) Function() string {
    return "melee"
}

func (g Glaive) Damage() (int, int, string) {
    return 1, 10, "slashing"
}
