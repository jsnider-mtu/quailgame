package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Alarm struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (a Alarm) PrettyPrint() string {
    return "Alarm"
}

func (a Alarm) GetLevel() int {
    return a.level
}

func (a Alarm) GetCastTime() string {
    return a.casttime
}

func (a Alarm) GetRange() int {
    return a.spellrange
}

func (a Alarm) GetComponents() string {
    return a.components
}

func (a Alarm) GetDuration() int {
    return a.duration
}

func (a Alarm) GetConcentration() bool {
    return a.concentration
}

func (a Alarm) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Alarm is not implemented yet")
}
