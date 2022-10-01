package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Hex struct {}

func (h Hex) PrettyPrint() string {
    return "Hex"
}

func (h Hex) GetLevel() int {
    return 0
}

func (a Hex) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hex is not implemented yet")
}
