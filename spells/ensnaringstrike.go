package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type EnsnaringStrike struct {}

func (e EnsnaringStrike) PrettyPrint() string {
    return "Ensnaring Strike"
}

func (e EnsnaringStrike) GetLevel() int {
    return 1
}

func (a EnsnaringStrike) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Ensnaring Strike is not implemented yet")
}
