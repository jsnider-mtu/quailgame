package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Sanctuary struct {}

func (s Sanctuary) PrettyPrint() string {
    return "Sanctuary"
}

func (a Sanctuary) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Sanctuary is not implemented yet")
}
