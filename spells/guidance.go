package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Guidance struct {}

func (g Guidance) PrettyPrint() string {
    return "Guidance"
}

func (a Guidance) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Guidance is not implemented yet")
}
