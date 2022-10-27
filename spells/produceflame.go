package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ProduceFlame struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (p ProduceFlame) PrettyPrint() string {
    return "Produce Flame"
}

func (p ProduceFlame) GetLevel() int {
    return p.level
}

func (p ProduceFlame) GetCastTime() string {
    return p.casttime
}

func (p ProduceFlame) GetRange() float64 {
    return float64(p.spellrange) * 4.8
}

func (p ProduceFlame) GetComponents() string {
    return p.components
}

func (p ProduceFlame) GetDuration() int {
    return p.duration
}

func (p ProduceFlame) GetConcentration() bool {
    return p.concentration
}

func (p ProduceFlame) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Produce Flame is not implemented yet")
}
