package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ThornWhip struct {}

func (t ThornWhip) PrettyPrint() string {
    return "Thorn Whip"
}

func (t ThornWhip) GetLevel() int {
    return 0
}

func (a ThornWhip) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thorn Whip is not implemented yet")
}
