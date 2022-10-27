package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type BurningHands struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (b BurningHands) PrettyPrint() string {
    return "Burning Hands"
}

func (b BurningHands) GetLevel() int {
    return b.level
}

func (b BurningHands) GetCastTime() string {
    return b.casttime
}

func (b BurningHands) GetRange() float64 {
    return float64(b.spellrange) * 4.8
}

func (b BurningHands) GetComponents() string {
    return b.components
}

func (b BurningHands) GetDuration() int {
    return b.duration
}

func (b BurningHands) GetConcentration() bool {
    return b.concentration
}

func (b BurningHands) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Burning Hands is not implemented yet")
}
