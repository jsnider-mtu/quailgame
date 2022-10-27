package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FaerieFire struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (f FaerieFire) PrettyPrint() string {
    return "Faerie Fire"
}

func (f FaerieFire) GetLevel() int {
    return f.level
}

func (f FaerieFire) GetCastTime() string {
    return f.casttime
}

func (f FaerieFire) GetRange() float64 {
    return float64(f.spellrange) * 4.8
}

func (f FaerieFire) GetComponents() string {
    return f.components
}

func (f FaerieFire) GetDuration() int {
    return f.duration
}

func (f FaerieFire) GetConcentration() bool {
    return f.concentration
}

func (f FaerieFire) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Faerie Fire is not implemented yet")
}
