package items

import "github.com/jsnider-mtu/quailgame/player"

type Dulcimer struct {
}

func (d Dulcimer) Slot() string {
    return "BothHands"
}

func (d Dulcimer) Use(p *player.Player) {
}

func (d Dulcimer) Save() string {
    return "Dulcimer"
}

func (d Dulcimer) PrettyPrint() string {
    return "Dulcimer"
}

func (d Dulcimer) Function() string {
    return "instrument"
}

func (d Dulcimer) Damage() (int, int, string) {
    return 0, 0, ""
}
