package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ShockingGrasp struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (s ShockingGrasp) PrettyPrint() string {
    return "Shocking Grasp"
}

func (s ShockingGrasp) GetLevel() int {
    return s.level
}

func (s ShockingGrasp) GetCastTime() string {
    return s.casttime
}

func (s ShockingGrasp) GetRange() int {
    return s.range
}

func (s ShockingGrasp) GetComponents() string {
    return s.components
}

func (s ShockingGrasp) GetDuration() int {
    return s.duration
}

func (s ShockingGrasp) GetConcentration() bool {
    return s.concentration
}

func (s ShockingGrasp) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shocking Grasp is not implemented yet")
}
