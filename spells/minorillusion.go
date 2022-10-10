package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MinorIllusion struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (m MinorIllusion) PrettyPrint() string {
    return "Minor Illusion"
}

func (m MinorIllusion) GetLevel() int {
    return m.level
}

func (m MinorIllusion) GetCastTime() string {
    return m.casttime
}

func (m MinorIllusion) GetRange() int {
    return m.spellrange
}

func (m MinorIllusion) GetComponents() string {
    return m.components
}

func (m MinorIllusion) GetDuration() int {
    return m.duration
}

func (m MinorIllusion) GetConcentration() bool {
    return m.concentration
}

func (m MinorIllusion) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Minor Illusion is not implemented yet")
}
