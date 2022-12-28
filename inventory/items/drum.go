package items

import "github.com/jsnider-mtu/quailgame/player"

type Drum struct {
}

func (d Drum) Slot() string {
    return "BothHands"
}

func (d Drum) Use(p *player.Player) {
}

func (d Drum) Save() string {
    return "Drum"
}

func (d Drum) PrettyPrint() string {
    return "Drum"
}

func (d Drum) Function() string {
    return "instrument"
}

func (d Drum) Damage() (int, int, string) {
    return 0, 0, ""
}
