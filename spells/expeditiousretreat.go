package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ExpeditiousRetreat struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (e ExpeditiousRetreat) PrettyPrint() string {
    return "Expeditious Retreat"
}

func (e ExpeditiousRetreat) GetLevel() int {
    return e.level
}

func (e ExpeditiousRetreat) GetCastTime() string {
    return e.casttime
}

func (e ExpeditiousRetreat) GetRange() float64 {
    return float64(e.spellrange) * 4.8
}

func (e ExpeditiousRetreat) GetComponents() string {
    return e.components
}

func (e ExpeditiousRetreat) GetDuration() int {
    return e.duration
}

func (e ExpeditiousRetreat) GetConcentration() bool {
    return e.concentration
}

func (e ExpeditiousRetreat) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Expeditious Retreat is not implemented yet")
}
