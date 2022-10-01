package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Resistance struct {}

func (r Resistance) PrettyPrint() string {
    return "Resistance"
}

func (a Resistance) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Resistance is not implemented yet")
}
