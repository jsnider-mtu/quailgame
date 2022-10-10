package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ThornWhip struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (t ThornWhip) PrettyPrint() string {
    return "Thorn Whip"
}

func (t ThornWhip) GetLevel() int {
    return t.level
}

func (t ThornWhip) GetCastTime() string {
    return t.casttime
}

func (t ThornWhip) GetRange() int {
    return t.spellrange
}

func (t ThornWhip) GetComponents() string {
    return t.components
}

func (t ThornWhip) GetDuration() int {
    return t.duration
}

func (t ThornWhip) GetConcentration() bool {
    return t.concentration
}

func (t ThornWhip) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thorn Whip is not implemented yet")
}
