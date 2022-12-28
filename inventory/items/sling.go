package items

import "github.com/jsnider-mtu/quailgame/player"

type Sling struct {
}

func (s Sling) Slot() string {
    return "RightHand"
}

func (s Sling) Use(p *player.Player) {
    // must be equipped to use
}

func (s Sling) Save() string {
    return "Sling"
}

func (s Sling) PrettyPrint() string {
    return "Sling"
}

func (s Sling) Function() string {
    return "range"
}

func (s Sling) Damage() (int, int, string) {
    return 1, 4, "bludgeoning"
}
