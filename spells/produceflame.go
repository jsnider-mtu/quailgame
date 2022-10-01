package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ProduceFlame struct {}

func (p ProduceFlame) PrettyPrint() string {
    return "Produce Flame"
}

func (p ProduceFlame) GetLevel() int {
    return 0
}

func (a ProduceFlame) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Produce Flame is not implemented yet")
}
