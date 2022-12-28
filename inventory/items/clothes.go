package items

import "github.com/jsnider-mtu/quailgame/player"

type Clothes struct {
    Quality string
}

func (c Clothes) Slot() string {
    return "Clothes"
}

func (c Clothes) Use(p *player.Player) {
}

func (c Clothes) Save() string {
    return "Clothes," + c.Quality
}

func (c Clothes) PrettyPrint() string {
    if c.Quality == "Costume" {
        return c.Quality
    } else {
        return c.Quality + " Clothes"
    }
}

func (c Clothes) Function() string {
    return "clothes"
}

func (c Clothes) Damage() (int, int, string) {
    return 0, 0, ""
}
