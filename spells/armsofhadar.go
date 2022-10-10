package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ArmsOfHadar struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (a ArmsOfHadar) PrettyPrint() string {
    return "Arms of Hadar"
}

func (a ArmsOfHadar) GetLevel() int {
    return a.level
}

func (a ArmsOfHadar) GetCastTime() string {
    return a.casttime
}

func (a ArmsOfHadar) GetRange() int {
    return a.spellrange
}

func (a ArmsOfHadar) GetComponents() string {
    return a.components
}

func (a ArmsOfHadar) GetDuration() int {
    return a.duration
}

func (a ArmsOfHadar) GetConcentration() bool {
    return a.concentration
}

func (a ArmsOfHadar) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Arms of Hadar is not implemented yet")
}
