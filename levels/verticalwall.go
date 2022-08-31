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

func VerticalWallLvl(entrance int) *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.VerticalWallLvlOne_PNG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)
    cloudimg, _, err := image.Decode(bytes.NewReader(lvlimages.Clouds_PNG))
    if err != nil {
        log.Fatal(err)
    }
    cloudImage := ebiten.NewImageFromImage(cloudimg)

    lvldoors := []*Door{}

    for x := 0; x < 15; x++ {
        lvldoors = append(lvldoors, &Door{Coords: [2]int{432 + (24 * x), 288}, NewLvl: []interface{}{"Two", 1}})
    }

    for x := 0; x < 15; x++ {
        lvldoors = append(lvldoors, &Door{Coords: [2]int{432 + (24 * x), 5808}, NewLvl: []interface{}{"Two", 2}})
    }

    var pos [2]int

    switch entrance {
    case 0:
        pos = [2]int{-600, -336}
    case 1:
        pos = [2]int{-600, -5760}
    default:
        log.Fatal(fmt.Sprintf("Entrace %d does not exist", entrance))
    }

    return &Level{
        Name: "VerticalWall", Cutscene: -1, Max:[2]int{816, 5856}, Pos: pos, Boxes: [][4]int{
            {0, 0, 432, 6144},
            {816, 0, 1248, 6144}}, Doors: lvldoors, NPCs: []*npcs.NPC{}, Image: lvlImg,
        Anim: func(screen *ebiten.Image, l *Level, count, w, h int) {
            iw, _ := l.Image.Size()
            gm := ebiten.GeoM{}
            gm.Translate(float64(((iw - 96) + l.Pos[0]) - count), float64(l.Pos[1]))
            screen.DrawImage(cloudImage, &ebiten.DrawImageOptions{GeoM: gm})}}
}
