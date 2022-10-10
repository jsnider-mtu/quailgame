package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Grease struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (g Grease) PrettyPrint() string {
    return "Grease"
}

func (g Grease) GetLevel() int {
    return g.level
}

func (g Grease) GetCastTime() string {
    return g.casttime
}

func (g Grease) GetRange() int {
    return g.range
}

func (g Grease) GetComponents() string {
    return g.components
}

func (g Grease) GetDuration() int {
    return g.duration
}

func (g Grease) GetConcentration() bool {
    return g.concentration
}

func (g Grease) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Grease is not implemented yet")
}
