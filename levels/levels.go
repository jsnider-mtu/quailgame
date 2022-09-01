package levels

import (
    "bytes"
    "fmt"
    "image"
    _ "image/png"
    "log"

    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/npcs"
)

var (
    grassImage *ebiten.Image
)

func init() {
    grassimg, _, err := image.Decode(bytes.NewReader(assets.Grass_Anim_PNG))
    if err != nil {
        log.Fatal(err)
    }
    grassImage = ebiten.NewImageFromImage(grassimg)
}

type Door struct {
    coords [2]int
    NewLvl []interface{}
}

func (d *Door) GetCoords() [2]int {
    return d.coords
}

type Level struct {
    name string
    Cutscene int
    max [2]int
    Pos [2]int
    Boxes [][4]int
    Doors []*Door
    NPCs []*npcs.NPC
    grasses [][2]int
    Image *ebiten.Image
    Anim func(*ebiten.Image, *Level, int, int, int) // screen, l, count, w, h
}

func (l *Level) GetName() string {
    return l.name
}

func (l *Level) GetMax() [2]int {
    return l.max
}

func (l *Level) GetGrasses() [][2]int {
    return l.grasses
}

func (l *Level) GrassAnim(screen *ebiten.Image, w, h int) {
    for _, g := range l.grasses {
        if l.Pos[0] == -g[0] && l.Pos[1] == -g[1] {
            gm := ebiten.GeoM{}
            gm.Translate(float64(w / 2), float64(h / 2))
            screen.DrawImage(grassImage, &ebiten.DrawImageOptions{GeoM: gm})
        }
    }
}

func LoadLvl(newlvl ...interface{}) *Level {
    if len(newlvl) != 4 && len(newlvl) != 2 {
        log.Fatal("Incorrect number of arguments passed to levels.LoadLvl; should be 2 or 4, got %d", len(newlvl))
        return nil
    }
    switch newlvl[0] {
    case "One":
        l := lvlOne(newlvl[1].(int))
        if len(newlvl) == 4 {
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
        }
        return l
    case "Two":
        l := lvlTwo(newlvl[1].(int))
        if len(newlvl) == 4 {
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
        }
        return l
    case "VerticalWall":
        l := verticalWallLvl(newlvl[1].(int))
        if len(newlvl) == 4 {
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
        }
        return l
    default:
        log.Fatal(fmt.Sprintf("Level %s does not exist", newlvl[0]))
    }
    return lvlOne(newlvl[1].(int))
}
