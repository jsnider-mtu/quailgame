package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Prestidigitation struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (p Prestidigitation) PrettyPrint() string {
    return "Prestidigitation"
}

func (p Prestidigitation) GetLevel() int {
    return p.level
}

func (p Prestidigitation) GetCastTime() string {
    return p.casttime
}

func (p Prestidigitation) GetRange() int {
    return p.range
}

func (p Prestidigitation) GetComponents() string {
    return p.components
}

func (p Prestidigitation) GetDuration() int {
    return p.duration
}

func (p Prestidigitation) GetConcentration() bool {
    return p.concentration
}

func (p Prestidigitation) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Prestidigitation is not implemented yet")
}
