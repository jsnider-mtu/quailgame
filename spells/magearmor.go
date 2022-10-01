package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MageArmor struct {}

func (m MageArmor) PrettyPrint() string {
    return "Mage Armor"
}

func (m MageArmor) GetLevel() int {
    return 0
}

func (a MageArmor) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Mage Armor is not implemented yet")
}
