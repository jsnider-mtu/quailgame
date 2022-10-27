package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ColorSpray struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (c ColorSpray) PrettyPrint() string {
    return "Color Spray"
}

func (c ColorSpray) GetLevel() int {
    return c.level
}

func (c ColorSpray) GetCastTime() string {
    return c.casttime
}

func (c ColorSpray) GetRange() float64 {
    return float64(c.spellrange) * 4.8
}

func (c ColorSpray) GetComponents() string {
    return c.components
}

func (c ColorSpray) GetDuration() int {
    return c.duration
}

func (c ColorSpray) GetConcentration() bool {
    return c.concentration
}

func (c ColorSpray) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Color Spray is not implemented yet")
}
