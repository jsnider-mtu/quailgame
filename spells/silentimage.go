package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SilentImage struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s SilentImage) PrettyPrint() string {
    return "Silent Image"
}

func (s SilentImage) GetLevel() int {
    return s.level
}

func (s SilentImage) GetCastTime() string {
    return s.casttime
}

func (s SilentImage) GetRange() int {
    return s.spellrange
}

func (s SilentImage) GetComponents() string {
    return s.components
}

func (s SilentImage) GetDuration() int {
    return s.duration
}

func (s SilentImage) GetConcentration() bool {
    return s.concentration
}

func (s SilentImage) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Silent Image is not implemented yet")
}
