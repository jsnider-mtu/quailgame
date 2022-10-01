package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ShockingGrasp struct {}

func (s ShockingGrasp) PrettyPrint() string {
    return "Shocking Grasp"
}

func (a ShockingGrasp) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shocking Grasp is not implemented yet")
}
