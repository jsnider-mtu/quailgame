package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ChillTouch struct {}

func (c ChillTouch) PrettyPrint() string {
    return "Chill Touch"
}

func (a ChillTouch) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Chill Touch is not implemented yet")
}
