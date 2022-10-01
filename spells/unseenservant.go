package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type UnseenServant struct {}

func (u UnseenServant) PrettyPrint() string {
    return "Unseen Servant"
}

func (a UnseenServant) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Unseen Servant is not implemented yet")
}
