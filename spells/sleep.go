package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Sleep struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s Sleep) PrettyPrint() string {
    return "Sleep"
}

func (s Sleep) GetLevel() int {
    return s.level
}

func (s Sleep) GetCastTime() string {
    return s.casttime
}

func (s Sleep) GetRange() int {
    return s.spellrange
}

func (s Sleep) GetComponents() string {
    return s.components
}

func (s Sleep) GetDuration() int {
    return s.duration
}

func (s Sleep) GetConcentration() bool {
    return s.concentration
}

func (s Sleep) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Sleep is not implemented yet")
}
