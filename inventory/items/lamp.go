package items

import "github.com/jsnider-mtu/quailgame/player"

type Lamp struct {
}

func (l Lamp) Slot() string {
    return "LeftHand"
}

func (l Lamp) Use(p *player.Player) {
}

func (l Lamp) Save() string {
    return "Lamp"
}

func (l Lamp) PrettyPrint() string {
    return "Lamp"
}

func (l Lamp) Function() string {
    return "light"
}

func (l Lamp) Damage() (int, int, string) {
    return 0, 0, ""
}
