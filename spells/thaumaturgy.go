package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Thaumaturgy struct {}

func (t Thaumaturgy) PrettyPrint() string {
    return "Thaumaturgy"
}

func (a Thaumaturgy) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thaumaturgy is not implemented yet")
}
