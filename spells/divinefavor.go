package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DivineFavor struct {}

func (d DivineFavor) PrettyPrint() string {
    return "Divine Favor"
}

func (d DivineFavor) GetLevel() int {
    return 1
}

func (a DivineFavor) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Divine Favor is not implemented yet")
}
