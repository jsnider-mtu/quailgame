package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MinorIllusion struct {}

func (m MinorIllusion) PrettyPrint() string {
    return "Minor Illusion"
}

func (a MinorIllusion) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Minor Illusion is not implemented yet")
}
