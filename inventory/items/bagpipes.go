package items

import "github.com/jsnider-mtu/quailgame/player"

type Bagpipes struct {
}

func (b Bagpipes) Slot() string {
    return "BothHands"
}

func (b Bagpipes) Use(p *player.Player) {
}

func (b Bagpipes) Save() string {
    return "Bagpipes"
}

func (b Bagpipes) PrettyPrint() string {
    return "Bagpipes" 
}

func (b Bagpipes) Function() string {
    return "instrument"
}

func (b Bagpipes) Damage() (int, int, string) {
    return 0, 0, ""
}
