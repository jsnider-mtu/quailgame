package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Entangle struct {}

func (e Entangle) PrettyPrint() string {
    return "Entangle"
}

func (e Entangle) GetLevel() int {
    return 1
}

func (a Entangle) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Entangle is not implemented yet")
}
