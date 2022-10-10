package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DisguiseSelf struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (d DisguiseSelf) PrettyPrint() string {
    return "Disguise Self"
}

func (d DisguiseSelf) GetLevel() int {
    return d.level
}

func (d DisguiseSelf) GetCastTime() string {
    return d.casttime
}

func (d DisguiseSelf) GetRange() int {
    return d.spellrange
}

func (d DisguiseSelf) GetComponents() string {
    return d.components
}

func (d DisguiseSelf) GetDuration() int {
    return d.duration
}

func (d DisguiseSelf) GetConcentration() bool {
    return d.concentration
}

func (d DisguiseSelf) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Disguise Self is not implemented yet")
}
