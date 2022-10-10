package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ViciousMockery struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (v ViciousMockery) PrettyPrint() string {
    return "Vicious Mockery"
}

func (v ViciousMockery) GetLevel() int {
    return v.level
}

func (v ViciousMockery) GetCastTime() string {
    return v.casttime
}

func (v ViciousMockery) GetRange() int {
    return v.spellrange
}

func (v ViciousMockery) GetComponents() string {
    return v.components
}

func (v ViciousMockery) GetDuration() int {
    return v.duration
}

func (v ViciousMockery) GetConcentration() bool {
    return v.concentration
}

func (v ViciousMockery) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Vicious Mockery is not implemented yet")
}
