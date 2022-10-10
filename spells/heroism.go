package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Heroism struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (h Heroism) PrettyPrint() string {
    return "Heroism"
}

func (h Heroism) GetLevel() int {
    return h.level
}

func (h Heroism) GetCastTime() string {
    return h.casttime
}

func (h Heroism) GetRange() int {
    return h.range
}

func (h Heroism) GetComponents() string {
    return h.components
}

func (h Heroism) GetDuration() int {
    return h.duration
}

func (h Heroism) GetConcentration() bool {
    return h.concentration
}

func (h Heroism) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Heroism is not implemented yet")
}
