package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type BurningHands struct {}

func (b BurningHands) PrettyPrint() string {
    return "Burning Hands"
}

func (b BurningHands) GetLevel() int {
    return 1
}

func (a BurningHands) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Burning Hands is not implemented yet")
}
