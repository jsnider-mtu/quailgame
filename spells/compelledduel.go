package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CompelledDuel struct {}

func (c CompelledDuel) PrettyPrint() string {
    return "Compelled Duel"
}

func (a CompelledDuel) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Compelled Duel is not implemented yet")
}
