package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ChromaticOrb struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (c ChromaticOrb) PrettyPrint() string {
    return "Chromatic Orb"
}

func (c ChromaticOrb) GetLevel() int {
    return c.level
}

func (c ChromaticOrb) GetCastTime() string {
    return c.casttime
}

func (c ChromaticOrb) GetRange() int {
    return c.range
}

func (c ChromaticOrb) GetComponents() string {
    return c.components
}

func (c ChromaticOrb) GetDuration() int {
    return c.duration
}

func (c ChromaticOrb) GetConcentration() bool {
    return c.concentration
}

func (c ChromaticOrb) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Chromatic Orb is not implemented yet")
}
