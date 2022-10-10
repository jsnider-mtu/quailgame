package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type UnseenServant struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (u UnseenServant) PrettyPrint() string {
    return "Unseen Servant"
}

func (u UnseenServant) GetLevel() int {
    return u.level
}

func (u UnseenServant) GetCastTime() string {
    return u.casttime
}

func (u UnseenServant) GetRange() int {
    return u.range
}

func (u UnseenServant) GetComponents() string {
    return u.components
}

func (u UnseenServant) GetDuration() int {
    return u.duration
}

func (u UnseenServant) GetConcentration() bool {
    return u.concentration
}

func (u UnseenServant) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Unseen Servant is not implemented yet")
}
