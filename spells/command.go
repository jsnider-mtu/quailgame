package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Command struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (c Command) PrettyPrint() string {
    return "Command"
}

func (c Command) GetLevel() int {
    return c.level
}

func (c Command) GetCastTime() string {
    return c.casttime
}

func (c Command) GetRange() int {
    return c.range
}

func (c Command) GetComponents() string {
    return c.components
}

func (c Command) GetDuration() int {
    return c.duration
}

func (c Command) GetConcentration() bool {
    return c.concentration
}

func (c Command) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Command is not implemented yet")
}
