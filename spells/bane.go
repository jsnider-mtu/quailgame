package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Bane struct {}

func (b Bane) PrettyPrint() string {
    return "Bane"
}

func (b Bane) GetLevel() int {
    return 0
}

func (a Bane) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Bane is not implemented yet")
}
