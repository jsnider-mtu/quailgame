package player

import (
    "fmt"
    "log"
    "math"
    "strconv"
    "strings"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/inventory"
    "github.com/jsnider-mtu/quailgame/spells"
)

type Stats struct {
    AC int
    Str int
    StrMod int
    Dex int
    DexMod int
    Con int
    ConMod int
    Intel int
    IntelMod int
    Wis int
    WisMod int
    Cha int
    ChaMod int
    ProfBonus int
    Initiative int
    SavingThrows map[string]int
    Skills map[string]int
    Speed int
    MaxHP int
    HP int
    TempHP int
    HitDice string
    DeathSaveSucc int
    DeathSaveFail int
    Languages []string
    Size int
    Inspiration bool
    Darkvision bool
    Proficiencies []string
    Resistances []string
    Lucky bool
    Nimbleness bool
    Brave bool
    Ancestry string
}

type Equipment struct {
    Armor inventory.Item
    Head inventory.Item
    Torso inventory.Item
    Legs inventory.Item
    Feet inventory.Item
    LeftHand inventory.Item
    RightHand inventory.Item
    BothHands inventory.Item
    LeftPinky inventory.Item
    LeftRing inventory.Item
    LeftMid inventory.Item
    LeftInd inventory.Item
    LeftThumb inventory.Item
    RightPinky inventory.Item
    RightRing inventory.Item
    RightMid inventory.Item
    RightInd inventory.Item
    RightThumb inventory.Item
    Clothes inventory.Item
}

type Spells struct {
    Spells []spells.Spell
}

type Player struct {
    Name string
    Pos [2]int
    Inv *inventory.Inv
    Image *ebiten.Image
    Stats *Stats
    Race string
    Class string
    Level int
    XP int
    Equipment *Equipment
    Spells *Spells
    Concentrating bool
    TurnsCasting int
    TurnsCastEnd int
}

func (s *Spells) Add(spellsslice []string) error {
    for _, spell := range spellsslice {
        if sp, ok := spells.New(spell); ok == nil {
            s.Spells = append(s.Spells, sp)
        } else {
            return ok
            //return errors.New(fmt.Sprintf("%s is not a valid spell", spell))
        }
    }
    return nil
}

func (s *Spells) Save() string {
    var spellsstr string
    for sind, spell := range s.Spells {
        if sind == len(s.Spells) - 1 {
            spellsstr += spell.PrettyPrint()
        } else {
            spellsstr += spell.PrettyPrint() + ","
        }
    }
    return spellsstr
}

func (s *Stats) Check() error {
    return nil
}

func (s *Stats) Save() string {
    var savingthrowsstr string
    for stkey, stval := range s.SavingThrows {
        savingthrowsstr += stkey + "=" + strconv.Itoa(stval) + ","
    }
    var skillsstr string
    for skkey, skval := range s.Skills {
        skillsstr += skkey + "=" + strconv.Itoa(skval) + ","
    }
    var languagestr string
    for lind, language := range s.Languages {
        if lind == len(s.Languages) - 1 {
            languagestr += language
        } else {
            languagestr += language + ","
        }
    }
    var proficienciesstr string
    for pind, prof := range s.Proficiencies {
        if pind == len(s.Proficiencies) - 1 {
            proficienciesstr += prof
        } else {
            proficienciesstr += prof + ","
        }
    }
    var resistancesstr string
    for rind, resist := range s.Resistances {
        if rind == len(s.Resistances) - 1 {
            resistancesstr += resist
        } else {
            resistancesstr += resist + ","
        }
    }
    return "AC:" + strconv.Itoa(s.AC) + ";" +
           "Str:" + strconv.Itoa(s.Str) + ";" +
           "StrMod:" + strconv.Itoa(s.StrMod) + ";" +
           "Dex:" + strconv.Itoa(s.Dex) + ";" +
           "DexMod:" + strconv.Itoa(s.DexMod) + ";" +
           "Con:" + strconv.Itoa(s.Con) + ";" +
           "ConMod:" + strconv.Itoa(s.ConMod) + ";" +
           "Intel:" + strconv.Itoa(s.Intel) + ";" +
           "IntelMod:" + strconv.Itoa(s.IntelMod) + ";" +
           "Wis:" + strconv.Itoa(s.Wis) + ";" +
           "WisMod:" + strconv.Itoa(s.WisMod) + ";" +
           "Cha:" + strconv.Itoa(s.Cha) + ";" +
           "ChaMod:" + strconv.Itoa(s.ChaMod) + ";" +
           "ProfBonus:" + strconv.Itoa(s.ProfBonus) + ";" +
           "Initiative:" + strconv.Itoa(s.Initiative) + ";" +
           "SavingThrows:" + savingthrowsstr + ";" +
           "Skills:" + skillsstr + ";" +
           "Speed:" + strconv.Itoa(s.Speed) + ";" +
           "MaxHP:" + strconv.Itoa(s.MaxHP) + ";" +
           "HP:" + strconv.Itoa(s.HP) + ";" +
           "TempHP:" + strconv.Itoa(s.TempHP) + ";" +
           "HitDice:" + s.HitDice + ";" +
           "DeathSaveSucc:" + strconv.Itoa(s.DeathSaveSucc) + ";" +
           "DeathSaveFail:" + strconv.Itoa(s.DeathSaveFail) + ";" +
           "Languages:" + languagestr + ";" +
           "Size:" + strconv.Itoa(s.Size) + ";" +
           "Inspiration:" + strconv.FormatBool(s.Inspiration) + ";" +
           "Darkvision:" + strconv.FormatBool(s.Darkvision) + ";" +
           "Proficiencies:" + proficienciesstr + ";" +
           "Resistances:" + resistancesstr + ";" +
           "Lucky:" + strconv.FormatBool(s.Lucky) + ";" +
           "Nimbleness:" + strconv.FormatBool(s.Nimbleness) + ";" +
           "Brave:" + strconv.FormatBool(s.Brave) + ";" +
           "Ancestry:" + s.Ancestry + ";"
}

func (e *Equipment) Save() string {
    var result string
    if e.Armor == nil {
        log.Println("e.Armor == nil")
    } else {
        result += fmt.Sprintf("Armor=%s|", e.Armor.Save())
    }
    if e.Head == nil {
        log.Println("e.Head == nil")
    } else {
        result += fmt.Sprintf("Head=%s|", e.Head.Save())
    }
    if e.Torso == nil {
        log.Println("e.Torso == nil")
    } else {
        result += fmt.Sprintf("Torso=%s|", e.Torso.Save())
    }
    if e.Legs == nil {
        log.Println("e.Legs == nil")
    } else {
        result += fmt.Sprintf("Legs=%s|", e.Legs.Save())
    }
    if e.Feet == nil {
        log.Println("e.Feet == nil")
    } else {
        result += fmt.Sprintf("Feet=%s|", e.Feet.Save())
    }
    if e.LeftPinky == nil {
        log.Println("e.LeftPinky == nil")
    } else {
        result += fmt.Sprintf("LeftPinky=%s|", e.LeftPinky.Save())
    }
    if e.LeftRing == nil {
        log.Println("e.LeftRing == nil")
    } else {
        result += fmt.Sprintf("LeftRing=%s|", e.LeftRing.Save())
    }
    if e.LeftMid == nil {
        log.Println("e.LeftMid == nil")
    } else {
        result += fmt.Sprintf("LeftMid=%s|", e.LeftMid.Save())
    }
    if e.LeftInd == nil {
        log.Println("e.LeftInd == nil")
    } else {
        result += fmt.Sprintf("LeftInd=%s|", e.LeftInd.Save())
    }
    if e.LeftThumb == nil {
        log.Println("e.LeftThumb == nil")
    } else {
        result += fmt.Sprintf("LeftThumb=%s|", e.LeftThumb.Save())
    }
    if e.RightPinky == nil {
        log.Println("e.RightPinky == nil")
    } else {
        result += fmt.Sprintf("RightPinky=%s|", e.RightPinky.Save())
    }
    if e.RightRing == nil {
        log.Println("e.RightRing == nil")
    } else {
        result += fmt.Sprintf("RightRing=%s|", e.RightRing.Save())
    }
    if e.RightMid == nil {
        log.Println("e.RightMid == nil")
    } else {
        result += fmt.Sprintf("RightMid=%s|", e.RightMid.Save())
    }
    if e.RightInd == nil {
        log.Println("e.RightInd == nil")
    } else {
        result += fmt.Sprintf("RightInd=%s|", e.RightInd.Save())
    }
    if e.RightThumb == nil {
        log.Println("e.RightThumb == nil")
    } else {
        result += fmt.Sprintf("RightThumb=%s|", e.RightThumb.Save())
    }
    if e.BothHands == nil {
        log.Println("e.BothHands == nil")
        if e.LeftHand == nil {
            log.Println("e.LeftHand == nil")
        } else {
            result += fmt.Sprintf("LeftHand=%s|", e.LeftHand.Save())
        }
        if e.RightHand == nil {
            log.Println("e.RightHand == nil")
        } else {
            result += fmt.Sprintf("RightHand=%s|", e.RightHand.Save())
        }
    } else {
        result += fmt.Sprintf("BothHands=%s|", e.BothHands.Save())
    }
    //if e.BothHands == nil {
    //    log.Println("e.BothHands == nil")
    //} else {
    //    result += fmt.Sprintf("BothHands=%s|", e.BothHands.Save())
    //}
    if e.Clothes == nil {
        log.Println("e.Clothes == nil")
    } else {
        result += fmt.Sprintf("Clothes=%s|", e.Clothes.Save())
    }
    result += ";"
    result = strings.Replace(result, "|;", ";", 1)
    log.Println("result = " + result)
    return result
}

func (p *Player) GetName() string {
    return p.Name
}

func (p *Player) CanSpeak() bool {
    return true
}

func (p *Player) CanMove() bool {
    return true
}

func (p *Player) HaveMat() bool {
    if p.Equipment.LeftHand.PrettyPrint() == "Arcane Focus" || p.Equipment.LeftHand.PrettyPrint() == "Component Pouch" {
        return true
    }
    return false
}

func (p *Player) CastSpell(spell spells.Spell, target *Player) int {
    log.Println(fmt.Sprintf("Casting spell %s, target is %s", spell.PrettyPrint(), target.GetName()))
    //spellrange := spell.GetRange()
    // is target within spellrange?
    // get x, y of player and target and calculate distance
    playerx := p.Pos[0] + 24
    playery := p.Pos[1] + 24
    targetx := target.Pos[0] + 24
    targety := target.Pos[1] + 24
    // 5 ft == 24 px
    dist := math.Sqrt(math.Pow(float64(targetx - playerx), 2.0) + math.Pow(float64(targety - playery), 2.0))
    if dist <= spell.GetRange() {
        // get casting time
        // if not instant start counting
        switch spell.GetComponents() {
        case "V,S,M":
            if !p.CanSpeak() || !p.CanMove() || !p.HaveMat() {
                return 0
            }
        case "S,M":
            if !p.CanMove() || !p.HaveMat() {
                return 0
            }
        case "M":
            if !p.HaveMat() {
                return 0
            }
        case "V,S":
            if !p.CanSpeak() || !p.CanMove() {
                return 0
            }
        case "V":
            if !p.CanSpeak() {
                return 0
            }
        case "S":
            if !p.CanMove() {
                return 0
            }
        default:
            log.Fatal(fmt.Sprintf("%s is an invalid spell component", spell.GetComponents()))
        }
        switch spell.GetCastTime() {
        case "minute":
            p.Concentrating = true
            p.TurnsCastEnd = 10
            p.TurnsCasting++
            if p.TurnsCasting == p.TurnsCastEnd {
                if !spell.GetConcentration() {
                    p.Concentrating = false
                }
                p.TurnsCastEnd = 0
                p.TurnsCasting = 0
                return 0
            }
            return p.TurnsCastEnd - p.TurnsCasting
        case "hour":
            p.Concentrating = true
            p.TurnsCastEnd = 600
            p.TurnsCasting++
            if p.TurnsCasting == p.TurnsCastEnd {
                if !spell.GetConcentration() {
                    p.Concentrating = false
                }
                p.TurnsCastEnd = 0
                p.TurnsCasting = 0
                return 0
            }
            return p.TurnsCastEnd - p.TurnsCasting
        case "action":
            if spell.GetConcentration() {
                p.Concentrating = true
            }
            return 0
        case "bonus":
            if spell.GetConcentration() {
                p.Concentrating = true
            }
            return 0
        case "reaction":
            if spell.GetConcentration() {
                p.Concentrating = true
            }
            return 0
        default:
            return 0
        }
        // V can I speak?
        // S can I move at least a hand?
        // M do I have arcane focus or component pouch?
        // Does target need to make a saving throw?
        // Deal damage or condition on target
    } else {
        log.Println(fmt.Sprintf("%s is out of range of %s", target.GetName(), spell.PrettyPrint()))
    }
    return 0
}

func (p *Player) Unequip(slot string) {
    switch slot {
    case "Armor":
        if p.Equipment.Armor == nil {
            log.Fatal("No armor equipped")
        }
        p.Inv.Add(p.Equipment.Armor)
        p.Equipment.Armor = nil
    case "Head":
        if p.Equipment.Head == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.Head)
        p.Equipment.Head = nil
    case "Torso":
        if p.Equipment.Torso == nil {
            log.Fatal("Nothing on my torso")
        }
        p.Inv.Add(p.Equipment.Torso)
        p.Equipment.Torso = nil
    case "Legs":
        if p.Equipment.Legs == nil {
            log.Fatal("Nothing on my legs")
        }
        p.Inv.Add(p.Equipment.Legs)
        p.Equipment.Legs = nil
    case "Feet":
        if p.Equipment.Feet == nil {
            log.Fatal("Nothing on my feet")
        }
        p.Inv.Add(p.Equipment.Feet)
        p.Equipment.Feet = nil
    case "LeftPinky":
        if p.Equipment.LeftPinky == nil {
            log.Fatal("Nothing on my left pinky")
        }
        p.Inv.Add(p.Equipment.LeftPinky)
        p.Equipment.LeftPinky = nil
    case "LeftRing":
        if p.Equipment.LeftRing == nil {
            log.Fatal("Nothing on my left ring finger")
        }
        p.Inv.Add(p.Equipment.LeftRing)
        p.Equipment.LeftRing = nil
    case "LeftMid":
        if p.Equipment.LeftMid == nil {
            log.Fatal("Nothing on my left middle finger")
        }
        p.Inv.Add(p.Equipment.LeftMid)
        p.Equipment.LeftMid = nil
    case "LeftInd":
        if p.Equipment.LeftInd == nil {
            log.Fatal("Nothing on my left index finger")
        }
        p.Inv.Add(p.Equipment.LeftInd)
        p.Equipment.LeftInd = nil
    case "LeftThumb":
        if p.Equipment.LeftThumb == nil {
            log.Fatal("Nothing on my left thumb")
        }
        p.Inv.Add(p.Equipment.LeftThumb)
        p.Equipment.LeftThumb = nil
    case "RightPinky":
        if p.Equipment.RightPinky == nil {
            log.Fatal("Nothing on my right pinky")
        }
        p.Inv.Add(p.Equipment.RightPinky)
        p.Equipment.RightPinky = nil
    case "RightRing":
        if p.Equipment.RightRing == nil {
            log.Fatal("Nothing on my right ring finger")
        }
        p.Inv.Add(p.Equipment.RightRing)
        p.Equipment.RightRing = nil
    case "RightMid":
        if p.Equipment.RightMid == nil {
            log.Fatal("Nothing on my right middle finger")
        }
        p.Inv.Add(p.Equipment.RightMid)
        p.Equipment.RightMid = nil
    case "RightInd":
        if p.Equipment.RightInd == nil {
            log.Fatal("Nothing on my right index finger")
        }
        p.Inv.Add(p.Equipment.RightInd)
        p.Equipment.RightInd = nil
    case "RightThumb":
        if p.Equipment.RightThumb == nil {
            log.Fatal("Nothing on my right thumb")
        }
        p.Inv.Add(p.Equipment.RightThumb)
        p.Equipment.RightThumb = nil
    case "LeftHand":
        if p.Equipment.LeftHand == nil {
            log.Fatal("Nothing in my left hand")
        }
        p.Inv.Add(p.Equipment.LeftHand)
        p.Equipment.LeftHand = nil
    case "RightHand":
        if p.Equipment.RightHand == nil {
            log.Fatal("Nothing in my right hand")
        }
        p.Inv.Add(p.Equipment.RightHand)
        p.Equipment.RightHand = nil
    case "BothHands":
        if p.Equipment.BothHands == nil {
            log.Fatal("Nothing in both my hands")
        }
        p.Inv.Add(p.Equipment.BothHands)
        p.Equipment.BothHands = nil
    case "Clothes":
        if p.Equipment.Clothes == nil {
            log.Fatal("No clothes equipped")
        }
        p.Inv.Add(p.Equipment.Clothes)
        p.Equipment.Clothes = nil
    default:
        log.Fatal(fmt.Sprintf("%s is not a valid slot", slot))
    }
}

func (p *Player) Equip(item inventory.Item) {
    for _, b := range p.Inv.GetItems() {
        if b.PrettyPrint() == item.PrettyPrint() {
            switch item.Slot() {
            case "Armor":
                if p.Equipment.Armor != nil {
                    p.Inv.Add(p.Equipment.Armor)
                }
                p.Equipment.Armor = item
                p.Inv.Drop(item)
            case "Head":
                if p.Equipment.Head != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.Head = item
                p.Inv.Drop(item)
            case "Torso":
                if p.Equipment.Torso != nil {
                    p.Inv.Add(p.Equipment.Torso)
                }
                p.Equipment.Torso = item
                p.Inv.Drop(item)
            case "Legs":
                if p.Equipment.Legs != nil {
                    p.Inv.Add(p.Equipment.Legs)
                }
                p.Equipment.Legs = item
                p.Inv.Drop(item)
            case "Feet":
                if p.Equipment.Feet != nil {
                    p.Inv.Add(p.Equipment.Feet)
                }
                p.Equipment.Feet = item
                p.Inv.Drop(item)
            case "LeftPinky":
                if p.Equipment.LeftPinky != nil {
                    p.Inv.Add(p.Equipment.LeftPinky)
                }
                p.Equipment.LeftPinky = item
                p.Inv.Drop(item)
            case "LeftRing":
                if p.Equipment.LeftRing != nil {
                    p.Inv.Add(p.Equipment.LeftRing)
                }
                p.Equipment.LeftRing = item
                p.Inv.Drop(item)
            case "LeftMid":
                if p.Equipment.LeftMid != nil {
                    p.Inv.Add(p.Equipment.LeftMid)
                }
                p.Equipment.LeftMid = item
                p.Inv.Drop(item)
            case "LeftInd":
                if p.Equipment.LeftInd != nil {
                    p.Inv.Add(p.Equipment.LeftInd)
                }
                p.Equipment.LeftInd = item
                p.Inv.Drop(item)
            case "LeftThumb":
                if p.Equipment.LeftThumb != nil {
                    p.Inv.Add(p.Equipment.LeftThumb)
                }
                p.Equipment.LeftThumb = item
                p.Inv.Drop(item)
            case "RightPinky":
                if p.Equipment.RightPinky != nil {
                    p.Inv.Add(p.Equipment.RightPinky)
                }
                p.Equipment.RightPinky = item
                p.Inv.Drop(item)
            case "RightRing":
                if p.Equipment.RightRing != nil {
                    p.Inv.Add(p.Equipment.RightRing)
                }
                p.Equipment.RightRing = item
                p.Inv.Drop(item)
            case "RightMid":
                if p.Equipment.RightMid != nil {
                    p.Inv.Add(p.Equipment.RightMid)
                }
                p.Equipment.RightMid = item
                p.Inv.Drop(item)
            case "RightInd":
                if p.Equipment.RightInd != nil {
                    p.Inv.Add(p.Equipment.RightInd)
                }
                p.Equipment.RightInd = item
                p.Inv.Drop(item)
            case "RightThumb":
                if p.Equipment.RightThumb != nil {
                    p.Inv.Add(p.Equipment.RightThumb)
                }
                p.Equipment.RightThumb = item
                p.Inv.Drop(item)
            case "LeftHand":
                if p.Equipment.LeftHand != nil {
                    p.Inv.Add(p.Equipment.LeftHand)
                }
                p.Equipment.LeftHand = item
                p.Inv.Drop(item)
            case "RightHand":
                if p.Equipment.RightHand != nil {
                    p.Inv.Add(p.Equipment.RightHand)
                }
                p.Equipment.RightHand = item
                p.Inv.Drop(item)
            case "BothHands":
                if p.Equipment.BothHands != nil {
                    p.Inv.Add(p.Equipment.BothHands)
                }
                p.Equipment.BothHands = item
                p.Inv.Drop(item)
            case "Clothes":
                if p.Equipment.Clothes != nil {
                    p.Inv.Add(p.Equipment.Clothes)
                }
                p.Equipment.Clothes = item
                p.Inv.Drop(item)
            default:
                log.Fatal(fmt.Sprintf("%s is not a valid slot", item.Slot()))
            }
        }
    }
}
