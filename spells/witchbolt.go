package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type WitchBolt struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (w WitchBolt) PrettyPrint() string {
    return "Witch Bolt"
}

func (w WitchBolt) GetLevel() int {
    return w.level
}

func (w WitchBolt) GetCastTime() string {
    return w.casttime
}

func (w WitchBolt) GetRange() float64 {
    return float64(w.spellrange) * 4.8
}

func (w WitchBolt) GetComponents() string {
    return w.components
}

func (w WitchBolt) GetDuration() int {
    return w.duration
}

func (w WitchBolt) GetConcentration() bool {
    return w.concentration
}

func (w WitchBolt) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Witch Bolt is not implemented yet")
}
