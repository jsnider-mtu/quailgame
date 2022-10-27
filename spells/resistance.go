package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Resistance struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (r Resistance) PrettyPrint() string {
    return "Resistance"
}

func (r Resistance) GetLevel() int {
    return r.level
}

func (r Resistance) GetCastTime() string {
    return r.casttime
}

func (r Resistance) GetRange() float64 {
    return float64(r.spellrange) * 4.8
}

func (r Resistance) GetComponents() string {
    return r.components
}

func (r Resistance) GetDuration() int {
    return r.duration
}

func (r Resistance) GetConcentration() bool {
    return r.concentration
}

func (r Resistance) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Resistance is not implemented yet")
}
