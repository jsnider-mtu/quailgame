package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Identify struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (i Identify) PrettyPrint() string {
    return "Identify"
}

func (i Identify) GetLevel() int {
    return i.level
}

func (i Identify) GetCastTime() string {
    return i.casttime
}

func (i Identify) GetRange() int {
    return i.spellrange
}

func (i Identify) GetComponents() string {
    return i.components
}

func (i Identify) GetDuration() int {
    return i.duration
}

func (i Identify) GetConcentration() bool {
    return i.concentration
}

func (i Identify) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Identify is not implemented yet")
}
