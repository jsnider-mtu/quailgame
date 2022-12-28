package items

import "github.com/jsnider-mtu/quailgame/player"

type Flute struct {
}

func (f Flute) Slot() string {
    return "BothHands"
}

func (f Flute) Use(p *player.Player) {
}

func (f Flute) Save() string {
    return "Flute"
}

func (f Flute) PrettyPrint() string {
    return "Flute"
}

func (f Flute) Function() string {
    return "instrument"
}

func (f Flute) Damage() (int, int, string) {
    return 0, 0, ""
}
