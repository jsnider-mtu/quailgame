package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Sanctuary struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (s Sanctuary) PrettyPrint() string {
    return "Sanctuary"
}

func (s Sanctuary) GetLevel() int {
    return s.level
}

func (s Sanctuary) GetCastTime() string {
    return s.casttime
}

func (s Sanctuary) GetRange() int {
    return s.range
}

func (s Sanctuary) GetComponents() string {
    return s.components
}

func (s Sanctuary) GetDuration() int {
    return s.duration
}

func (s Sanctuary) GetConcentration() bool {
    return s.concentration
}

func (s Sanctuary) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Sanctuary is not implemented yet")
}
