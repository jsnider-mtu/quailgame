package items

import "github.com/jsnider-mtu/quailgame/player"

type PanFlute struct {
}

func (p PanFlute) Slot() string {
    return "BothHands"
}

func (p PanFlute) Use(p *player.Player) {
}

func (p PanFlute) Save() string {
    return "PanFlute"
}

func (p PanFlute) PrettyPrint() string {
    return "Pan Flute"
}

func (p PanFlute) Function() string {
    return "instrument"
}

func (p PanFlute) Damage() (int, int, string) {
    return 0, 0, ""
}
