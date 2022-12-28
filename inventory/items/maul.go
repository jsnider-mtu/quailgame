package items

import "github.com/jsnider-mtu/quailgame/player"

type Maul struct {
}

func (m Maul) Slot() string {
    return "BothHands"
}

func (m Maul) Use(p *player.Player) {
    // must be equipped to use
}

func (m Maul) Save() string {
    return "Maul"
}

func (m Maul) PrettyPrint() string {
    return "Maul"
}

func (m Maul) Function() string {
    return "melee"
}

func (m Maul) Damage() (int, int, string) {
    return 2, 6, "bludgeoning"
}
