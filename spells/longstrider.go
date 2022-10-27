package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Longstrider struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (l Longstrider) PrettyPrint() string {
    return "Longstrider"
}

func (l Longstrider) GetLevel() int {
    return l.level
}

func (l Longstrider) GetCastTime() string {
    return l.casttime
}

func (l Longstrider) GetRange() float64 {
    return float64(l.spellrange) * 4.8
}

func (l Longstrider) GetComponents() string {
    return l.components
}

func (l Longstrider) GetDuration() int {
    return l.duration
}

func (l Longstrider) GetConcentration() bool {
    return l.concentration
}

func (l Longstrider) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Longstrider is not implemented yet")
}
