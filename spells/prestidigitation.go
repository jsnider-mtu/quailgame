package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Prestidigitation struct {}

func (p Prestidigitation) PrettyPrint() string {
    return "Prestidigitation"
}

func (a Prestidigitation) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Prestidigitation is not implemented yet")
}
