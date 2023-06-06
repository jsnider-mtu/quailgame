package player

import (
    "crypto/md5"
    "fmt"
    "log"
    "strings"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/classes"
    "github.com/jsnider-mtu/quailgame/inventory"
    "github.com/jsnider-mtu/quailgame/inventory/items"
)

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

type Player struct {
    Name string
    Pos [2]int
    Inv *inventory.Inv
    Image *ebiten.Image
    Class classes.Class
    Equipment *Equipment
    WriteMsg string
    PageMsgs [][]interface{}
}

func (e *Equipment) Save() string {
    var result string
    if e.Armor != nil {
        result += fmt.Sprintf("Armor=%s|", e.Armor.Save())
    }
    if e.Head != nil {
        result += fmt.Sprintf("Head=%s|", e.Head.Save())
    }
    if e.Torso != nil {
        result += fmt.Sprintf("Torso=%s|", e.Torso.Save())
    }
    if e.Legs != nil {
        result += fmt.Sprintf("Legs=%s|", e.Legs.Save())
    }
    if e.Feet != nil {
        result += fmt.Sprintf("Feet=%s|", e.Feet.Save())
    }
    if e.LeftPinky != nil {
        result += fmt.Sprintf("LeftPinky=%s|", e.LeftPinky.Save())
    }
    if e.LeftRing != nil {
        result += fmt.Sprintf("LeftRing=%s|", e.LeftRing.Save())
    }
    if e.LeftMid != nil {
        result += fmt.Sprintf("LeftMid=%s|", e.LeftMid.Save())
    }
    if e.LeftInd != nil {
        result += fmt.Sprintf("LeftInd=%s|", e.LeftInd.Save())
    }
    if e.LeftThumb != nil {
        result += fmt.Sprintf("LeftThumb=%s|", e.LeftThumb.Save())
    }
    if e.RightPinky != nil {
        result += fmt.Sprintf("RightPinky=%s|", e.RightPinky.Save())
    }
    if e.RightRing != nil {
        result += fmt.Sprintf("RightRing=%s|", e.RightRing.Save())
    }
    if e.RightMid != nil {
        result += fmt.Sprintf("RightMid=%s|", e.RightMid.Save())
    }
    if e.RightInd != nil {
        result += fmt.Sprintf("RightInd=%s|", e.RightInd.Save())
    }
    if e.RightThumb != nil {
        result += fmt.Sprintf("RightThumb=%s|", e.RightThumb.Save())
    }
    if e.BothHands == nil {
        if e.LeftHand != nil {
            result += fmt.Sprintf("LeftHand=%s|", e.LeftHand.Save())
        }
        if e.RightHand != nil {
            result += fmt.Sprintf("RightHand=%s|", e.RightHand.Save())
        }
    } else {
        result += fmt.Sprintf("BothHands=%s|", e.BothHands.Save())
    }
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

func (p *Player) Effects(action string, data []int, c chan int) {
    switch action {
    case "illuminate":
        if len(data) != 3 {
            log.Fatal(fmt.Sprintf("Incorrect # of arguments %d for illuminate data", len(data)))
        }
        dataarr := [3]int{}
        copy(dataarr[:], data)
        p.Class.Illuminate(dataarr)
        return
    case "disguise":
        log.Println("Need to implement disguise action")
        return
    case "write":
        reqs := []string{"Ink Bottle", "Paper"}
        for _, i := range p.Inv.GetItems() {
            for x := 0; x < len(reqs); x++ {
                for in, j := range reqs {
                    if strings.HasPrefix(i.PrettyPrint(), j) {
                        if i.GetQuantity() > 0 {
                            reqs = append(reqs[:in], reqs[in + 1:]...)
                            log.Println(fmt.Sprint(reqs))
                            break
                        }
                    }
                }
            }
        }
        if len(reqs) == 0 {
            for {
                msg := <-c
                switch msg {
                case 0:
                    if p.WriteMsg != "" {
                        p.Inv.GetItems()[data[0]].(*items.Paper).Write(fmt.Sprintf("%x", md5.Sum([]byte(p.WriteMsg))), p.WriteMsg)
                        p.WriteMsg = ""
                        return
                    } else {
                        log.Println("p.WriteMsg was empty, waiting some more")
                    }
                case 1:
                    return
                default:
                    log.Println("Waiting on p.WriteMsg to be populated")
                }
            }
        } else {
            msg := "Missing requirements: "
            for _, k := range reqs {
                msg += k + ", "
            }
            strings.TrimRight(msg, ", ")
            log.Println(msg)
        }
        return
    case "read":
        return
    case "throw":
        var origitem inventory.Item
        if p.Equipment.BothHands != nil {
            origitem = p.Equipment.BothHands
        } else if p.Equipment.RightHand != nil {
            origitem = p.Equipment.RightHand
        }
        c <- data[0]
        c <- data[1]
        savename := strings.Split(p.Inv.GetItems()[data[2]].Save(), ",")[0]
        p.Equip(p.Inv.GetItems()[data[2]])
        switch savename {
        case "OilFlask":
            c <- 0
        case "Dagger":
            c <- 1
        case "Handaxe":
            c <- 2
        case "LightHammer":
            c <- 3
        default:
            log.Println("Not a valid throwable item")
        }
        msg := <-c
        if origitem != nil {
            p.Equip(origitem)
        }
        switch msg {
        case 0:
            p.Inv.GetItems()[len(p.Inv.GetItems()) - 1].(*items.OilFlask).Quantity--
            return
        case 1:
            return
        case 2:
            p.Inv.GetItems()[len(p.Inv.GetItems()) - 1].(*items.Handaxe).Quantity--
        case 3:
            p.Inv.GetItems()[len(p.Inv.GetItems()) - 1].(*items.LightHammer).Quantity--
        default:
            return
        }
        return
    case "playmusic":
        var origitem inventory.Item
        if p.Equipment.BothHands != nil {
            origitem = p.Equipment.BothHands
        } else if p.Equipment.RightHand != nil {
            origitem = p.Equipment.RightHand
        }
        p.Equip(p.Inv.GetItems()[data[0]])
        log.Println("Started playing music")
        msg := <-c
        if origitem != nil {
            p.Equip(origitem)
        }
        if msg == 0 {
            log.Println("Stopped playing music")
            return
        }
        return
    default:
        log.Println(action + " is not a recognized action")
        return
    }
    return
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
                if p.Equipment.BothHands != nil {
                    p.Unequip("BothHands")
                }
                p.Equipment.LeftHand = item
                p.Inv.Drop(item)
            case "RightHand":
                if p.Equipment.RightHand != nil {
                    p.Inv.Add(p.Equipment.RightHand)
                }
                if p.Equipment.BothHands != nil {
                    p.Unequip("BothHands")
                }
                p.Equipment.RightHand = item
                p.Inv.Drop(item)
            case "BothHands":
                if p.Equipment.BothHands != nil {
                    p.Inv.Add(p.Equipment.BothHands)
                }
                if p.Equipment.RightHand != nil {
                    p.Unequip("RightHand")
                }
                if p.Equipment.LeftHand != nil {
                    p.Unequip("LeftHand")
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
