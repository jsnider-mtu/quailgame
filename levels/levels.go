package levels

import (
    "fmt"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/npcs"
)

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
    Image *ebiten.Image
    Anim func(*ebiten.Image, *Level, int, int, int)
}

func (l *Level) GetName() string {
    return l.name
}

func (l *Level) GetMax() [2]int {
    return l.max
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
