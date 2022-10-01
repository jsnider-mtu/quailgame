package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Friends struct {}

func (f Friends) PrettyPrint() string {
    return "Friends"
}

func (f Friends) GetLevel() int {
    return 0
}

func (a Friends) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Friends is not implemented yet")
}
