package levels

import (
    "bytes"
    "fmt"
    "image"
//    _ "image/jpeg"
    _ "image/png"
    "log"
//    "math/rand"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/levels/lvlimages"
    "github.com/jsnider-mtu/quailgame/npcs"
//    "github.com/jsnider-mtu/quailgame/npcs/npcimages"
//    "github.com/jsnider-mtu/quailgame/player"
)

func LvlTwo(entrance int) *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.Two_PNG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)

    lvldoors := []*Door{
        &Door{Coords: [2]int{48, 96}, NewLvl: [2]int{1, 1}},
        &Door{Coords: [2]int{144, 0}, NewLvl: [2]int{3, 1}},
        &Door{Coords: [2]int{240, 96}, NewLvl: [2]int{1, 1}}}

    var pos [2]int

    switch entrance {
    case 0:
        pos = [2]int{-48, -144}
    case 1:
        pos = [2]int{-96, -192}
    case 2:
        pos = [2]int{-144, 0}
    default:
        log.Fatal(fmt.Sprintf("Entrace %d does not exist", entrance))
    }

    return &Level{
        Name: "Two", Cutscene: 1, Max: [2]int{312, 216}, Pos: pos, Boxes: [][4]int{
            {0, 0, 144, 96},
            {0, 96, 48, 240},
            {96, 96, 144, 192},
            {192, 0, 336, 96},
            {192, 96, 240, 144},
            {288, 96, 336, 240},
            {192, 192, 288, 240}}, Doors: lvldoors, NPCs: []*npcs.NPC{}, Image: lvlImg, Anim: func(a *ebiten.Image, l *Level, b, c, d int) {}}
}
