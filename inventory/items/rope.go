package items

import (
    "fmt"
    "strconv"

    import "github.com/jsnider-mtu/quailgame/player"
)

type Rope struct {
    Length int
}

func (r Rope) Slot() string {
    return "BothHands"
}

func (r Rope) Use(p *player.Player) {
}

func (r Rope) Save() string {
    return "Rope," + strconv.Itoa(r.Length)
}

func (r Rope) PrettyPrint() string {
    return fmt.Sprintf("Rope (%d)", r.Length)
}

func (r Rope) Function() string {
    return "climbing"
}

func (r Rope) Damage() (int, int, string) {
    return 0, 0, ""
}
