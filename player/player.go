package player

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/inventory"
)

type Player struct {
    Pos [2]int
    Inv *inventory.Inv
    Image *ebiten.Image
}
