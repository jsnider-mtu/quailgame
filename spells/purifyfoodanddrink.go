package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type PurifyFoodAndDrink struct {}

func (p PurifyFoodAndDrink) Cast(target *npcs.NPC) bool {
    log.Println("The spell Purify Food and Drink is not implemented yet")
}

func (p PurifyFoodAndDrink) PrettyPrint() string {
    return "Purify Food and Drink"
}

func (a PurifyFoodAndDrink) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Purify Food and Drink is not implemented yet")
}
