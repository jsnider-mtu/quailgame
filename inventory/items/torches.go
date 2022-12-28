package items

import (
    "fmt"
    "strconv"

    import "github.com/jsnider-mtu/quailgame/player"
)

type Torches struct {
    Quantity int
}

func (t Torches) Slot() string {
    return "LeftHand"
}

func (t Torches) Use(p *player.Player) {
}

func (t Torches) Save() string {
    return "Torches," + strconv.Itoa(t.Quantity)
}

func (t Torches) PrettyPrint() string {
    return fmt.Sprintf("Torches (%d)", t.Quantity)
}

func (t Torches) Function() string {
    return "light"
}

func (t Torches) Damage() (int, int, string) {
    return 0, 0, ""
}
