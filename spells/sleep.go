package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Sleep struct {}

func (s Sleep) PrettyPrint() string {
    return "Sleep"
}

func (s Sleep) GetLevel() int {
    return 0
}

func (a Sleep) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Sleep is not implemented yet")
}
