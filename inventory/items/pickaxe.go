package items

import (
    "github.com/jsnider-mtu/quailgame/inventory"
)

type Pickaxe struct {
    Material string
    Durability int
}

func (p *Pickaxe) Use() {
}
