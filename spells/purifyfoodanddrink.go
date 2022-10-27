package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type PurifyFoodAndDrink struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (p PurifyFoodAndDrink) PrettyPrint() string {
    return "Purify Food and Drink"
}

func (p PurifyFoodAndDrink) GetLevel() int {
    return p.level
}

func (p PurifyFoodAndDrink) GetCastTime() string {
    return p.casttime
}

func (p PurifyFoodAndDrink) GetRange() float64 {
    return float64(p.spellrange) * 4.8
}

func (p PurifyFoodAndDrink) GetComponents() string {
    return p.components
}

func (p PurifyFoodAndDrink) GetDuration() int {
    return p.duration
}

func (p PurifyFoodAndDrink) GetConcentration() bool {
    return p.concentration
}

func (p PurifyFoodAndDrink) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Purify Food and Drink is not implemented yet")
}
