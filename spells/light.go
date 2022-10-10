package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Light struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (l Light) PrettyPrint() string {
    return "Light"
}

func (l Light) GetLevel() int {
    return l.level
}

func (l Light) GetCastTime() string {
    return l.casttime
}

func (l Light) GetRange() int {
    return l.range
}

func (l Light) GetComponents() string {
    return l.components
}

func (l Light) GetDuration() int {
    return l.duration
}

func (l Light) GetConcentration() bool {
    return l.concentration
}

func (l Light) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Light is not implemented yet")
}
