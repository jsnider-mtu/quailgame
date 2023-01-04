package player

import (
    "crypto/md5"
    "errors"
    "fmt"
    "log"
    "strconv"
    "strings"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/inventory"
    "github.com/jsnider-mtu/quailgame/inventory/items"
    "github.com/jsnider-mtu/quailgame/spells"
//    "github.com/jsnider-mtu/quailgame/utils"

//    "github.com/hajimehoshi/ebiten/v2/text"
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
    Illuminated []int
    Oiled int
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
    CurLevel string
    WriteMsg string
    PageMsgs [][]interface{}
}

func (s *Spells) Add(spellsslice []string) error {
    for _, spell := range spellsslice {
        switch spell {
        case "Acid Splash":
            var acidsplash = spells.AcidSplash{}
            s.Spells = append(s.Spells, acidsplash)
        case "Alarm":
            var alarm = spells.Alarm{}
            s.Spells = append(s.Spells, alarm)
        case "Animal Friendship":
            var animalfriendship = spells.AnimalFriendship{}
            s.Spells = append(s.Spells, animalfriendship)
        case "Armor of Agathys":
            var armorofagathys = spells.ArmorOfAgathys{}
            s.Spells = append(s.Spells, armorofagathys)
        case "Arms of Hadar":
            var armsofhadar = spells.ArmsOfHadar{}
            s.Spells = append(s.Spells, armsofhadar)
        case "Bane":
            var bane = spells.Bane{}
            s.Spells = append(s.Spells, bane)
        case "Blade Ward":
            var bladeward = spells.BladeWard{}
            s.Spells = append(s.Spells, bladeward)
        case "Bless":
            var bless = spells.Bless{}
            s.Spells = append(s.Spells, bless)
        case "Burning Hands":
            var burninghands = spells.BurningHands{}
            s.Spells = append(s.Spells, burninghands)
        case "Charm Person":
            var charmperson = spells.CharmPerson{}
            s.Spells = append(s.Spells, charmperson)
        case "Chill Touch":
            var chilltouch = spells.ChillTouch{}
            s.Spells = append(s.Spells, chilltouch)
        case "Chromatic Orb":
            var chromaticorb = spells.ChromaticOrb{}
            s.Spells = append(s.Spells, chromaticorb)
        case "Color Spray":
            var colorspray = spells.ColorSpray{}
            s.Spells = append(s.Spells, colorspray)
        case "Command":
            var command = spells.Command{}
            s.Spells = append(s.Spells, command)
        case "Compelled Duel":
            var compelledduel = spells.CompelledDuel{}
            s.Spells = append(s.Spells, compelledduel)
        case "Comprehend Languages":
            var comprehendlanguages = spells.ComprehendLanguages{}
            s.Spells = append(s.Spells, comprehendlanguages)
        case "Create or Destroy Water":
            var createordestroywater = spells.CreateOrDestroyWater{}
            s.Spells = append(s.Spells, createordestroywater)
        case "Cure Wounds":
            var curewounds = spells.CureWounds{}
            s.Spells = append(s.Spells, curewounds)
        case "Dancing Lights":
            var dancinglights = spells.DancingLights{}
            s.Spells = append(s.Spells, dancinglights)
        case "Detect Evil and Good":
            var detectevilandgood = spells.DetectEvilAndGood{}
            s.Spells = append(s.Spells, detectevilandgood)
        case "Detect Magic":
            var detectmagic = spells.DetectMagic{}
            s.Spells = append(s.Spells, detectmagic)
        case "Detect Poison and Disease":
            var detectpoisonanddisease = spells.DetectPoisonAndDisease{}
            s.Spells = append(s.Spells, detectpoisonanddisease)
        case "Disguise Self":
            var disguiseself = spells.DisguiseSelf{}
            s.Spells = append(s.Spells, disguiseself)
        case "Dissonant Whispers":
            var dissonantwhispers = spells.DissonantWhispers{}
            s.Spells = append(s.Spells, dissonantwhispers)
        case "Divine Favor":
            var divinefavor = spells.DivineFavor{}
            s.Spells = append(s.Spells, divinefavor)
        case "Druidcraft":
            var druidcraft = spells.Druidcraft{}
            s.Spells = append(s.Spells, druidcraft)
        case "Eldritch Blast":
            var eldritchblast = spells.EldritchBlast{}
            s.Spells = append(s.Spells, eldritchblast)
        case "Ensnaring Strike":
            var ensnaringstrike = spells.EnsnaringStrike{}
            s.Spells = append(s.Spells, ensnaringstrike)
        case "Entangle":
            var entangle = spells.Entangle{}
            s.Spells = append(s.Spells, entangle)
        case "Expeditious Retreat":
            var expeditiousretreat = spells.ExpeditiousRetreat{}
            s.Spells = append(s.Spells, expeditiousretreat)
        case "Faerie Fire":
            var faeriefire = spells.FaerieFire{}
            s.Spells = append(s.Spells, faeriefire)
        case "False Life":
            var falselife = spells.FalseLife{}
            s.Spells = append(s.Spells, falselife)
        case "Feather Fall":
            var featherfall = spells.FeatherFall{}
            s.Spells = append(s.Spells, featherfall)
        case "Find Familiar":
            var findfamiliar = spells.FindFamiliar{}
            s.Spells = append(s.Spells, findfamiliar)
        case "Fire Bolt":
            var firebolt = spells.FireBolt{}
            s.Spells = append(s.Spells, firebolt)
        case "Fog Cloud":
            var fogcloud = spells.FogCloud{}
            s.Spells = append(s.Spells, fogcloud)
        case "Friends":
            var friends = spells.Friends{}
            s.Spells = append(s.Spells, friends)
        case "Goodberry":
            var goodberry = spells.Goodberry{}
            s.Spells = append(s.Spells, goodberry)
        case "Grease":
            var grease = spells.Grease{}
            s.Spells = append(s.Spells, grease)
        case "Guidance":
            var guidance = spells.Guidance{}
            s.Spells = append(s.Spells, guidance)
        case "Guiding Bolt":
            var guidingbolt = spells.GuidingBolt{}
            s.Spells = append(s.Spells, guidingbolt)
        case "Hail of Thorns":
            var hailofthorns = spells.HailOfThorns{}
            s.Spells = append(s.Spells, hailofthorns)
        case "Healing Word":
            var healingword = spells.HealingWord{}
            s.Spells = append(s.Spells, healingword)
        case "Hellish Rebuke":
            var hellishrebuke = spells.HellishRebuke{}
            s.Spells = append(s.Spells, hellishrebuke)
        case "Heroism":
            var heroism = spells.Heroism{}
            s.Spells = append(s.Spells, heroism)
        case "Hex":
            var hex = spells.Hex{}
            s.Spells = append(s.Spells, hex)
        case "Hunter's Mark":
            var huntersmark = spells.HuntersMark{}
            s.Spells = append(s.Spells, huntersmark)
        case "Identify":
            var identify = spells.Identify{}
            s.Spells = append(s.Spells, identify)
        case "Illusory Script":
            var illusoryscript = spells.IllusoryScript{}
            s.Spells = append(s.Spells, illusoryscript)
        case "Inflict Wounds":
            var inflictwounds = spells.InflictWounds{}
            s.Spells = append(s.Spells, inflictwounds)
        case "Jump":
            var jump = spells.Jump{}
            s.Spells = append(s.Spells, jump)
        case "Light":
            var light = spells.Light{}
            s.Spells = append(s.Spells, light)
        case "Longstrider":
            var longstrider = spells.Longstrider{}
            s.Spells = append(s.Spells, longstrider)
        case "Mage Armor":
            var magearmor = spells.MageArmor{}
            s.Spells = append(s.Spells, magearmor)
        case "Mage Hand":
            var magehand = spells.MageHand{}
            s.Spells = append(s.Spells, magehand)
        case "Magic Missile":
            var magicmissile = spells.MagicMissile{}
            s.Spells = append(s.Spells, magicmissile)
        case "Mending":
            var mending = spells.Mending{}
            s.Spells = append(s.Spells, mending)
        case "Message":
            var message = spells.Message{}
            s.Spells = append(s.Spells, message)
        case "Minor Illusion":
            var minorillusion = spells.MinorIllusion{}
            s.Spells = append(s.Spells, minorillusion)
        case "Poison Spray":
            var poisonspray = spells.PoisonSpray{}
            s.Spells = append(s.Spells, poisonspray)
        case "Prestidigitation":
            var prestidigitation = spells.Prestidigitation{}
            s.Spells = append(s.Spells, prestidigitation)
        case "Produce Flame":
            var produceflame = spells.ProduceFlame{}
            s.Spells = append(s.Spells, produceflame)
        case "Protection from Good and Evil":
            var protectionfromgoodandevil = spells.ProtectionFromGoodAndEvil{}
            s.Spells = append(s.Spells, protectionfromgoodandevil)
        case "Purify Food and Drink":
            var purifyfoodanddrink = spells.PurifyFoodAndDrink{}
            s.Spells = append(s.Spells, purifyfoodanddrink)
        case "Ray of Frost":
            var rayoffrost = spells.RayOfFrost{}
            s.Spells = append(s.Spells, rayoffrost)
        case "Ray of Sickness":
            var rayofsickness = spells.RayOfSickness{}
            s.Spells = append(s.Spells, rayofsickness)
        case "Resistance":
            var resistance = spells.Resistance{}
            s.Spells = append(s.Spells, resistance)
        case "Sacred Flame":
            var sacredflame = spells.SacredFlame{}
            s.Spells = append(s.Spells, sacredflame)
        case "Sanctuary":
            var sanctuary = spells.Sanctuary{}
            s.Spells = append(s.Spells, sanctuary)
        case "Searing Smite":
            var searingsmite = spells.SearingSmite{}
            s.Spells = append(s.Spells, searingsmite)
        case "Shield":
            var shield = spells.Shield{}
            s.Spells = append(s.Spells, shield)
        case "Shield of Faith":
            var shieldoffaith = spells.ShieldOfFaith{}
            s.Spells = append(s.Spells, shieldoffaith)
        case "Shillelagh":
            var shillelagh = spells.Shillelagh{}
            s.Spells = append(s.Spells, shillelagh)
        case "Shocking Grasp":
            var shockinggrasp = spells.ShockingGrasp{}
            s.Spells = append(s.Spells, shockinggrasp)
        case "Silent Image":
            var silentimage = spells.SilentImage{}
            s.Spells = append(s.Spells, silentimage)
        case "Sleep":
            var sleep = spells.Sleep{}
            s.Spells = append(s.Spells, sleep)
        case "Spare the Dying":
            var sparethedying = spells.SpareTheDying{}
            s.Spells = append(s.Spells, sparethedying)
        case "Speak with Animals":
            var speakwithanimals = spells.SpeakWithAnimals{}
            s.Spells = append(s.Spells, speakwithanimals)
        case "Tasha's Hideous Laughter":
            var tashashideouslaughter = spells.TashasHideousLaughter{}
            s.Spells = append(s.Spells, tashashideouslaughter)
        case "Tenser's Floating Disk":
            var tensersfloatingdisk = spells.TensersFloatingDisk{}
            s.Spells = append(s.Spells, tensersfloatingdisk)
        case "Thaumaturgy":
            var thaumaturgy = spells.Thaumaturgy{}
            s.Spells = append(s.Spells, thaumaturgy)
        case "Thorn Whip":
            var thornwhip = spells.ThornWhip{}
            s.Spells = append(s.Spells, thornwhip)
        case "Thunderous Smite":
            var thunderoussmite = spells.ThunderousSmite{}
            s.Spells = append(s.Spells, thunderoussmite)
        case "Thunderwave":
            var thunderwave = spells.Thunderwave{}
            s.Spells = append(s.Spells, thunderwave)
        case "True Strike":
            var truestrike = spells.TrueStrike{}
            s.Spells = append(s.Spells, truestrike)
        case "Unseen Servant":
            var unseenservant = spells.UnseenServant{}
            s.Spells = append(s.Spells, unseenservant)
        case "Vicious Mockery":
            var viciousmockery = spells.ViciousMockery{}
            s.Spells = append(s.Spells, viciousmockery)
        case "Witch Bolt":
            var witchbolt = spells.WitchBolt{}
            s.Spells = append(s.Spells, witchbolt)
        case "Wrathful Smite":
            var wrathfulsmite = spells.WrathfulSmite{}
            s.Spells = append(s.Spells, wrathfulsmite)
        default:
            return errors.New(fmt.Sprintf("%s is not a valid spell", spell))
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
    var illuminatedstr string
    for iind, illum := range s.Illuminated {
        if iind == len(s.Illuminated) - 1 {
            illuminatedstr += strconv.Itoa(illum)
        } else {
            illuminatedstr += strconv.Itoa(illum) + ","
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
           "Ancestry:" + s.Ancestry + ";" +
           "Illuminated:" + illuminatedstr + ";" +
           "Oiled:" + strconv.Itoa(s.Oiled) + ";"
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

func (p *Player) Effects(action string, data []int, c chan int) {
    switch action {
    case "illuminate":
        if len(data) != 3 {
            log.Fatal(fmt.Sprintf("Incorrect # of arguments %d for illuminate data", len(data)))
        }
        p.Stats.Illuminated = data
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
        default:
            log.Println("Not a valid throwable item")
        }
        msg := <-c
        if origitem != nil {
            p.Equip(origitem)
        }
        //if msg == 0 || msg == 1 {
        switch msg {
        case 0:
            p.Inv.GetItems()[len(p.Inv.GetItems()) - 1].(*items.OilFlask).Quantity--
            return
        case 1:
            return
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

func (p *Player) CastSpell(spell spells.Spell, target *Player) {
    log.Println(fmt.Sprintf("Casting spell %s, target is %s", spell.PrettyPrint(), target.GetName()))
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
        p.Stats.Illuminated = []int{}
        if strings.HasPrefix(p.Equipment.LeftHand.PrettyPrint(), "Candles") {
            if p.Equipment.LeftHand.(*items.Candles).Turns == 0 {
                p.Equipment.LeftHand.(*items.Candles).Turns = 600
            }
        } else if strings.HasPrefix(p.Equipment.LeftHand.PrettyPrint(), "Lamp") {
            if p.Equipment.LeftHand.(*items.Lamp).Turns == 0 {
                p.Equipment.LeftHand.(*items.Lamp).Turns = 3600
            }
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
                    p.Stats.Illuminated = []int{}
                    if strings.HasPrefix(p.Equipment.LeftHand.PrettyPrint(), "Candles") {
                        if p.Equipment.LeftHand.(*items.Candles).Turns == 0 {
                            p.Equipment.LeftHand.(*items.Candles).Turns = 600
                        }
                    } else if strings.HasPrefix(p.Equipment.LeftHand.PrettyPrint(), "Lamp") {
                        if p.Equipment.LeftHand.(*items.Lamp).Turns == 0 {
                            p.Equipment.LeftHand.(*items.Lamp).Turns = 3600
                        }
                    }
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
