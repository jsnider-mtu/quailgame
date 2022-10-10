package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SearingSmite struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s SearingSmite) PrettyPrint() string {
    return "Searing Smite"
}

func (s SearingSmite) GetLevel() int {
    return s.level
}

func (s SearingSmite) GetCastTime() string {
    return s.casttime
}

func (s SearingSmite) GetRange() int {
    return s.spellrange
}

func (s SearingSmite) GetComponents() string {
    return s.components
}

func (s SearingSmite) GetDuration() int {
    return s.duration
}

func (s SearingSmite) GetConcentration() bool {
    return s.concentration
}

func (s SearingSmite) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Searing Smite is not implemented yet")
}
