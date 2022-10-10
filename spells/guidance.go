package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Guidance struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (g Guidance) PrettyPrint() string {
    return "Guidance"
}

func (g Guidance) GetLevel() int {
    return g.level
}

func (g Guidance) GetCastTime() string {
    return g.casttime
}

func (g Guidance) GetRange() int {
    return g.spellrange
}

func (g Guidance) GetComponents() string {
    return g.components
}

func (g Guidance) GetDuration() int {
    return g.duration
}

func (g Guidance) GetConcentration() bool {
    return g.concentration
}

func (g Guidance) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Guidance is not implemented yet")
}
