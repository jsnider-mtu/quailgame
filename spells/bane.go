package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Bane struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (b Bane) PrettyPrint() string {
    return "Bane"
}

func (b Bane) GetLevel() int {
    return b.level
}

func (b Bane) GetCastTime() string {
    return b.casttime
}

func (b Bane) GetRange() float64 {
    return float64(b.spellrange) * 4.8
}

func (b Bane) GetComponents() string {
    return b.components
}

func (b Bane) GetDuration() int {
    return b.duration
}

func (b Bane) GetConcentration() bool {
    return b.concentration
}

func (a Bane) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Bane is not implemented yet")
}
