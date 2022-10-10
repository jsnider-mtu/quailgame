package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Shillelagh struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s Shillelagh) PrettyPrint() string {
    return "Shillelagh"
}

func (s Shillelagh) GetLevel() int {
    return s.level
}

func (s Shillelagh) GetCastTime() string {
    return s.casttime
}

func (s Shillelagh) GetRange() int {
    return s.spellrange
}

func (s Shillelagh) GetComponents() string {
    return s.components
}

func (s Shillelagh) GetDuration() int {
    return s.duration
}

func (s Shillelagh) GetConcentration() bool {
    return s.concentration
}

func (s Shillelagh) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shillelagh is not implemented yet")
}
