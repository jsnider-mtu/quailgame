package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Command struct {}

func (c Command) PrettyPrint() string {
    return "Command"
}

func (c Command) GetLevel() int {
    return 0
}

func (a Command) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Command is not implemented yet")
}
