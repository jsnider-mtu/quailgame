package levels

import (
    "fmt"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/npcs"
)

type Door struct {
    Coords [2]int
    NewLvl [2]int
}

type Level struct {
    Name string
    Cutscene int
    Max [2]int
    Pos [2]int
    Boxes [][4]int
    Doors []*Door
    NPCs []*npcs.NPC
    Image *ebiten.Image
    Anim func(*ebiten.Image, *Level, int, int, int)
}

func LoadLvl(name string, x, y int) *Level {
    switch name {
    case "One":
        l := LvlOne(0)
        l.Pos = [2]int{x, y}
        return l
    case "Two":
        l := LvlTwo(0)
        l.Pos = [2]int{x, y}
        return l
    case "VerticalWall":
        l := VerticalWallLvl(0)
        l.Pos = [2]int{x, y}
        return l
    default:
        log.Fatal(fmt.Sprintf("Level %s does not exist"))
    }
    return LvlOne(0)
}
