package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type RayOfFrost struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (r RayOfFrost) PrettyPrint() string {
    return "Ray of Frost"
}

func (r RayOfFrost) GetLevel() int {
    return r.level
}

func (r RayOfFrost) GetCastTime() string {
    return r.casttime
}

func (r RayOfFrost) GetRange() float64 {
    return float64(r.spellrange) * 4.8
}

func (r RayOfFrost) GetComponents() string {
    return r.components
}

func (r RayOfFrost) GetDuration() int {
    return r.duration
}

func (r RayOfFrost) GetConcentration() bool {
    return r.concentration
}

func (r RayOfFrost) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Ray of Frost is not implemented yet")
}
