package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type RayOfSickness struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (r RayOfSickness) PrettyPrint() string {
    return "Ray of Sickness"
}

func (r RayOfSickness) GetLevel() int {
    return r.level
}

func (r RayOfSickness) GetCastTime() string {
    return r.casttime
}

func (r RayOfSickness) GetRange() int {
    return r.range
}

func (r RayOfSickness) GetComponents() string {
    return r.components
}

func (r RayOfSickness) GetDuration() int {
    return r.duration
}

func (r RayOfSickness) GetConcentration() bool {
    return r.concentration
}

func (r RayOfSickness) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Ray of Sickness is not implemented yet")
}
