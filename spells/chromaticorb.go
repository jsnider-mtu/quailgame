package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ChromaticOrb struct {}

func (c ChromaticOrb) PrettyPrint() string {
    return "Chromatic Orb"
}

func (a ChromaticOrb) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Chromatic Orb is not implemented yet")
}
