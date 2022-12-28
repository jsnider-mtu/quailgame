package items

import "github.com/jsnider-mtu/quailgame/player"

type HandCrossbow struct {
}

func (h HandCrossbow) Slot() string {
    return "RightHand"
}

func (h HandCrossbow) Use(p *player.Player) {
    // must be equipped to use
}

func (h HandCrossbow) Save() string {
    return "HandCrossbow"
}

func (h HandCrossbow) PrettyPrint() string {
    return "Hand Crossbow"
}

func (h HandCrossbow) Function() string {
    return "range"
}

func (h HandCrossbow) Damage() (int, int, string) {
    return 1, 6, "piercing"
}
