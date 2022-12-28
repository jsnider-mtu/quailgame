package items

import "github.com/jsnider-mtu/quailgame/player"

type ThievesTools struct {
}

func (t ThievesTools) Slot() string {
    return ""
}

func (t ThievesTools) Use(p *player.Player) {
}

func (t ThievesTools) Save() string {
    return "ThievesTools"
}

func (t ThievesTools) PrettyPrint() string {
    return "Thieves Tools"
}

func (t ThievesTools) Function() string {
    return "theft"
}

func (t ThievesTools) Damage() (int, int, string) {
    return 0, 0, ""
}
