package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type PurifyFoodAndDrink struct {}

func (p PurifyFoodAndDrink) Cast(target *npcs.NPC) bool {
    log.Println("The spell Purify Food and Drink is not implemented yet")
}

func (p PurifyFoodAndDrink) PrettyPrint() string {
    return "Purify Food and Drink"
}
