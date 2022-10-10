package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Druidcraft struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (d Druidcraft) PrettyPrint() string {
    return "Druidcraft"
}

func (d Druidcraft) GetLevel() int {
    return d.level
}

func (d Druidcraft) GetCastTime() string {
    return d.casttime
}

func (d Druidcraft) GetRange() int {
    return d.spellrange
}

func (d Druidcraft) GetComponents() string {
    return d.components
}

func (d Druidcraft) GetDuration() int {
    return d.duration
}

func (d Druidcraft) GetConcentration() bool {
    return d.concentration
}

func (d Druidcraft) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Druidcraft is not implemented yet")
}
