package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Mending struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (m Mending) PrettyPrint() string {
    return "Mending"
}

func (m Mending) GetLevel() int {
    return m.level
}

func (m Mending) GetCastTime() string {
    return m.casttime
}

func (m Mending) GetRange() int {
    return m.range
}

func (m Mending) GetComponents() string {
    return m.components
}

func (m Mending) GetDuration() int {
    return m.duration
}

func (m Mending) GetConcentration() bool {
    return m.concentration
}

func (m Mending) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Mending is not implemented yet")
}
