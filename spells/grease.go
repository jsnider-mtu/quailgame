package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Grease struct {}

func (g Grease) PrettyPrint() string {
    return "Grease"
}

func (a Grease) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Grease is not implemented yet")
}
