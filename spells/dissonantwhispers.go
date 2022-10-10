package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DissonantWhispers struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (d DissonantWhispers) PrettyPrint() string {
    return "Dissonant Whispers"
}

func (d DissonantWhispers) GetLevel() int {
    return d.level
}

func (d DissonantWhispers) GetCastTime() string {
    return d.casttime
}

func (d DissonantWhispers) GetRange() int {
    return d.range
}

func (d DissonantWhispers) GetComponents() string {
    return d.components
}

func (d DissonantWhispers) GetDuration() int {
    return d.duration
}

func (d DissonantWhispers) GetConcentration() bool {
    return d.concentration
}

func (d DissonantWhispers) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Dissonant Whispers is not implemented yet")
}
