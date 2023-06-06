package items

import (
    "fmt"
    "log"
    "strconv"

    "github.com/jsnider-mtu/quailgame/inventory"
)

func LoadItem(name string, quaity interface{}) inventory.Item {
    switch name {
    case "ArcaneFocus":
        return &ArcaneFocus{}
    case "Bagpipes":
        return &Bagpipes{}
    case "Battleaxe":
        return &Battleaxe{}
    case "Blowgun":
        return &Blowgun{}
    case "Candles":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &Candles{Quantity: quaityint}
    case "Chainmail":
        return &Chainmail{}
    case "Club":
        return &Club{}
    case "ComponentPouch":
        return &ComponentPouch{}
    case "Dagger":
        return &Dagger{}
    case "Darts":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &Darts{Quantity: quaityint}
    case "DisguiseKit":
        return &DisguiseKit{}
    case "Drum":
        return &Drum{}
    case "Dulcimer":
        return &Dulcimer{}
    case "Flail":
        return &Flail{}
    case "Flute":
        return &Flute{}
    case "Glaive":
        return &Glaive{}
    case "Greataxe":
        return &Greataxe{}
    case "Greatclub":
        return &Greatclub{}
    case "Greatsword":
        return &Greatsword{}
    case "Halberd":
        return &Halberd{}
    case "Handaxe":
        return &Handaxe{}
    case "Handcrossbow":
        return &HandCrossbow{}
    case "Heavycrossbow":
        return &HeavyCrossbow{}
    case "Horn":
        return &Horn{}
    case "InkBottle":
        return &InkBottle{}
    case "InkPen":
        return &InkPen{}
    case "Javelin":
        return &Javelin{}
    case "Lamp":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &Lamp{Quantity: quaityint}
    case "Lance":
        return &Lance{}
    case "LeatherArmor":
        return &LeatherArmor{}
    case "LightCrossbow":
        return &LightCrossbow{}
    case "LightHammer":
        return &LightHammer{}
    case "Longbow":
        return &Longbow{}
    case "Longsword":
        return &Longsword{}
    case "Lute":
        return &Lute{}
    case "Lyre":
        return &Lyre{}
    case "Mace":
        return &Mace{}
    case "Maul":
        return &Maul{}
    case "Morningstar":
        return &Morningstar{}
    case "Net":
        return &Net{}
    case "OilFlask":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &OilFlask{Quantity: quaityint}
    case "PanFlute":
        return &PanFlute{}
    case "Paper":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &Paper{Quantity: quaityint}
    case "Perfume":
        return &Perfume{}
    case "Pike":
        return &Pike{}
    case "Quarterstaff":
        return &Quarterstaff{}
    case "Quiver":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &Quiver{Arrows: quaityint}
    case "Rapier":
        return &Rapier{}
    case "Rope":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &Rope{Length: quaityint}
    case "Scalemail":
        return &Scalemail{}
    case "Scimitar":
        return &Scimitar{}
    case "SealingWax":
        return &SealingWax{}
    case "Shawm":
        return &Shawm{}
    case "Shield":
        return &Shield{}
    case "Shortbow":
        return &Shortbow{}
    case "Shortsword":
        return &Shortsword{}
    case "Sickle":
        return &Sickle{}
    case "Sling":
        return &Sling{}
    case "Soap":
        return &Soap{}
    case "Spear":
        return &Spear{}
    case "ThievesTools":
        return &ThievesTools{}
    case "Tinderbox":
        return &Tinderbox{}
    case "Torches":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return &Torches{Quantity: quaityint}
    case "Trident":
        return &Trident{}
    case "Viol":
        return &Viol{}
    case "Warhammer":
        return &Warhammer{}
    case "WarPick":
        return &WarPick{}
    case "Whip":
        return &Whip{}
    default:
        log.Fatal(fmt.Sprintf("Item with name %s does not exist", name))
        return nil
    }
}
