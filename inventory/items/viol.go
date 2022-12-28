package items

import "github.com/jsnider-mtu/quailgame/player"

type Viol struct {
}

func (v Viol) Slot() string {
    return "BothHands"
}

func (v Viol) Use(p *player.Player) {
}

func (v Viol) Save() string {
    return "Viol"
}

func (v Viol) PrettyPrint() string {
    return "Viol"
}

func (v Viol) Function() string {
    return "instrument"
}

func (v Viol) Damage() (int, int, string) {
    return 0, 0, ""
}
