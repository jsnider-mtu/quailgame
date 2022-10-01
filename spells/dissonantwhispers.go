package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DissonantWhispers struct {}

func (d DissonantWhispers) PrettyPrint() string {
    return "Dissonant Whispers"
}

func (a DissonantWhispers) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Dissonant Whispers is not implemented yet")
}
