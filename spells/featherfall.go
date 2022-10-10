package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FeatherFall struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (f FeatherFall) PrettyPrint() string {
    return "Feather Fall"
}

func (f FeatherFall) GetLevel() int {
    return f.level
}

func (f FeatherFall) GetCastTime() string {
    return f.casttime
}

func (f FeatherFall) GetRange() int {
    return f.range
}

func (f FeatherFall) GetComponents() string {
    return f.components
}

func (f FeatherFall) GetDuration() int {
    return f.duration
}

func (f FeatherFall) GetConcentration() bool {
    return f.concentration
}

func (f FeatherFall) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Feather Fall is not implemented yet")
}
