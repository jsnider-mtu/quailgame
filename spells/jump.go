package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Jump struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (j Jump) PrettyPrint() string {
    return "Jump"
}

func (j Jump) GetLevel() int {
    return j.level
}

func (j Jump) GetCastTime() string {
    return j.casttime
}

func (j Jump) GetRange() int {
    return j.spellrange
}

func (j Jump) GetComponents() string {
    return j.components
}

func (j Jump) GetDuration() int {
    return j.duration
}

func (j Jump) GetConcentration() bool {
    return j.concentration
}

func (j Jump) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Jump is not implemented yet")
}
