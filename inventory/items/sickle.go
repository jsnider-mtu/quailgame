package items

import "github.com/jsnider-mtu/quailgame/player"

type Sickle struct {
}

func (s Sickle) Slot() string {
    return "RightHand"
}

func (s Sickle) Use(p *player.Player) {
    // must be equipped to use
}

func (s Sickle) Save() string {
    return "Sickle"
}

func (s Sickle) PrettyPrint() string {
    return "Sickle"
}

func (s Sickle) Function() string {
    return "melee"
}

func (s Sickle) Damage() (int, int, string) {
    return 1, 4, "slashing"
}
