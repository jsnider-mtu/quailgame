package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Goodberry struct {}

func (g Goodberry) PrettyPrint() string {
    return "Goodberry"
}

func (g Goodberry) GetLevel() int {
    return 0
}

func (a Goodberry) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Goodberry is not implemented yet")
}
