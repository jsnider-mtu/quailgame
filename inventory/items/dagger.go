package items

import "github.com/jsnider-mtu/quailgame/player"

type Dagger struct {
}

func (d Dagger) Slot() string {
    return "RightHand"
}

func (d Dagger) Use(p *player.Player) {
    // must be equipped to use
}

func (d Dagger) Save() string {
    return "Dagger"
}

func (d Dagger) PrettyPrint() string {
    return "Dagger"
}

func (d Dagger) Function() string {
    return "melee"
}

func (d Dagger) Damage() (int, int, string) {
    return 1, 4, "piercing"
}
