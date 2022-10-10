package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DetectEvilAndGood struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (d DetectEvilAndGood) PrettyPrint() string {
    return "Detect Evil and Good"
}

func (d DetectEvilAndGood) GetLevel() int {
    return d.level
}

func (d DetectEvilAndGood) GetCastTime() string {
    return d.casttime
}

func (d DetectEvilAndGood) GetRange() int {
    return d.spellrange
}

func (d DetectEvilAndGood) GetComponents() string {
    return d.components
}

func (d DetectEvilAndGood) GetDuration() int {
    return d.duration
}

func (d DetectEvilAndGood) GetConcentration() bool {
    return d.concentration
}

func (d DetectEvilAndGood) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Detect Evil and Good is not implemented yet")
}
