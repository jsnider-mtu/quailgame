package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Goodberry struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (g Goodberry) PrettyPrint() string {
    return "Goodberry"
}

func (g Goodberry) GetLevel() int {
    return g.level
}

func (g Goodberry) GetCastTime() string {
    return g.casttime
}

func (g Goodberry) GetRange() int {
    return g.range
}

func (g Goodberry) GetComponents() string {
    return g.components
}

func (g Goodberry) GetDuration() int {
    return g.duration
}

func (g Goodberry) GetConcentration() bool {
    return g.concentration
}

func (g Goodberry) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Goodberry is not implemented yet")
}
