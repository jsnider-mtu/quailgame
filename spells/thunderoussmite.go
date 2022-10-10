package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ThunderousSmite struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (t ThunderousSmite) PrettyPrint() string {
    return "Thunderous Smite"
}

func (t ThunderousSmite) GetLevel() int {
    return t.level
}

func (t ThunderousSmite) GetCastTime() string {
    return t.casttime
}

func (t ThunderousSmite) GetRange() int {
    return t.range
}

func (t ThunderousSmite) GetComponents() string {
    return t.components
}

func (t ThunderousSmite) GetDuration() int {
    return t.duration
}

func (t ThunderousSmite) GetConcentration() bool {
    return t.concentration
}

func (t ThunderousSmite) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thunderous Smite is not implemented yet")
}
