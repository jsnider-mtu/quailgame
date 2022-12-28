package items

import "github.com/jsnider-mtu/quailgame/player"

type Quarterstaff struct {
}

func (q Quarterstaff) Slot() string {
    return "BothHands"
}

func (q Quarterstaff) Use(p *player.Player) {
    // must be equipped to use
}

func (q Quarterstaff) Save() string {
    return "Quarterstaff"
}

func (q Quarterstaff) PrettyPrint() string {
    return "Quarterstaff"
}

func (q Quarterstaff) Function() string {
    return "melee"
}

func (q Quarterstaff) Damage() (int, int, string) {
    return 1, 6, "bludgeoning"
}
