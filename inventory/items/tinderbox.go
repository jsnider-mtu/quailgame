package items

import "github.com/jsnider-mtu/quailgame/player"

type Tinderbox struct {
}

func (t Tinderbox) Slot() string {
    return ""
}

func (t Tinderbox) Use(p *player.Player) {
}

func (t Tinderbox) Save() string {
    return "Tinderbox"
}

func (t Tinderbox) PrettyPrint() string {
    return "Tinderbox"
}

func (t Tinderbox) Function() string {
    return "fire"
}

func (t Tinderbox) Damage() (int, int, string) {
    return 0, 0, ""
}
