package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MageHand struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (m MageHand) PrettyPrint() string {
    return "Mage Hand"
}

func (m MageHand) GetLevel() int {
    return m.level
}

func (m MageHand) GetCastTime() string {
    return m.casttime
}

func (m MageHand) GetRange() int {
    return m.range
}

func (m MageHand) GetComponents() string {
    return m.components
}

func (m MageHand) GetDuration() int {
    return m.duration
}

func (m MageHand) GetConcentration() bool {
    return m.concentration
}

func (m MageHand) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Mage Hand is not implemented yet")
}
