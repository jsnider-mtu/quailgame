package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Thaumaturgy struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (t Thaumaturgy) PrettyPrint() string {
    return "Thaumaturgy"
}

func (t Thaumaturgy) GetLevel() int {
    return t.level
}

func (t Thaumaturgy) GetCastTime() string {
    return t.casttime
}

func (t Thaumaturgy) GetRange() int {
    return t.range
}

func (t Thaumaturgy) GetComponents() string {
    return t.components
}

func (t Thaumaturgy) GetDuration() int {
    return t.duration
}

func (t Thaumaturgy) GetConcentration() bool {
    return t.concentration
}

func (t Thaumaturgy) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thaumaturgy is not implemented yet")
}
