package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type BladeWard struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (b BladeWard) PrettyPrint() string {
    return "Blade Ward"
}

func (b BladeWard) GetLevel() int {
    return b.level
}

func (b BladeWard) GetCastTime() string {
    return b.casttime
}

func (b BladeWard) GetRange() int {
    return b.spellrange
}

func (b BladeWard) GetComponents() string {
    return b.components
}

func (b BladeWard) GetDuration() int {
    return b.duration
}

func (b BladeWard) GetConcentration() bool {
    return b.concentration
}

func (a BladeWard) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Blade Ward is not implemented yet")
}
