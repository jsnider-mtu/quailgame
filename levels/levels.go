package levels

import (
    "bytes"
    "image"
    "image/color"
    _ "image/jpeg"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/projectx/levels/lvlimages"
)

type Level struct {
    Max [2]int
    Pos [2]int
    Boxes [][4]int
    Doors []*Door
    Image *ebiten.Image
}

type Door struct {
    Coords [2]int
    NewLvl int
    Image *ebiten.Image
}

func LvlOne() *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.One_JPEG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)

    lvldoors := []*Door{&Door{Coords: [2]int{96, 96}, NewLvl: 2, Image: ebiten.NewImage(48, 48)}}
    for _, ld := range lvldoors {
        ld.Image.Fill(color.Black)
    }

    return &Level{Max: [2]int{720, 528}, Pos: [2]int{-48, -144}, Boxes: [][4]int{{48, 48, 96, 96}}, Doors: lvldoors, Image: lvlImg}
}

func LvlTwo() *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.One_JPEG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)

    lvldoors := []*Door{&Door{Coords: [2]int{192, 192}, NewLvl: 1, Image: ebiten.NewImage(48, 48)}}
    for _, ld := range lvldoors {
        ld.Image.Fill(color.Black)
    }

    return &Level{Max: [2]int{720, 528}, Pos: [2]int{-96, -144}, Boxes: [][4]int{{0, 0, 48, 48}}, Doors: lvldoors, Image: lvlImg}
}
