package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MageHand struct {}

func (m MageHand) PrettyPrint() string {
    return "Mage Hand"
}

func (a MageHand) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Mage Hand is not implemented yet")
}
