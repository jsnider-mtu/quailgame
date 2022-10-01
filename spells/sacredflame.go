package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SacredFlame struct {}

func (s SacredFlame) PrettyPrint() string {
    return "Sacred Flame"
}

func (s SacredFlame) GetLevel() int {
    return 0
}

func (a SacredFlame) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Sacred Flame is not implemented yet")
}
