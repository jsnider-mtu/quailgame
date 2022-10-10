package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DancingLights struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (d DancingLights) PrettyPrint() string {
    return "Dancing Lights"
}

func (d DancingLights) GetLevel() int {
    return d.level
}

func (d DancingLights) GetCastTime() string {
    return d.casttime
}

func (d DancingLights) GetRange() int {
    return d.range
}

func (d DancingLights) GetComponents() string {
    return d.components
}

func (d DancingLights) GetDuration() int {
    return d.duration
}

func (d DancingLights) GetConcentration() bool {
    return d.concentration
}

func (d DancingLights) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Dancing Lights is not implemented yet")
}
