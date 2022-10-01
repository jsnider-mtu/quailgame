package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CureWounds struct {}

func (c CureWounds) PrettyPrint() string {
    return "Cure Wounds"
}

func (c CureWounds) GetLevel() int {
    return 0
}

func (a CureWounds) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Cure Wounds is not implemented yet")
}
