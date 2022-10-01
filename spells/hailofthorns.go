package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HailOfThorns struct {}

func (h HailOfThorns) PrettyPrint() string {
    return "Hail of Thorns"
}

func (a HailOfThorns) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hail of Thorns is not implemented yet")
}
