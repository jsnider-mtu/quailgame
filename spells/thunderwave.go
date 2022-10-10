package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Thunderwave struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (t Thunderwave) PrettyPrint() string {
    return "Thunderwave"
}

func (t Thunderwave) GetLevel() int {
    return t.level
}

func (t Thunderwave) GetCastTime() string {
    return t.casttime
}

func (t Thunderwave) GetRange() int {
    return t.spellrange
}

func (t Thunderwave) GetComponents() string {
    return t.components
}

func (t Thunderwave) GetDuration() int {
    return t.duration
}

func (t Thunderwave) GetConcentration() bool {
    return t.concentration
}

func (t Thunderwave) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thunderwave is not implemented yet")
}
