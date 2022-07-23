package levels

import (
    "github.com/hajimehoshi/ebiten/v2"
//    "github.com/jsnider-mtu/projectx/levels/lvlimages"
)

//var (
//    first *Level = &Level{Max: [2]int{
//)

type Level struct {
    Max [2]int
    Pos [2]int
    Boxes [][4]int
    Doors []*Door
    Image *ebiten.Image
}

type Door struct {
    Coords [4]int
    Direction string // "up", "down", "left", "right"
    Image *ebiten.Image
}
