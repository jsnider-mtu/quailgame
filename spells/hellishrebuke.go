package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HellishRebuke struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (h HellishRebuke) PrettyPrint() string {
    return "Hellish Rebuke"
}

func (h HellishRebuke) GetLevel() int {
    return h.level
}

func (h HellishRebuke) GetCastTime() string {
    return h.casttime
}

func (h HellishRebuke) GetRange() int {
    return h.spellrange
}

func (h HellishRebuke) GetComponents() string {
    return h.components
}

func (h HellishRebuke) GetDuration() int {
    return h.duration
}

func (h HellishRebuke) GetConcentration() bool {
    return h.concentration
}

func (h HellishRebuke) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hellish Rebuke is not implemented yet")
}
