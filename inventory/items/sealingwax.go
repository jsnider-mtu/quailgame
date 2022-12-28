package items

import "github.com/jsnider-mtu/quailgame/player"

type SealingWax struct {
}

func (s SealingWax) Slot() string {
    return ""
}

func (s SealingWax) Use(p *player.Player) {
}

func (s SealingWax) Save() string {
    return "SealingWax"
}

func (s SealingWax) PrettyPrint() string {
    return "Sealing Wax"
}

func (s SealingWax) Function() string {
    return "writing"
}

func (s SealingWax) Damage() (int, int, string) {
    return 0, 0, ""
}
