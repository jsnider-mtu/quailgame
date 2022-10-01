package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type TrueStrike struct {}

func (t TrueStrike) PrettyPrint() string {
    return "True Strike"
}

func (t TrueStrike) GetLevel() int {
    return 0
}

func (a TrueStrike) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell True Strike is not implemented yet")
}
