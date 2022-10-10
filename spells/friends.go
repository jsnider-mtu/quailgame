package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Friends struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (f Friends) PrettyPrint() string {
    return "Friends"
}

func (f Friends) GetLevel() int {
    return f.level
}

func (f Friends) GetCastTime() string {
    return f.casttime
}

func (f Friends) GetRange() int {
    return f.spellrange
}

func (f Friends) GetComponents() string {
    return f.components
}

func (f Friends) GetDuration() int {
    return f.duration
}

func (f Friends) GetConcentration() bool {
    return f.concentration
}

func (f Friends) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Friends is not implemented yet")
}
