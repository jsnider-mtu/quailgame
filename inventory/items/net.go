package items

import "github.com/jsnider-mtu/quailgame/player"

type Net struct {
}

func (n Net) Slot() string {
    return "RightHand"
}

func (n Net) Use(p *player.Player) {
    // must be equipped to use
}

func (n Net) Save() string {
    return "Net"
}

func (n Net) PrettyPrint() string {
    return "Net"
}

func (n Net) Function() string {
    return "range"
}

func (n Net) Damage() (int, int, string) {
    return 0, 0, ""
}
