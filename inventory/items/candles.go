package items

import (
    "fmt"
    "strconv"

    "github.com/jsnider-mtu/quailgame/player"
)

type Candles struct {
    Quantity int
}

func (c Candles) Slot() string {
    return "LeftHand"
}

func (c Candles) Use(p *player.Player) {
    // illuminate surroundings
    if c.Quantity > 0 {
        p.Light = [3]int{5, 5, 600}
    }
}

func (c Candles) Save() string {
    return "Candles," + strconv.Itoa(c.Quantity)
}

func (c Candles) PrettyPrint() string {
    return fmt.Sprintf("Candles (%d)", c.Quantity)
}

func (c Candles) Function() string {
    return "light"
}

func (c Candles) Damage() (int, int, string) {
    return 0, 0, ""
}
