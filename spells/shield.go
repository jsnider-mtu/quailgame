package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Shield struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s Shield) PrettyPrint() string {
    return "Shield"
}

func (s Shield) GetLevel() int {
    return s.level
}

func (s Shield) GetCastTime() string {
    return s.casttime
}

func (s Shield) GetRange() int {
    return s.spellrange
}

func (s Shield) GetComponents() string {
    return s.components
}

func (s Shield) GetDuration() int {
    return s.duration
}

func (s Shield) GetConcentration() bool {
    return s.concentration
}

func (s Shield) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shield is not implemented yet")
}
