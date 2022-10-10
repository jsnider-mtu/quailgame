package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FogCloud struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (f FogCloud) PrettyPrint() string {
    return "Fog Cloud"
}

func (f FogCloud) GetLevel() int {
    return f.level
}

func (f FogCloud) GetCastTime() string {
    return f.casttime
}

func (f FogCloud) GetRange() int {
    return f.spellrange
}

func (f FogCloud) GetComponents() string {
    return f.components
}

func (f FogCloud) GetDuration() int {
    return f.duration
}

func (f FogCloud) GetConcentration() bool {
    return f.concentration
}

func (f FogCloud) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Fog Cloud is not implemented yet")
}
