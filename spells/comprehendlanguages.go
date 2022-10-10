package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ComprehendLanguages struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (c ComprehendLanguages) PrettyPrint() string {
    return "Comprehend Languages"
}

func (c ComprehendLanguages) GetLevel() int {
    return c.level
}

func (c ComprehendLanguages) GetCastTime() string {
    return c.casttime
}

func (c ComprehendLanguages) GetRange() int {
    return c.range
}

func (c ComprehendLanguages) GetComponents() string {
    return c.components
}

func (c ComprehendLanguages) GetDuration() int {
    return c.duration
}

func (c ComprehendLanguages) GetConcentration() bool {
    return c.concentration
}

func (c ComprehendLanguages) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Comprehend Languages is not implemented yet")
}
