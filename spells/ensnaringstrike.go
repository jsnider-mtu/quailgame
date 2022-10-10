package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type EnsnaringStrike struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (e EnsnaringStrike) PrettyPrint() string {
    return "Ensnaring Strike"
}

func (e EnsnaringStrike) GetLevel() int {
    return e.level
}

func (e EnsnaringStrike) GetCastTime() string {
    return e.casttime
}

func (e EnsnaringStrike) GetRange() int {
    return e.spellrange
}

func (e EnsnaringStrike) GetComponents() string {
    return e.components
}

func (e EnsnaringStrike) GetDuration() int {
    return e.duration
}

func (e EnsnaringStrike) GetConcentration() bool {
    return e.concentration
}

func (e EnsnaringStrike) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Ensnaring Strike is not implemented yet")
}
