package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Shield struct {}

func (s Shield) PrettyPrint() string {
    return "Shield"
}

func (s Shield) GetLevel() int {
    return 1
}

func (a Shield) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shield is not implemented yet")
}
