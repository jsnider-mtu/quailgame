package items

import "github.com/jsnider-mtu/quailgame/player"

type ArcaneFocus struct {
}

func (a ArcaneFocus) Slot() string {
    return "LeftHand"
}

func (a ArcaneFocus) Use(p *player.Player) {
}

func (a ArcaneFocus) Save() string {
    return "ArcaneFocus"
}

func (a ArcaneFocus) PrettyPrint() string {
    return "Arcane Focus" 
}

func (a ArcaneFocus) Function() string {
    return "spells"
}

func (a ArcaneFocus) Damage() (int, int, string) {
    return 0, 0, ""
}
