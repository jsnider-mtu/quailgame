package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type WrathfulSmite struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (w WrathfulSmite) PrettyPrint() string {
    return "Wrathful Smite"
}

func (w WrathfulSmite) GetLevel() int {
    return w.level
}

func (w WrathfulSmite) GetCastTime() string {
    return w.casttime
}

func (w WrathfulSmite) GetRange() float64 {
    return float64(w.spellrange) * 4.8
}

func (w WrathfulSmite) GetComponents() string {
    return w.components
}

func (w WrathfulSmite) GetDuration() int {
    return w.duration
}

func (w WrathfulSmite) GetConcentration() bool {
    return w.concentration
}

func (w WrathfulSmite) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Wrathful Smite is not implemented yet")
}
