package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Mending struct {}

func (m Mending) PrettyPrint() string {
    return "Mending"
}

func (a Mending) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Mending is not implemented yet")
}
