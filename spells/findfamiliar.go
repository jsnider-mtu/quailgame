package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FindFamiliar struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (f FindFamiliar) PrettyPrint() string {
    return "Find Familiar"
}

func (f FindFamiliar) GetLevel() int {
    return f.level
}

func (f FindFamiliar) GetCastTime() string {
    return f.casttime
}

func (f FindFamiliar) GetRange() int {
    return f.range
}

func (f FindFamiliar) GetComponents() string {
    return f.components
}

func (f FindFamiliar) GetDuration() int {
    return f.duration
}

func (f FindFamiliar) GetConcentration() bool {
    return f.concentration
}

func (f FindFamiliar) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Find Familiar is not implemented yet")
}
