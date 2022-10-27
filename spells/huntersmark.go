package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HuntersMark struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (h HuntersMark) PrettyPrint() string {
    return "Hunter's Mark"
}

func (h HuntersMark) GetLevel() int {
    return h.level
}

func (h HuntersMark) GetCastTime() string {
    return h.casttime
}

func (h HuntersMark) GetRange() float64 {
    return float64(h.spellrange) * 4.8
}

func (h HuntersMark) GetComponents() string {
    return h.components
}

func (h HuntersMark) GetDuration() int {
    return h.duration
}

func (h HuntersMark) GetConcentration() bool {
    return h.concentration
}

func (h HuntersMark) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hunter's Mark is not implemented yet")
}
