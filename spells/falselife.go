package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FalseLife struct {}

func (f FalseLife) PrettyPrint() string {
    return "False Life"
}

func (f FalseLife) GetLevel() int {
    return 1
}

func (a FalseLife) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell False Life is not implemented yet")
}
