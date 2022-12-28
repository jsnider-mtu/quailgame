package items

import "github.com/jsnider-mtu/quailgame/player"

type Scimitar struct {
}

func (s Scimitar) Slot() string {
    return "RightHand"
}

func (s Scimitar) Use(p *player.Player) {
    // must be equipped to use
}

func (s Scimitar) Save() string {
    return "Scimitar"
}

func (s Scimitar) PrettyPrint() string {
    return "Scimitar"
}

func (s Scimitar) Function() string {
    return "melee"
}

func (s Scimitar) Damage() (int, int, string) {
    return 1, 6, "slashing"
}
