package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FindFamiliar struct {}

func (f FindFamiliar) PrettyPrint() string {
    return "Find Familiar"
}

func (f FindFamiliar) GetLevel() int {
    return 1
}

func (a FindFamiliar) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Find Familiar is not implemented yet")
}
