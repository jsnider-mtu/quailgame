package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DivineFavor struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (d DivineFavor) PrettyPrint() string {
    return "Divine Favor"
}

func (d DivineFavor) GetLevel() int {
    return d.level
}

func (d DivineFavor) GetCastTime() string {
    return d.casttime
}

func (d DivineFavor) GetRange() int {
    return d.spellrange
}

func (d DivineFavor) GetComponents() string {
    return d.components
}

func (d DivineFavor) GetDuration() int {
    return d.duration
}

func (d DivineFavor) GetConcentration() bool {
    return d.concentration
}

func (d DivineFavor) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Divine Favor is not implemented yet")
}
