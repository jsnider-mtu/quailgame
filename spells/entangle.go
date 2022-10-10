package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Entangle struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (e Entangle) PrettyPrint() string {
    return "Entangle"
}

func (e Entangle) GetLevel() int {
    return e.level
}

func (e Entangle) GetCastTime() string {
    return e.casttime
}

func (e Entangle) GetRange() int {
    return e.spellrange
}

func (e Entangle) GetComponents() string {
    return e.components
}

func (e Entangle) GetDuration() int {
    return e.duration
}

func (e Entangle) GetConcentration() bool {
    return e.concentration
}

func (e Entangle) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Entangle is not implemented yet")
}
