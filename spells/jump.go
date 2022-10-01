package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Jump struct {}

func (j Jump) PrettyPrint() string {
    return "Jump"
}

func (j Jump) GetLevel() int {
    return 0
}

func (a Jump) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Jump is not implemented yet")
}
