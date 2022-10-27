package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HailOfThorns struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (h HailOfThorns) PrettyPrint() string {
    return "Hail of Thorns"
}

func (h HailOfThorns) GetLevel() int {
    return h.level
}

func (h HailOfThorns) GetCastTime() string {
    return h.casttime
}

func (h HailOfThorns) GetRange() float64 {
    return float64(h.spellrange) * 4.8
}

func (h HailOfThorns) GetComponents() string {
    return h.components
}

func (h HailOfThorns) GetDuration() int {
    return h.duration
}

func (h HailOfThorns) GetConcentration() bool {
    return h.concentration
}

func (h HailOfThorns) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hail of Thorns is not implemented yet")
}
