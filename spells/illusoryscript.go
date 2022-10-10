package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type IllusoryScript struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (i IllusoryScript) PrettyPrint() string {
    return "Illusory Script"
}

func (i IllusoryScript) GetLevel() int {
    return i.level
}

func (i IllusoryScript) GetCastTime() string {
    return i.casttime
}

func (i IllusoryScript) GetRange() int {
    return i.range
}

func (i IllusoryScript) GetComponents() string {
    return i.components
}

func (i IllusoryScript) GetDuration() int {
    return i.duration
}

func (i IllusoryScript) GetConcentration() bool {
    return i.concentration
}

func (i IllusoryScript) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Illusory Script is not implemented yet")
}
