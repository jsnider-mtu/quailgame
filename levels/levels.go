package levels

import (
    "bytes"
    "errors"
    "fmt"
    "image"
    _ "image/png"
    "log"
    "math"
    "math/rand"
    "strconv"
    "strings"

    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/npcs"
    "github.com/jsnider-mtu/quailgame/player"
//    "github.com/jsnider-mtu/quailgame/utils"
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

func (l *Level) Attack(p, target *player.Player) (bool, error) {
    // Check equipped weapon, if melee check npc's position before attacking, if range check lineofsight
    if p.Equipment.BothHands != nil {
        if strings.HasPrefix(p.Equipment.BothHands.Function(), "melee") {
            if (p.Pos[0] == target.Pos[0] + 24 && p.Pos[1] == target.Pos[1] + 24) ||
                (p.Pos[0] == target.Pos[0] && p.Pos[1] == target.Pos[1] + 24) ||
                (p.Pos[0] == target.Pos[0] - 24 && p.Pos[1] == target.Pos[1] + 24) ||
                (p.Pos[0] == target.Pos[0] + 24 && p.Pos[1] == target.Pos[1]) ||
                (p.Pos[0] == target.Pos[0] - 24 && p.Pos[1] == target.Pos[1]) ||
                (p.Pos[0] == target.Pos[0] + 24 && p.Pos[1] == target.Pos[1] - 24) ||
                (p.Pos[0] == target.Pos[0] && p.Pos[1] == target.Pos[1] - 24) ||
                (p.Pos[0] == target.Pos[0] - 24 && p.Pos[1] == target.Pos[1] - 24) {
                // Roll to hit (modifiers), if meets or exceeds AC then roll for damage
                hitroll := rand.Intn(20) + 1
                hittotal := hitroll + p.Stats.StrMod
                fmt.Println(fmt.Sprintf("Melee attack roll: %d, plus modifier: %d, equals: %d", hitroll, p.Stats.StrMod, hittotal))
                if hit := hittotal >= target.Stats.AC; hit {
                    fmt.Println(fmt.Sprintf("%d hits against %d", hittotal, target.Stats.AC))
                    damnd, dam, damtype := p.Equipment.BothHands.Damage()
                    damage := 0
                    for x := 0; x < damnd; x++ {
                        newdam := rand.Intn(dam) + 1
                        damage += newdam
                        fmt.Println(fmt.Sprintf("#%d die rolled %d, total damage is now %d", x + 1, newdam, damage))
                    }
                    target.Stats.HP -= damage
                    if target.Stats.Oiled > 0 && damtype == "fire" {
                        target.Stats.HP -= 5
                    }
                    fmt.Println(fmt.Sprintf("%s has lost %d hit points, now has %d hit points remaining", target.GetName(), damage, target.Stats.HP))
                    return true, nil
                } else {
                    fmt.Println(fmt.Sprintf("Melee attack missed, Roll: %d, AC: %d", hittotal, target.Stats.AC))
                    return false, nil
                }
            } else {
                fmt.Println("Enemy is not within melee range")
                return false, errors.New("Enemy is not within melee range")
            }
        } else if strings.HasPrefix(p.Equipment.BothHands.Function(), "range") {
            if l.Distance(p, target.Pos) <= p.Equipment.BothHands.GetRange()[1] {
                if ok, _, _ := l.LineOfSight(p, target.Pos); ok {
                    // Check distance, if ok roll to hit (modifiers), if meets or exceeds AC then roll for damage
                    hitroll := rand.Intn(20) + 1
                    if l.Distance(p, target.Pos) >= p.Equipment.BothHands.GetRange()[0] {
                        hitroll2 := rand.Intn(20) + 1
                        if hitroll2 < hitroll {
                            hitroll = hitroll2
                        }
                    }
                    hittotal := hitroll + p.Stats.DexMod
                    fmt.Println(fmt.Sprintf("Ranged attack roll: %d, plus modifier: %d, equals: %d", hitroll, p.Stats.DexMod, hittotal))
                    if hit := hittotal >= target.Stats.AC; hit {
                        fmt.Println(fmt.Sprintf("%d hits against %d", hittotal, target.Stats.AC))
                        damnd, dam, damtype := p.Equipment.BothHands.Damage()
                        damage := 0
                        for x := 0; x < damnd; x++ {
                            newdam := rand.Intn(dam) + 1
                            damage += newdam
                            fmt.Println(fmt.Sprintf("#%d die rolled %d, total damage is now %d", x + 1, newdam, damage))
                        }
                        target.Stats.HP -= damage
                        if target.Stats.Oiled > 0 && damtype == "fire" {
                            target.Stats.HP -= 5
                        }
                        fmt.Println(fmt.Sprintf("%s has lost %d hit points, now has %d hit points remaining", target.GetName(), damage, target.Stats.HP))
                        return true, nil
                    } else {
                        fmt.Println(fmt.Sprintf("Ranged attack missed. Roll: %d, AC: %d", hittotal, target.Stats.AC))
                        return false, nil
                    }
                } else {
                    fmt.Println("You do not have line of sight to the enemy")
                    return false, errors.New("You do not have line of sight to the enemy")
                }
            } else {
                return false, errors.New("Enemy is too far away")
            }
        } else {
            fmt.Println(fmt.Sprintf("Can't attack with %s, function %s", p.Equipment.BothHands.PrettyPrint(), p.Equipment.BothHands.Function()))
            return false, errors.New(fmt.Sprintf("Can't attack with %s, function %s", p.Equipment.BothHands.PrettyPrint(), p.Equipment.BothHands.Function()))
        }
    } else if p.Equipment.RightHand != nil {
        if strings.HasPrefix(p.Equipment.RightHand.Function(), "melee") {
            if (p.Pos[0] == target.Pos[0] + 24 && p.Pos[1] == target.Pos[1] + 24) ||
                (p.Pos[0] == target.Pos[0] && p.Pos[1] == target.Pos[1] + 24) ||
                (p.Pos[0] == target.Pos[0] - 24 && p.Pos[1] == target.Pos[1] + 24) ||
                (p.Pos[0] == target.Pos[0] + 24 && p.Pos[1] == target.Pos[1]) ||
                (p.Pos[0] == target.Pos[0] - 24 && p.Pos[1] == target.Pos[1]) ||
                (p.Pos[0] == target.Pos[0] + 24 && p.Pos[1] == target.Pos[1] - 24) ||
                (p.Pos[0] == target.Pos[0] && p.Pos[1] == target.Pos[1] - 24) ||
                (p.Pos[0] == target.Pos[0] - 24 && p.Pos[1] == target.Pos[1] - 24) {
                // Roll to hit (modifiers), if meets or exceeds AC then roll for damage
                hitroll := rand.Intn(20) + 1
                hittotal := hitroll + p.Stats.StrMod
                fmt.Println(fmt.Sprintf("Melee attack roll: %d, plus modifier: %d, equals: %d", hitroll, p.Stats.StrMod, hittotal))
                if hit := hittotal >= target.Stats.AC; hit {
                    fmt.Println(fmt.Sprintf("%d hits against %d", hittotal, target.Stats.AC))
                    damnd, dam, damtype := p.Equipment.RightHand.Damage()
                    damage := 0
                    for x := 0; x < damnd; x++ {
                        newdam := rand.Intn(dam) + 1
                        damage += newdam
                        fmt.Println(fmt.Sprintf("#%d die rolled %d, total damage is now %d", x + 1, newdam, damage))
                    }
                    target.Stats.HP -= damage
                    if target.Stats.Oiled > 0 && damtype == "fire" {
                        target.Stats.HP -= 5
                    }
                    fmt.Println(fmt.Sprintf("%s has lost %d hit points, now has %d hit points remaining", target.GetName(), damage, target.Stats.HP))
                    return true, nil
                } else {
                    fmt.Println(fmt.Sprintf("Melee attack missed, Roll: %d, AC: %d", hittotal, target.Stats.AC))
                    return false, nil
                }
            } else {
                if strings.Contains(p.Equipment.RightHand.Function(), "throw") {
                    if l.Distance(p, target.Pos) <= p.Equipment.RightHand.GetRange()[1] {
                        if ok, _, _ := l.LineOfSight(p, target.Pos); ok {
                            // Check distance, if ok roll to hit (modifiers), if meets or exceeds AC then roll for damage
                            hitroll := rand.Intn(20) + 1
                            if l.Distance(p, target.Pos) >= p.Equipment.RightHand.GetRange()[0] {
                                hitroll2 := rand.Intn(20) + 1
                                if hitroll2 < hitroll {
                                    hitroll = hitroll2
                                }
                            }
                            hittotal := hitroll + p.Stats.DexMod
                            fmt.Println(fmt.Sprintf("Ranged attack roll: %d, plus modifier: %d, equals: %d", hitroll, p.Stats.DexMod, hittotal))
                            if hit := hittotal >= target.Stats.AC; hit {
                                fmt.Println(fmt.Sprintf("%d hits against %d", hittotal, target.Stats.AC))
                                damnd, dam, damtype := p.Equipment.RightHand.Damage()
                                damage := 0
                                for x := 0; x < damnd; x++ {
                                    newdam := rand.Intn(dam) + 1
                                    damage += newdam
                                    fmt.Println(fmt.Sprintf("#%d die rolled %d, total damage is now %d", x + 1, newdam, damage))
                                }
                                target.Stats.HP -= damage
                                if strings.HasPrefix(p.Equipment.RightHand.PrettyPrint(), "Oil Flask") {
                                    target.Stats.Oiled = 10
                                }
                                if target.Stats.Oiled > 0 && damtype == "fire" {
                                    target.Stats.HP -= 5
                                }
                                fmt.Println(fmt.Sprintf("%s has lost %d hit points, now has %d hit points remaining", target.GetName(), damage, target.Stats.HP))
                                return true, nil
                            } else {
                                fmt.Println(fmt.Sprintf("Ranged attack missed. Roll: %d, AC: %d", hittotal, target.Stats.AC))
                                return false, nil
                            }
                        } else {
                            fmt.Println("You do not have line of sight to the enemy")
                            return false, errors.New("You do not have line of sight to the enemy")
                        }
                    } else {
                        return false, errors.New("Enemy is too far away")
                    }
                } else {
                    fmt.Println("Enemy is not within melee range")
                    return false, errors.New("Enemy is not within melee range")
                }
            }
        } else if strings.HasPrefix(p.Equipment.RightHand.Function(), "range") {
            if l.Distance(p, target.Pos) <= p.Equipment.RightHand.GetRange()[1] {
                if ok, _, _ := l.LineOfSight(p, target.Pos); ok {
                    // Check distance, if ok roll to hit (modifiers), if meets or exceeds AC then roll for damage
                    hitroll := rand.Intn(20) + 1
                    if l.Distance(p, target.Pos) >= p.Equipment.RightHand.GetRange()[0] {
                        hitroll2 := rand.Intn(20) + 1
                        if hitroll2 < hitroll {
                            hitroll = hitroll2
                        }
                    }
                    hittotal := hitroll + p.Stats.DexMod
                    fmt.Println(fmt.Sprintf("Ranged attack roll: %d, plus modifier: %d, equals: %d", hitroll, p.Stats.DexMod, hittotal))
                    if hit := hittotal >= target.Stats.AC; hit {
                        fmt.Println(fmt.Sprintf("%d hits against %d", hittotal, target.Stats.AC))
                        damnd, dam, damtype := p.Equipment.RightHand.Damage()
                        damage := 0
                        for x := 0; x < damnd; x++ {
                            newdam := rand.Intn(dam) + 1
                            damage += newdam
                            fmt.Println(fmt.Sprintf("#%d die rolled %d, total damage is now %d", x + 1, newdam, damage))
                        }
                        target.Stats.HP -= damage
                        if strings.HasPrefix(p.Equipment.RightHand.PrettyPrint(), "Oil Flask") {
                            target.Stats.Oiled = 10
                        }
                        if target.Stats.Oiled > 0 && damtype == "fire" {
                            target.Stats.HP -= 5
                        }
                        fmt.Println(fmt.Sprintf("%s has lost %d hit points, now has %d hit points remaining", target.GetName(), damage, target.Stats.HP))
                        return true, nil
                    } else {
                        fmt.Println(fmt.Sprintf("Ranged attack missed. Roll: %d, AC: %d", hittotal, target.Stats.AC))
                        return false, nil
                    }
                } else {
                    fmt.Println("You do not have line of sight to the enemy")
                    return false, errors.New("You do not have line of sight to the enemy")
                }
            } else {
                return false, errors.New("Enemy is too far away")
            }
        } else {
            fmt.Println(fmt.Sprintf("Can't attack with %s, function %s", p.Equipment.RightHand.PrettyPrint(), p.Equipment.RightHand.Function()))
            return false, errors.New(fmt.Sprintf("Can't attack with %s, function %s", p.Equipment.RightHand.PrettyPrint(), p.Equipment.RightHand.Function()))
        }
    } else {
        fmt.Println("No weapon equipped")
    }
    return false, nil
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

func (l *Level) SaveNPCs() string {
    result := ""
    for _, npc := range l.NPCs {
        result += npc.SaveHP()
    }
    return result
}

func LoadLvl(newlvl ...interface{}) *Level {
    if len(newlvl) != 5 && len(newlvl) != 3 {
        log.Fatal("Incorrect number of arguments passed to levels.LoadLvl; should be 3 or 5, got %d", len(newlvl))
        return nil
    }
    switch newlvl[0] {
    case "One":
        var l *Level
        if len(newlvl) == 5 {
            l = lvlOne(newlvl[1].(int), newlvl[4].(string))
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
            if newlvl[4].(string) != "" {
                npchps := strings.Split(newlvl[4].(string), ";")
                npchps = npchps[:len(npchps) - 1]
                for x := 0; x < len(npchps); x++ {
                    npcname := strings.Split(npchps[x], "=")[0]
                    npchp := strings.Split(npchps[x], "=")[1]
                    for _, npc := range l.NPCs {
                        if npc.GetName() == npcname {
                            hpint, err := strconv.Atoi(npchp)
                            if err != nil {
                                panic(err)
                            }
                            npc.PC.Stats.HP = hpint
                        }
                    }
                }
            }
        } else {
            l = lvlOne(newlvl[1].(int), newlvl[2].(string))
        }
        return l
    case "Two":
        var l *Level
        if len(newlvl) == 5 {
            l = lvlTwo(newlvl[1].(int), newlvl[4].(string))
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
            if newlvl[4] != "" {
                npchps := strings.Split(newlvl[4].(string), ";")
                npchps = npchps[:len(npchps) - 1]
                for x := 0; x < len(npchps); x++ {
                    npcname := strings.Split(npchps[x], "=")[0]
                    npchp := strings.Split(npchps[x], "=")[1]
                    for _, npc := range l.NPCs {
                        if npc.GetName() == npcname {
                            hpint, err := strconv.Atoi(npchp)
                            if err != nil {
                                panic(err)
                            }
                            npc.PC.Stats.HP = hpint
                        }
                    }
                }
            }
        } else {
            l = lvlTwo(newlvl[1].(int), newlvl[2].(string))
        }
        return l
    case "VerticalWall":
        var l *Level
        if len(newlvl) == 5 {
            l = verticalWallLvl(newlvl[1].(int), newlvl[4].(string))
            l.Pos = [2]int{newlvl[2].(int), newlvl[3].(int)}
            if newlvl[4] != "" {
                npchps := strings.Split(newlvl[4].(string), ";")
                npchps = npchps[:len(npchps) - 1]
                for x := 0; x < len(npchps); x++ {
                    npcname := strings.Split(npchps[x], "=")[0]
                    npchp := strings.Split(npchps[x], "=")[1]
                    for _, npc := range l.NPCs {
                        if npc.GetName() == npcname {
                            hpint, err := strconv.Atoi(npchp)
                            if err != nil {
                                panic(err)
                            }
                            npc.PC.Stats.HP = hpint
                        }
                    }
                }
            }
        } else {
            l = verticalWallLvl(newlvl[1].(int), newlvl[2].(string))
        }
        return l
    default:
        log.Fatal(fmt.Sprintf("Level %s does not exist", newlvl[0]))
    }
    return lvlOne(newlvl[1].(int), "")
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
        if p.Pos[1] + dist > 0 && p.Pos[1] + dist < l.GetMax()[1] {
            if dist < 0 {
                if !pc {
                    if p.Pos[0] == mc.Pos[0] && p.Pos[1] + dist == mc.Pos[1] {
                        return false, "player"
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] > a[0] - 48 && p.Pos[0] < a[2] && p.Pos[1] + dist >= a[1] && p.Pos[1] + dist < a[3] - 24 {
                        return false, "box"
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist < b.PC.Pos[1] + 48 {
                            return false, "npc"
                        }
                    } else {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist == b.PC.Pos[1] {
                            if attempt < 10 {
                                return false, "npc"
                            }
                        }
                    }
                }
            } else {
                if !pc {
                    if p.Pos[0] == mc.Pos[0] && p.Pos[1] + dist == mc.Pos[1] {
                        return false, "player"
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] > a[0] - 24 && p.Pos[0] < a[2] - 24 && p.Pos[1] + dist >= a[1] - 24 && p.Pos[1] + dist < a[3] {
                        return false, "box"
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist > b.PC.Pos[1] - 48 {
                            return false, "npc"
                        }
                    } else {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist == b.PC.Pos[1] {
                            if attempt < 10 {
                                return false, "npc"
                            }
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
        if p.Pos[0] + dist > 0 && p.Pos[0] + dist < l.GetMax()[0] {
            if dist < 0 {
                if !pc {
                    if p.Pos[0] + dist == mc.Pos[0] && p.Pos[1] == mc.Pos[1] {
                        return false, "player"
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] + dist >= a[0] && p.Pos[0] + dist < a[2] && p.Pos[1] >= a[1] - 24 && p.Pos[1] < a[3] - 24 {
                        return false, "box"
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] + dist == b.PC.Pos[0] && p.Pos[1] >= b.PC.Pos[1] - 24 && p.Pos[1] <= b.PC.Pos[1] + 24 {
                            return false, "npc"
                        }
                    } else {
                        if p.Pos[0] + dist >= b.PC.Pos[0] && p.Pos[0] + dist < b.PC.Pos[0] + 24 && p.Pos[1] == b.PC.Pos[1] {
                            if attempt < 10 {
                                return false, "npc"
                            }
                        }
                    }
                }
            } else {
                if !pc {
                    if p.Pos[0] + dist == mc.Pos[0] && p.Pos[1] == mc.Pos[1] {
                        return false, "player"
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] + dist >= a[0] - 24 && p.Pos[0] + dist < a[2] && p.Pos[1] >= a[1] - 24 && p.Pos[1] < a[3] - 24 {
                        return false, "box"
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] + dist >= b.PC.Pos[0] && p.Pos[1] >= b.PC.Pos[1] - 24 && p.Pos[1] <= b.PC.Pos[1] + 24 {
                            return false, "npc"
                        }
                    } else {
                        if p.Pos[0] + dist >= b.PC.Pos[0] && p.Pos[0] + dist < b.PC.Pos[0] + 24 && p.Pos[1] == b.PC.Pos[1] {
                            if attempt < 10 {
                                return false, "npc"
                            }
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
    return math.Sqrt(math.Pow(float64(target[0] - p.Pos[0]), 2.0) - math.Pow(float64(target[1] - p.Pos[1]), 2.0))
}
