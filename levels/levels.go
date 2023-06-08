package levels

import (
    "bytes"
    "fmt"
    "image"
    _ "image/png"
    "log"
    "math"

    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/npcs"
    "github.com/jsnider-mtu/quailgame/player"
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
    oils [][3]int
}

func (l *Level) OilSpot(spot [2]int) {
    for oi, oil := range l.oils {
        if [2]int{oil[0], oil[1]} == spot {
            l.oils[oi][2] = 10
            return
        }
    }
    l.oils = append(l.oils, [3]int{spot[0], spot[1], 10})
    return
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

func (l *Level) GetOils() [][3]int {
    return l.oils
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
        log.Fatal("Incorrect number of arguments passed to levels.LoadLvl; should be 2 or 4, got", len(newlvl))
        return nil
    }
    fmt.Println("Level name is", newlvl[0])
    switch newlvl[0] {
    case "One":
        var l *Level
        l = lvlOne(newlvl[1].(int))
        if len(newlvl) == 4 {
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
        }
        return l
    case "Two":
        var l *Level
        l = lvlTwo(newlvl[1].(int))
        if len(newlvl) == 4 {
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
        }
        return l
    case "VerticalWall":
        var l *Level
        l = verticalWallLvl(newlvl[1].(int))
        if len(newlvl) == 4 {
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
        }
        return l
    default:
        log.Fatal(fmt.Sprintf("Level %s does not exist", newlvl[0]))
    }
    return lvlOne(newlvl[1].(int))
}

func (l *Level) LineOfSight(p *player.Player, target [2]int) (bool, bool, float64) {
    var slope float64
    var slopevert bool = false
    if target[0] > p.Pos[0] {
        slope = float64((target[1] + 24) - (p.Pos[1] + 24)) / float64((target[0] + 24) - (p.Pos[0] + 24))
    } else if target[0] < p.Pos[0] {
        slope = float64((p.Pos[1] + 24) - (target[1] + 24)) / float64((p.Pos[0] + 24) - (target[0] + 24))
    } else {
        slopevert = true
    }
    if slopevert {
        if target[1] > p.Pos[1] {
            for _, box := range l.Boxes {
                if p.Pos[0] + 24 > box[0] && p.Pos[0] + 24 < box[2] && p.Pos[1] + 24 < box[1] && target[1] + 24 > box[3] {
                    return false, true, slope
                }
            }
            return true, true, slope
        } else {
            for _, box := range l.Boxes {
                if p.Pos[0] + 24 > box[0] && p.Pos[0] + 24 < box[2] && p.Pos[1] + 24 > box[3] && target[1] + 24 < box[1] {
                    return false, true, slope
                }
            }
            return true, true, slope
        }
    } else {
        if target[0] > p.Pos[0] {
            for x := p.Pos[0] + 24; x <= target[0] + 24; x++ {
                y := int((float64(x - (p.Pos[0] + 24)) * slope) + float64(p.Pos[1] + 24))
                for _, box := range l.Boxes {
                    if x > box[0] && x < box[2] && y > box[1] && y < box[3] {
                        return false, false, slope
                    }
                }
            }
            return true, false, slope
        } else {
            for x := target[0] + 24; x <= p.Pos[0] + 24; x++ {
                y := int((float64(x - (target[0] + 24)) * slope) + float64(target[1] + 24))
                for _, box := range l.Boxes {
                    if x > box[0] && x < box[2] && y > box[1] && y < box[3] {
                        return false, false, slope
                    }
                }
            }
            return true, false, slope
        }
    }
}

func (l *Level) TryUpdatePos(pc bool, p *player.Player, vert bool, dist int, attempt int, mc *player.Player) (bool, string) {
    if vert {
        if p.Pos[1] + dist > -12 && p.Pos[1] + dist < l.GetMax()[1] - 12 {
            if !pc {
                if p.Pos[0] > mc.Pos[0] - 24 && p.Pos[0] < mc.Pos[0] + 24 && p.Pos[1] + dist < mc.Pos[1] + 24 && p.Pos[1] + dist > mc.Pos[1] - 24 {
                    return false, "player"
                }
            }
            for _, a := range l.Boxes {
                if p.Pos[0] > a[0] - 36 && p.Pos[0] < a[2] - 12 && p.Pos[1] + dist < a[3] - 12 && p.Pos[1] + dist > a[1] - 36 {
                    return false, "box"
                }
            }
            for _, b := range l.NPCs {
                if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                    continue
                }
                if p.Pos[0] > b.PC.Pos[0] - 24 && p.Pos[0] < b.PC.Pos[0] + 24 && p.Pos[1] + dist < b.PC.Pos[1] + 24 && p.Pos[1] + dist > b.PC.Pos[1] - 24 {
                    if !pc {
                        return false, "npc"
                    } else {
                        if attempt < 10 {
                            return false, "npc"
                        }
                    }
                }
            }
            p.Pos[1] += dist
            if pc {
                l.Pos[1] -= dist
            }
            return true, ""
        }
        return false, "mapedge"
    } else {
        if p.Pos[0] + dist > -12 && p.Pos[0] + dist < l.GetMax()[0] - 12 {
            if !pc {
                if p.Pos[0] + dist < mc.Pos[0] + 24 && p.Pos[0] + dist > mc.Pos[0] - 24 && p.Pos[1] > mc.Pos[1] - 24 && p.Pos[1] < mc.Pos[1] + 24 {
                    return false, "player"
                }
            }
            for _, a := range l.Boxes {
                if p.Pos[0] + dist > a[0] - 36 && p.Pos[0] + dist < a[2] - 12 && p.Pos[1] > a[1] - 36 && p.Pos[1] < a[3] - 12 {
                    return false, "box"
                }
            }
            for _, b := range l.NPCs {
                if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                    continue
                }
                if p.Pos[0] + dist > b.PC.Pos[0] - 24 && p.Pos[0] + dist < b.PC.Pos[0] + 24 && p.Pos[1] < b.PC.Pos[1] + 24 && p.Pos[1] > b.PC.Pos[1] - 24 {
                    if !pc {
                        return false, "npc"
                    } else {
                        if attempt < 10 {
                            return false, "npc"
                        }
                    }
                }
            }
            p.Pos[0] += dist
            if pc {
                l.Pos[0] -= dist
            }
            return true, ""
        }
        return false, "mapedge"
    }
}

func (l *Level) Distance(p *player.Player, target [2]int) float64 {
    return math.Sqrt(math.Pow(float64(target[0] - p.Pos[0]), 2.0) + math.Pow(float64(target[1] - p.Pos[1]), 2.0))
}
