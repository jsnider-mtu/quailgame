package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FireBolt struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (f FireBolt) PrettyPrint() string {
    return "Fire Bolt"
}

func (f FireBolt) GetLevel() int {
    return f.level
}

func (f FireBolt) GetCastTime() string {
    return f.casttime
}

func (f FireBolt) GetRange() float64 {
    return float64(f.spellrange) * 4.8
}

func (f FireBolt) GetComponents() string {
    return f.components
}

func (f FireBolt) GetDuration() int {
    return f.duration
}

func (f FireBolt) GetConcentration() bool {
    return f.concentration
}

func (f FireBolt) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Fire Bolt is not implemented yet")
}
