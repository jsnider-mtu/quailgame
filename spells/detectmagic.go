package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DetectMagic struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (d DetectMagic) PrettyPrint() string {
    return "Detect Magic"
}

func (d DetectMagic) GetLevel() int {
    return d.level
}

func (d DetectMagic) GetCastTime() string {
    return d.casttime
}

func (d DetectMagic) GetRange() int {
    return d.spellrange
}

func (d DetectMagic) GetComponents() string {
    return d.components
}

func (d DetectMagic) GetDuration() int {
    return d.duration
}

func (d DetectMagic) GetConcentration() bool {
    return d.concentration
}

func (d DetectMagic) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Detect Magic is not implemented yet")
}
