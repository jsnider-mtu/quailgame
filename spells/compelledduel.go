package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CompelledDuel struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (c CompelledDuel) PrettyPrint() string {
    return "Compelled Duel"
}

func (c CompelledDuel) GetLevel() int {
    return c.level
}

func (c CompelledDuel) GetCastTime() string {
    return c.casttime
}

func (c CompelledDuel) GetRange() int {
    return c.range
}

func (c CompelledDuel) GetComponents() string {
    return c.components
}

func (c CompelledDuel) GetDuration() int {
    return c.duration
}

func (c CompelledDuel) GetConcentration() bool {
    return c.concentration
}

func (c CompelledDuel) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Compelled Duel is not implemented yet")
}
