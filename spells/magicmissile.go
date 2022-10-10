package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MagicMissile struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (m MagicMissile) PrettyPrint() string {
    return "Magic Missile"
}

func (m MagicMissile) GetLevel() int {
    return m.level
}

func (m MagicMissile) GetCastTime() string {
    return m.casttime
}

func (m MagicMissile) GetRange() int {
    return m.spellrange
}

func (m MagicMissile) GetComponents() string {
    return m.components
}

func (m MagicMissile) GetDuration() int {
    return m.duration
}

func (m MagicMissile) GetConcentration() bool {
    return m.concentration
}

func (m MagicMissile) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Magic Missile is not implemented yet")
}
