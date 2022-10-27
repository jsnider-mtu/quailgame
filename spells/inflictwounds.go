package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type InflictWounds struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (i InflictWounds) PrettyPrint() string {
    return "Inflict Wounds"
}

func (i InflictWounds) GetLevel() int {
    return i.level
}

func (i InflictWounds) GetCastTime() string {
    return i.casttime
}

func (i InflictWounds) GetRange() float64 {
    return float64(i.spellrange) * 4.8
}

func (i InflictWounds) GetComponents() string {
    return i.components
}

func (i InflictWounds) GetDuration() int {
    return i.duration
}

func (i InflictWounds) GetConcentration() bool {
    return i.concentration
}

func (i InflictWounds) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Inflict Wounds is not implemented yet")
}
