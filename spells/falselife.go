package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FalseLife struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (f FalseLife) PrettyPrint() string {
    return "False Life"
}

func (f FalseLife) GetLevel() int {
    return f.level
}

func (f FalseLife) GetCastTime() string {
    return f.casttime
}

func (f FalseLife) GetRange() int {
    return f.range
}

func (f FalseLife) GetComponents() string {
    return f.components
}

func (f FalseLife) GetDuration() int {
    return f.duration
}

func (f FalseLife) GetConcentration() bool {
    return f.concentration
}

func (f FalseLife) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell False Life is not implemented yet")
}
