package items

import (
    "fmt"
    "log"
    "strconv"

    "github.com/jsnider-mtu/quailgame/inventory"
)

func LoadItem(name, material, durability string) inventory.Item {
    dur, err := strconv.Atoi(durability)
    if err != nil {
        log.Fatal(err)
    }
    switch name {
    case "Pickaxe":
        return Pickaxe{material: material, durability: dur}
    //case "Axe":
    default:
        log.Fatal(fmt.Sprintf("Item with name %s does not exist", name))
        return nil
    }
}
