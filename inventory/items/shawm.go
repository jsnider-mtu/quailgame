package items

import "github.com/jsnider-mtu/quailgame/player"

type Shawm struct {
}

func (s Shawm) Slot() string {
    return "BothHands"
}

func (s Shawm) Use(p *player.Player) {
}

func (s Shawm) Save() string {
    return "Shawm"
}

func (s Shawm) PrettyPrint() string {
    return "Shawm"
}

func (s Shawm) Function() string {
    return "instrument"
}

func (s Shawm) Damage() (int, int, string) {
    return 0, 0, ""
}
