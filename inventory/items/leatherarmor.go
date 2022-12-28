package items

import "github.com/jsnider-mtu/quailgame/player"

type LeatherArmor struct {
}

func (l LeatherArmor) Slot() string {
    return "Armor"
}

func (l LeatherArmor) Use(p *player.Player) {
}

func (l LeatherArmor) Save() string {
    return "LeatherArmor"
}

func (l LeatherArmor) PrettyPrint() string {
    return "Leather Armor"
}

func (l LeatherArmor) Function() string {
    return "armor"
}

func (l LeatherArmor) Damage() (int, int, string) {
    return 0, 0, ""
}
