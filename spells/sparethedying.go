package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SpareTheDying struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s SpareTheDying) PrettyPrint() string {
    return "Spare the Dying"
}

func (s SpareTheDying) GetLevel() int {
    return s.level
}

func (s SpareTheDying) GetCastTime() string {
    return s.casttime
}

func (s SpareTheDying) GetRange() float64 {
    return float64(s.spellrange) * 4.8
}

func (s SpareTheDying) GetComponents() string {
    return s.components
}

func (s SpareTheDying) GetDuration() int {
    return s.duration
}

func (s SpareTheDying) GetConcentration() bool {
    return s.concentration
}

func (s SpareTheDying) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Spare the Dying is not implemented yet")
}
