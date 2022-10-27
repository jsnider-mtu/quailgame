package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HealingWord struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (h HealingWord) PrettyPrint() string {
    return "Healing Word"
}

func (h HealingWord) GetLevel() int {
    return h.level
}

func (h HealingWord) GetCastTime() string {
    return h.casttime
}

func (h HealingWord) GetRange() float64 {
    return float64(h.spellrange) * 4.8
}

func (h HealingWord) GetComponents() string {
    return h.components
}

func (h HealingWord) GetDuration() int {
    return h.duration
}

func (h HealingWord) GetConcentration() bool {
    return h.concentration
}

func (h HealingWord) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Healing Word is not implemented yet")
}
