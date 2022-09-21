package player

import (
    "fmt"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/inventory"
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
    Darkvision bool
    Proficiencies []string
    Resistances []string
    Lucky bool
    Nimbleness bool
    Brave bool
    Ancestry string
}

type Equipment struct {
    Head inventory.Item
    Torso inventory.Item
    Legs inventory.Item
    Feet inventory.Item
    LeftHand inventory.Item
    RightHand inventory.Item
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
    Background string
    Alignment string
    XP int
    Equipment *Equipment
}

func (s *Stats) Check() error {
    return nil
}

func (p *Player) Unequip(slot string) {
    switch slot {
    case "Head":
        if p.Equipment.Head == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.Head)
        p.Equipment.Head = nil
    case "Torso":
        if p.Equipment.Torso == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.Torso)
        p.Equipment.Torso = nil
    case "Legs":
        if p.Equipment.Legs == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.Legs)
        p.Equipment.Legs = nil
    case "Feet":
        if p.Equipment.Feet == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.Feet)
        p.Equipment.Feet = nil
    case "LeftPinky":
        if p.Equipment.LeftPinky == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.LeftPinky)
        p.Equipment.LeftPinky = nil
    case "LeftRing":
        if p.Equipment.LeftRing == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.LeftRing)
        p.Equipment.LeftRing = nil
    case "LeftMid":
        if p.Equipment.LeftMid == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.LeftMid)
        p.Equipment.LeftMid = nil
    case "LeftInd":
        if p.Equipment.LeftInd == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.LeftInd)
        p.Equipment.LeftInd = nil
    case "LeftThumb":
        if p.Equipment.LeftThumb == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.LeftThumb)
        p.Equipment.LeftThumb = nil
    case "RightPinky":
        if p.Equipment.RightPinky == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.RightPinky)
        p.Equipment.RightPinky = nil
    case "RightRing":
        if p.Equipment.RightRing == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.RightRing)
        p.Equipment.RightRing = nil
    case "RightMid":
        if p.Equipment.RightMid == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.RightMid)
        p.Equipment.RightMid = nil
    case "RightInd":
        if p.Equipment.RightInd == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.RightInd)
        p.Equipment.RightInd = nil
    case "RightThumb":
        if p.Equipment.RightThumb == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.RightThumb)
        p.Equipment.RightThumb = nil
    case "LeftHand":
        if p.Equipment.LeftHand == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.LeftHand)
        p.Equipment.LeftHand = nil
    case "RightHand":
        if p.Equipment.RightHand == nil {
            log.Fatal("Nothing on my head")
        }
        p.Inv.Add(p.Equipment.RightHand)
        p.Equipment.RightHand = nil
    default:
        log.Fatal(fmt.Sprintf("%s is not a valid slot", slot))
    }
}

func (p *Player) Equip(item inventory.Item) {
    for _, b := range p.Inv.GetItems() {
        if b == item {
            switch item.Slot() {
            case "Head":
                if p.Equipment.Head != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.Head = item
                p.Inv.Drop(item)
            case "Torso":
                if p.Equipment.Torso != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.Torso = item
                p.Inv.Drop(item)
            case "Legs":
                if p.Equipment.Legs != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.Legs = item
                p.Inv.Drop(item)
            case "Feet":
                if p.Equipment.Feet != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.Feet = item
                p.Inv.Drop(item)
            case "LeftPinky":
                if p.Equipment.LeftPinky != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.LeftPinky = item
                p.Inv.Drop(item)
            case "LeftRing":
                if p.Equipment.LeftRing != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.LeftRing = item
                p.Inv.Drop(item)
            case "LeftMid":
                if p.Equipment.LeftMid != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.LeftMid = item
                p.Inv.Drop(item)
            case "LeftInd":
                if p.Equipment.LeftInd != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.LeftInd = item
                p.Inv.Drop(item)
            case "LeftThumb":
                if p.Equipment.LeftThumb != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.LeftThumb = item
                p.Inv.Drop(item)
            case "RightPinky":
                if p.Equipment.RightPinky != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.RightPinky = item
                p.Inv.Drop(item)
            case "RightRing":
                if p.Equipment.RightRing != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.RightRing = item
                p.Inv.Drop(item)
            case "RightMid":
                if p.Equipment.RightMid != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.RightMid = item
                p.Inv.Drop(item)
            case "RightInd":
                if p.Equipment.RightInd != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.RightInd = item
                p.Inv.Drop(item)
            case "RightThumb":
                if p.Equipment.RightThumb != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.RightThumb = item
                p.Inv.Drop(item)
            case "LeftHand":
                if p.Equipment.LeftHand != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.LeftHand = item
                p.Inv.Drop(item)
            case "RightHand":
                if p.Equipment.RightHand != nil {
                    p.Inv.Add(p.Equipment.Head)
                }
                p.Equipment.RightHand = item
                p.Inv.Drop(item)
            default:
                log.Fatal(fmt.Sprintf("%s is not a valid slot", item.Slot))
            }
        }
    }
}
