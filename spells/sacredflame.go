package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SacredFlame struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s SacredFlame) PrettyPrint() string {
    return "Sacred Flame"
}

func (s SacredFlame) GetLevel() int {
    return s.level
}

func (s SacredFlame) GetCastTime() string {
    return s.casttime
}

func (s SacredFlame) GetRange() int {
    return s.spellrange
}

func (s SacredFlame) GetComponents() string {
    return s.components
}

func (s SacredFlame) GetDuration() int {
    return s.duration
}

func (s SacredFlame) GetConcentration() bool {
    return s.concentration
}

func (s SacredFlame) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Sacred Flame is not implemented yet")
}
