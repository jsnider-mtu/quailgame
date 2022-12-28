package items

import "github.com/jsnider-mtu/quailgame/player"

type Lute struct {
}

func (l Lute) Slot() string {
    return "BothHands"
}

func (l Lute) Use(p *player.Player) {
}

func (l Lute) Save() string {
    return "Lute"
}

func (l Lute) PrettyPrint() string {
    return "Lute"
}

func (l Lute) Function() string {
    return "instrument"
}

func (l Lute) Damage() (int, int, string) {
    return 0, 0, ""
}
