package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Bless struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (b Bless) PrettyPrint() string {
    return "Bless"
}

func (b Bless) GetLevel() int {
    return b.level
}

func (b Bless) GetCastTime() string {
    return b.casttime
}

func (b Bless) GetRange() int {
    return b.spellrange
}

func (b Bless) GetComponents() string {
    return b.components
}

func (b Bless) GetDuration() int {
    return b.duration
}

func (b Bless) GetConcentration() bool {
    return b.concentration
}

func (b Bless) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Bless is not implemented yet")
}
