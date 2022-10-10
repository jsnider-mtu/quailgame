package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ChillTouch struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (c ChillTouch) PrettyPrint() string {
    return "Chill Touch"
}

func (c ChillTouch) GetLevel() int {
    return c.level
}

func (c ChillTouch) GetCastTime() string {
    return c.casttime
}

func (c ChillTouch) GetRange() int {
    return c.spellrange
}

func (c ChillTouch) GetComponents() string {
    return c.components
}

func (c ChillTouch) GetDuration() int {
    return c.duration
}

func (c ChillTouch) GetConcentration() bool {
    return c.concentration
}

func (c ChillTouch) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Chill Touch is not implemented yet")
}
