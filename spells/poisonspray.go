package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type PoisonSpray struct {}

func (p PoisonSpray) PrettyPrint() string {
    return "Poison Spray"
}

func (p PoisonSpray) GetLevel() int {
    return 0
}

func (a PoisonSpray) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Poison Spray is not implemented yet")
}
