package lvlone

import (
    "bytes"
    "image"
    "image/color"
    _ "image/jpeg"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/projectx/levels"
    "github.com/jsnider-mtu/projectx/levels/lvlimages"
)

func Setup() *levels.Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.One_JPEG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)
    
    lvldoors := []*levels.Door{&levels.Door{Coords: [4]int{96, 96, 192, 144}, Direction: "up", Image: ebiten.NewImage(192-96, 144-96)}}
    for _, ld := range lvldoors {
        ld.Image.Fill(color.Black)
    }
    
    return &levels.Level{Max: [2]int{1152, 504}, Pos: [2]int{384, 288}, Boxes: [][4]int{{-384, -288, 1154-384, 172-288}}, Doors: lvldoors, Image: lvlImg}
}
