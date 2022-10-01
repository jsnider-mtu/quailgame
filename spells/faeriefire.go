package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FaerieFire struct {}

func (f FaerieFire) PrettyPrint() string {
    return "Faerie Fire"
}

func (a FaerieFire) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Faerie Fire is not implemented yet")
}
