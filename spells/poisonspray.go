package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type PoisonSpray struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (p PoisonSpray) PrettyPrint() string {
    return "Poison Spray"
}

func (p PoisonSpray) GetLevel() int {
    return p.level
}

func (p PoisonSpray) GetCastTime() string {
    return p.casttime
}

func (p PoisonSpray) GetRange() int {
    return p.spellrange
}

func (p PoisonSpray) GetComponents() string {
    return p.components
}

func (p PoisonSpray) GetDuration() int {
    return p.duration
}

func (p PoisonSpray) GetConcentration() bool {
    return p.concentration
}

func (p PoisonSpray) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Poison Spray is not implemented yet")
}
