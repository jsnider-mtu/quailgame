package items

import "github.com/jsnider-mtu/quailgame/player"

type Halberd struct {
}

func (h Halberd) Slot() string {
    return "BothHands"
}

func (h Halberd) Use(p *player.Player) {
    // must be equipped to use
}

func (h Halberd) Save() string {
    return "Halberd"
}

func (h Halberd) PrettyPrint() string {
    return "Halberd"
}

func (h Halberd) Function() string {
    return "melee"
}

func (h Halberd) Damage() (int, int, string) {
    return 1, 10, "slashing"
}
