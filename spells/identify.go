package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Identify struct {}

func (i Identify) PrettyPrint() string {
    return "Identify"
}

func (a Identify) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Identify is not implemented yet")
}
