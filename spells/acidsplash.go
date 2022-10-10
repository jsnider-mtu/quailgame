package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type AcidSplash struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (a AcidSplash) PrettyPrint() string {
    return "Acid Splash"
}

func (a AcidSplash) GetLevel() int {
    return a.level
}

func (a AcidSplash) GetCastTime() string {
    return a.casttime
}

func (a AcidSplash) GetRange() int {
    return a.spellrange
}

func (a AcidSplash) GetComponents() string {
    return a.components
}

func (a AcidSplash) GetDuration() int {
    return a.duration
}

func (a AcidSplash) GetConcentration() bool {
    return a.concentration
}

func (a AcidSplash) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Acid Splash is not implemented yet")
}
