package items

import "github.com/jsnider-mtu/quailgame/player"

type Perfume struct {
}

func (p Perfume) Slot() string {
    return ""
}

func (p Perfume) Use(p *player.Player) {
}

func (p Perfume) Save() string {
    return "Perfume"
}

func (p Perfume) PrettyPrint() string {
    return "Perfume"
}

func (p Perfume) Function() string {
    return "disguise"
}

func (p Perfume) Damage() (int, int, string) {
    return 0, 0, ""
}
