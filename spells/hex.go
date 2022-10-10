package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Hex struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (h Hex) PrettyPrint() string {
    return "Hex"
}

func (h Hex) GetLevel() int {
    return h.level
}

func (h Hex) GetCastTime() string {
    return h.casttime
}

func (h Hex) GetRange() int {
    return h.range
}

func (h Hex) GetComponents() string {
    return h.components
}

func (h Hex) GetDuration() int {
    return h.duration
}

func (h Hex) GetConcentration() bool {
    return h.concentration
}

func (h Hex) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hex is not implemented yet")
}
