package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ArmorOfAgathys struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (a ArmorOfAgathys) PrettyPrint() string {
    return "Armor of Agathys"
}

func (a ArmorOfAgathys) GetLevel() int {
    return a.level
}

func (a ArmorOfAgathys) GetCastTime() string {
    return a.casttime
}

func (a ArmorOfAgathys) GetRange() int {
    return a.range
}

func (a ArmorOfAgathys) GetComponents() string {
    return a.components
}

func (a ArmorOfAgathys) GetDuration() int {
    return a.duration
}

func (a ArmorOfAgathys) GetConcentration() bool {
    return a.concentration
}

func (a ArmorOfAgathys) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Armor of Agathys is not implemented yet")
}
