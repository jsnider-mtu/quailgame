package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ShieldOfFaith struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s ShieldOfFaith) PrettyPrint() string {
    return "Shield of Faith"
}

func (s ShieldOfFaith) GetLevel() int {
    return s.level
}

func (s ShieldOfFaith) GetCastTime() string {
    return s.casttime
}

func (s ShieldOfFaith) GetRange() int {
    return s.spellrange
}

func (s ShieldOfFaith) GetComponents() string {
    return s.components
}

func (s ShieldOfFaith) GetDuration() int {
    return s.duration
}

func (s ShieldOfFaith) GetConcentration() bool {
    return s.concentration
}

func (s ShieldOfFaith) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shield of Faith is not implemented yet")
}
