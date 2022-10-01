package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Longstrider struct {}

func (l Longstrider) PrettyPrint() string {
    return "Longstrider"
}

func (a Longstrider) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Longstrider is not implemented yet")
}
