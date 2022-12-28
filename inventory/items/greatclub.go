package items

import "github.com/jsnider-mtu/quailgame/player"

type Greatclub struct {
}

func (g Greatclub) Slot() string {
    return "BothHands"
}

func (g Greatclub) Use(p *player.Player) {
    // must be equipped to use
}

func (g Greatclub) Save() string {
    return "Greatclub"
}

func (g Greatclub) PrettyPrint() string {
    return "Greatclub"
}

func (g Greatclub) Function() string {
    return "melee"
}

func (g Greatclub) Damage() (int, int, string) {
    return 1, 8, "bludgeoning"
}
