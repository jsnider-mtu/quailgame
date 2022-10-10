package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type TashasHideousLaughter struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (t TashasHideousLaughter) PrettyPrint() string {
    return "Tasha's Hideous Laughter"
}

func (t TashasHideousLaughter) GetLevel() int {
    return t.level
}

func (t TashasHideousLaughter) GetCastTime() string {
    return t.casttime
}

func (t TashasHideousLaughter) GetRange() int {
    return t.spellrange
}

func (t TashasHideousLaughter) GetComponents() string {
    return t.components
}

func (t TashasHideousLaughter) GetDuration() int {
    return t.duration
}

func (t TashasHideousLaughter) GetConcentration() bool {
    return t.concentration
}

func (t TashasHideousLaughter) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Tasha's Hideous Laughter is not implemented yet")
}
