package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type GuidingBolt struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (g GuidingBolt) PrettyPrint() string {
    return "Guiding Bolt"
}

func (g GuidingBolt) GetLevel() int {
    return g.level
}

func (g GuidingBolt) GetCastTime() string {
    return g.casttime
}

func (g GuidingBolt) GetRange() float64 {
    return float64(g.spellrange) * 4.8
}

func (g GuidingBolt) GetComponents() string {
    return g.components
}

func (g GuidingBolt) GetDuration() int {
    return g.duration
}

func (g GuidingBolt) GetConcentration() bool {
    return g.concentration
}

func (g GuidingBolt) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Guiding Bolt is not implemented yet")
}
