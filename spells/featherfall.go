package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FeatherFall struct {}

func (f FeatherFall) PrettyPrint() string {
    return "Feather Fall"
}

func (f FeatherFall) GetLevel() int {
    return 0
}

func (a FeatherFall) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Feather Fall is not implemented yet")
}
