package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CharmPerson struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (c CharmPerson) PrettyPrint() string {
    return "Charm Person"
}

func (c CharmPerson) GetLevel() int {
    return c.level
}

func (c CharmPerson) GetCastTime() string {
    return c.casttime
}

func (c CharmPerson) GetRange() int {
    return c.range
}

func (c CharmPerson) GetComponents() string {
    return c.components
}

func (c CharmPerson) GetDuration() int {
    return c.duration
}

func (c CharmPerson) GetConcentration() bool {
    return c.concentration
}

func (c CharmPerson) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Charm Person is not implemented yet")
}
