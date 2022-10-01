package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type PurifyFoodAndDrink struct {}

func (p PurifyFoodAndDrink) PrettyPrint() string {
    return "Purify Food and Drink"
}

func (a PurifyFoodAndDrink) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Purify Food and Drink is not implemented yet")
}
