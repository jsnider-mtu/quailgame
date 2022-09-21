package items

import (
    "fmt"
    "log"
    "strconv"

    "github.com/jsnider-mtu/quailgame/inventory"
)

func LoadItem(name string, quaity interface{}) inventory.Item {
    switch name {
    case "Arcanefocus":
        return Arcanefocus{}
    case "Bagpipes":
        return Bagpipes{}
    case "Battleaxe":
        return Battleaxe{}
    case "Blowgun":
        return Blowgun{}
    case "Candles":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return Candles{Quantity: quaityint}
    case "Chainmail":
        return Chainmail{}
    case "Clothes":
        return Clothes{Quality: quaity.(string)}
    case "Club":
        return Club{}
    case "Componentpouch":
        return Componentpouch{}
    case "Dagger":
        return Dagger{}
    case "Darts":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return Darts{Quantity: quaityint}
    case "Disguisekit":
        return Disguisekit{}
    case "Drum":
        return Drum{}
    case "Dulcimer":
        return Dulcimer{}
    case "Flail":
        return Flail{}
    case "Flute":
        return Flute{}
    case "Glaive":
        return Glaive{}
    case "Greataxe":
        return Greataxe{}
    case "Greatclub":
        return Greatclub{}
    case "Greatsword":
        return Greatsword{}
    case "Halberd":
        return Halberd{}
    case "Handaxe":
        return Handaxe{}
    case "Handcrossbow":
        return Handcrossbow{}
    case "Heavycrossbow":
        return Heavycrossbow{}
    case "Horn":
        return Horn{}
    case "Inkbottle":
        return Inkbottle{}
    case "Inkpen":
        return Inkpen{}
    case "Javelin":
        return Javelin{}
    case "Lamp":
        return Lamp{}
    case "Lance":
        return Lance{}
    case "Leatherarmor":
        return Leatherarmor{}
    case "Lightcrossbow":
        return Lightcrossbow{}
    case "Lighthammer":
        return Lighthammer{}
    case "Longbow":
        return Longbow{}
    case "Longsword":
        return Longsword{}
    case "Lute":
        return Lute{}
    case "Lyre":
        return Lyre{}
    case "Mace":
        return Mace{}
    case "Maul":
        return Maul{}
    case "Morningstar":
        return Morningstar{}
    case "Net":
        return Net{}
    case "Oilflask":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return Oilflask{Quantity: quaityint}
    case "Panflute":
        return Panflute{}
    case "Paper":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return Paper{Quantity: quaityint}
    case "Perfume":
        return Perfume{}
    case "Pike":
        return Pike{}
    case "Quarterstaff":
        return Quarterstaff{}
    case "Quiver":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return Quiver{Arrows: quaityint}
    case "Rapier":
        return Rapier{}
    case "Rope":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return Rope{Length: quaityint}
    case "Scalemail":
        return Scalemail{}
    case "Scimitar":
        return Scimitar{}
    case "Sealingwax":
        return Sealingwax{}
    case "Shawm":
        return Shawm{}
    case "Shield":
        return Shield{}
    case "Shortbow":
        return Shortbow{}
    case "Shortsword":
        return Shortsword{}
    case "Sickle":
        return Sickle{}
    case "Sling":
        return Sling{}
    case "Soap":
        return Soap{}
    case "Spear":
        return Spear{}
    case "Thievestools":
        return Thievestools{}
    case "Tinderbox":
        return Tinderbox{}
    case "Torches":
        quaityint, err := strconv.Atoi(quaity.(string))
        if err != nil {
            log.Fatal("quaity is not an integer")
        }
        return Torches{Quantity: quaityint}
    case "Trident":
        return Trident{}
    case "Viol":
        return Viol{}
    case "Warhammer":
        return Warhammer{}
    case "Warpick":
        return Warpick{}
    case "Whip":
        return Whip{}
    default:
        log.Fatal(fmt.Sprintf("Item with name %s does not exist", name))
        return nil
    }
}
