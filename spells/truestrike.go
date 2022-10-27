package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type TrueStrike struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (t TrueStrike) PrettyPrint() string {
    return "True Strike"
}

func (t TrueStrike) GetLevel() int {
    return t.level
}

func (t TrueStrike) GetCastTime() string {
    return t.casttime
}

func (t TrueStrike) GetRange() float64 {
    return float64(t.spellrange) * 4.8
}

func (t TrueStrike) GetComponents() string {
    return t.components
}

func (t TrueStrike) GetDuration() int {
    return t.duration
}

func (t TrueStrike) GetConcentration() bool {
    return t.concentration
}

func (t TrueStrike) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell True Strike is not implemented yet")
}
