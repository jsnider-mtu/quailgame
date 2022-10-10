package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CureWounds struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (c CureWounds) PrettyPrint() string {
    return "Cure Wounds"
}

func (c CureWounds) GetLevel() int {
    return c.level
}

func (c CureWounds) GetCastTime() string {
    return c.casttime
}

func (c CureWounds) GetRange() int {
    return c.range
}

func (c CureWounds) GetComponents() string {
    return c.components
}

func (c CureWounds) GetDuration() int {
    return c.duration
}

func (c CureWounds) GetConcentration() bool {
    return c.concentration
}

func (c CureWounds) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Cure Wounds is not implemented yet")
}
