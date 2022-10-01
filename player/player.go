package player

import (
    "errors"
    "fmt"
    "log"
    "strconv"

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
}

func (s *Spells) Add(spellsslice []string) error {
    for _, spell := range spellsslice {
        switch spell {
        case "Acid Splash":
            var acidsplash spells.AcidSplash
            s.Spells = append(s.Spells, acidsplash)
        case "Alarm":
            var alarm spells.Alarm
            s.Spells = append(s.Spells, alarm)
        case "Animal Friendship":
            var animalfriendship spells.AnimalFriendship
            s.Spells = append(s.Spells, animalfriendship)
        case "Armor of Agathys":
            var armorofagathys spells.ArmorOfAgathys
            s.Spells = append(s.Spells, armorofagathys)
        case "Arms of Hadar":
            var armsofhadar spells.ArmsOfHadar
            s.Spells = append(s.Spells, armsofhadar)
        case "Bane":
            var bane spells.Bane
            s.Spells = append(s.Spells, bane)
        case "Blade Ward":
            var bladeward spells.BladeWard
            s.Spells = append(s.Spells, bladeward)
        case "Bless":
            var bless spells.Bless
            s.Spells = append(s.Spells, bless)
        case "Burning Hands":
            var burninghands spells.BurningHands
            s.Spells = append(s.Spells, burninghands)
        case "Charm Person":
            var charmperson spells.CharmPerson
            s.Spells = append(s.Spells, charmperson)
        case "Chill Touch":
            var chilltouch spells.ChillTouch
            s.Spells = append(s.Spells, chilltouch)
        case "Chromatic Orb":
            var chromaticorb spells.ChromaticOrb
            s.Spells = append(s.Spells, chromaticorb)
        case "Color Spray":
            var colorspray spells.ColorSpray
            s.Spells = append(s.Spells, colorspray)
        case "Command":
            var command spells.Command
            s.Spells = append(s.Spells, command)
        case "Compelled Duel":
            var compelledduel spells.CompelledDuel
            s.Spells = append(s.Spells, compelledduel)
        case "Comprehend Languages":
            var comprehendlanguages spells.ComprehendLanguages
            s.Spells = append(s.Spells, comprehendlanguages)
        case "Create or Destroy Water":
            var createordestroywater spells.CreateOrDestroyWater
            s.Spells = append(s.Spells, createordestroywater)
        case "Cure Wounds":
            var curewounds spells.CureWounds
            s.Spells = append(s.Spells, curewounds)
        case "Dancing Lights":
            var dancinglights spells.DancingLights
            s.Spells = append(s.Spells, dancinglights)
        case "Detect Evil and Good":
            var detectevilandgood spells.DetectEvilAndGood
            s.Spells = append(s.Spells, detectevilandgood)
        case "Detect Magic":
            var detectmagic spells.DetectMagic
            s.Spells = append(s.Spells, detectmagic)
        case "Detect Poison and Disease":
            var detectpoisonanddisease spells.DetectPoisonAndDisease
            s.Spells = append(s.Spells, detectpoisonanddisease)
        case "Disguise Self":
            var disguiseself spells.DisguiseSelf
            s.Spells = append(s.Spells, disguiseself)
        case "Dissonant Whispers":
            var dissonantwhispers spells.DissonantWhispers
            s.Spells = append(s.Spells, dissonantwhispers)
        case "Divine Favor":
            var divinefavor spells.DivineFavor
            s.Spells = append(s.Spells, divinefavor)
        case "Druidcraft":
            var druidcraft spells.Druidcraft
            s.Spells = append(s.Spells, druidcraft)
        case "Eldritch Blast":
            var eldritchblast spells.EldritchBlast
            s.Spells = append(s.Spells, eldritchblast)
        case "Ensnaring Strike":
            var ensnaringstrike spells.EnsnaringStrike
            s.Spells = append(s.Spells, ensnaringstrike)
        case "Entangle":
            var entangle spells.Entangle
            s.Spells = append(s.Spells, entangle)
        case "Expeditious Retreat":
            var expeditiousretreat spells.ExpeditiousRetreat
            s.Spells = append(s.Spells, expeditiousretreat)
        case "Faerie Fire":
            var faeriefire spells.FaerieFire
            s.Spells = append(s.Spells, faeriefire)
        case "False Life":
            var falselife spells.FalseLife
            s.Spells = append(s.Spells, falselife)
        case "Feather Fall":
            var featherfall spells.FeatherFall
            s.Spells = append(s.Spells, featherfall)
        case "Find Familiar":
            var findfamiliar spells.FindFamiliar
            s.Spells = append(s.Spells, findfamiliar)
        case "Fire Bolt":
            var firebolt spells.FireBolt
            s.Spells = append(s.Spells, firebolt)
        case "Fog Cloud":
            var fogcloud spells.FogCloud
            s.Spells = append(s.Spells, fogcloud)
        case "Friends":
            var friends spells.Friends
            s.Spells = append(s.Spells, friends)
        case "Goodberry":
            var goodberry spells.Goodberry
            s.Spells = append(s.Spells, goodberry)
        case "Grease":
            var grease spells.Grease
            s.Spells = append(s.Spells, grease)
        case "Guidance":
            var guidance spells.Guidance
            s.Spells = append(s.Spells, guidance)
        case "Guiding Bolt":
            var guidingbolt spells.GuidingBolt
            s.Spells = append(s.Spells, guidingbolt)
        case "Hail of Thorns":
            var hailofthorns spells.HailOfThorns
            s.Spells = append(s.Spells, hailofthorns)
        case "Healing Word":
            var healingword spells.HealingWord
            s.Spells = append(s.Spells, healingword)
        case "Hellish Rebuke":
            var hellishrebuke spells.HellishRebuke
            s.Spells = append(s.Spells, hellishrebuke)
        case "Heroism":
            var heroism spells.Heroism
            s.Spells = append(s.Spells, heroism)
        case "Hex":
            var hex spells.Hex
            s.Spells = append(s.Spells, hex)
        case "Hunter's Mark":
            var huntersmark spells.HuntersMark
            s.Spells = append(s.Spells, huntersmark)
        case "Identify":
            var identify spells.Identify
            s.Spells = append(s.Spells, identify)
        case "Illusory Script":
            var illusoryscript spells.IllusoryScript
            s.Spells = append(s.Spells, illusoryscript)
        case "Inflict Wounds":
            var inflictwounds spells.InflictWounds
            s.Spells = append(s.Spells, inflictwounds)
        case "Jump":
            var jump spells.Jump
            s.Spells = append(s.Spells, jump)
        case "Light":
            var light spells.Light
            s.Spells = append(s.Spells, light)
        case "Longstrider":
            var longstrider spells.Longstrider
            s.Spells = append(s.Spells, longstrider)
        case "Mage Armor":
            var magearmor spells.MageArmor
            s.Spells = append(s.Spells, magearmor)
        case "Mage Hand":
            var magehand spells.MageHand
            s.Spells = append(s.Spells, magehand)
        case "Magic Missile":
            var magicmissile spells.MagicMissile
            s.Spells = append(s.Spells, magicmissile)
        case "Mending":
            var mending spells.Mending
            s.Spells = append(s.Spells, mending)
        case "Message":
            var message spells.Message
            s.Spells = append(s.Spells, message)
        case "Minor Illusion":
            var minorillusion spells.MinorIllusion
            s.Spells = append(s.Spells, minorillusion)
        case "Poison Spray":
            var poisonspray spells.PoisonSpray
            s.Spells = append(s.Spells, poisonspray)
        case "Prestidigitation":
            var prestidigitation spells.Prestidigitation
            s.Spells = append(s.Spells, prestidigitation)
        case "Produce Flame":
            var produceflame spells.ProduceFlame
            s.Spells = append(s.Spells, produceflame)
        case "Protection from Good and Evil":
            var protectionfromgoodandevil spells.ProtectionFromGoodAndEvil
            s.Spells = append(s.Spells, protectionfromgoodandevil)
        case "Purify Food and Drink":
            var purifyfoodanddrink spells.PurifyFoodAndDrink
            s.Spells = append(s.Spells, purifyfoodanddrink)
        case "Ray of Frost":
            var rayoffrost spells.RayOfFrost
            s.Spells = append(s.Spells, rayoffrost)
        case "Ray of Sickness":
            var rayofsickness spells.RayOfSickness
            s.Spells = append(s.Spells, rayofsickness)
        case "Resistance":
            var resistance spells.Resistance
            s.Spells = append(s.Spells, resistance)
        case "Sacred Flame":
            var sacredflame spells.SacredFlame
            s.Spells = append(s.Spells, sacredflame)
        case "Sanctuary":
            var sanctuary spells.Sanctuary
            s.Spells = append(s.Spells, sanctuary)
        case "Searing Smite":
            var searingsmite spells.SearingSmite
            s.Spells = append(s.Spells, searingsmite)
        case "Shield":
            var shield spells.Shield
            s.Spells = append(s.Spells, shield)
        case "Shield of Faith":
            var shieldoffaith spells.ShieldOfFaith
            s.Spells = append(s.Spells, shieldoffaith)
        case "Shillelagh":
            var shillelagh spells.Shillelagh
            s.Spells = append(s.Spells, shillelagh)
        case "Shocking Grasp":
            var shockinggrasp spells.ShockingGrasp
            s.Spells = append(s.Spells, shockinggrasp)
        case "Silent Image":
            var silentimage spells.SilentImage
            s.Spells = append(s.Spells, silentimage)
        case "Sleep":
            var sleep spells.Sleep
            s.Spells = append(s.Spells, sleep)
        case "Spare the Dying":
            var sparethedying spells.SpareTheDying
            s.Spells = append(s.Spells, sparethedying)
        case "Speak with Animals":
            var speakwithanimals spells.SpeakWithAnimals
            s.Spells = append(s.Spells, speakwithanimals)
        case "Tasha's Hideous Laughter":
            var tashashideouslaughter spells.TashasHideousLaughter
            s.Spells = append(s.Spells, tashashideouslaughter)
        case "Tenser's Floating Disk":
            var tensersfloatingdisk spells.TensersFloatingDisk
            s.Spells = append(s.Spells, tensersfloatingdisk)
        case "Thaumaturgy":
            var thaumaturgy spells.Thaumaturgy
            s.Spells = append(s.Spells, thaumaturgy)
        case "Thorn Whip":
            var thornwhip spells.ThornWhip
            s.Spells = append(s.Spells, thornwhip)
        case "Thunderous Smite":
            var thunderoussmite spells.ThunderousSmite
            s.Spells = append(s.Spells, thunderoussmite)
        case "Thunderwave":
            var thunderwave spells.Thunderwave
            s.Spells = append(s.Spells, thunderwave)
        case "True Strike":
            var truestrike spells.TrueStrike
            s.Spells = append(s.Spells, truestrike)
        case "Unseen Servant":
            var unseenservant spells.UnseenServant
            s.Spells = append(s.Spells, unseenservant)
        case "Vicious Mockery":
            var viciousmockery spells.ViciousMockery
            s.Spells = append(s.Spells, viciousmockery)
        case "Witch Bolt":
            var witchbolt spells.WitchBolt
            s.Spells = append(s.Spells, witchbolt)
        case "Wrathful Smite":
            var wrathfulsmite spells.WrathfulSmite
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
    return ""
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
            default:
                log.Fatal(fmt.Sprintf("%s is not a valid slot", item.Slot))
            }
        }
    }
}
