package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MageArmor struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (m MageArmor) PrettyPrint() string {
    return "Mage Armor"
}

func (m MageArmor) GetLevel() int {
    return m.level
}

func (m MageArmor) GetCastTime() string {
    return m.casttime
}

func (m MageArmor) GetRange() int {
    return m.range
}

func (m MageArmor) GetComponents() string {
    return m.components
}

func (m MageArmor) GetDuration() int {
    return m.duration
}

func (m MageArmor) GetConcentration() bool {
    return m.concentration
}

func (m MageArmor) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Mage Armor is not implemented yet")
}
