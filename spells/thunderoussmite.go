package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ThunderousSmite struct {}

func (t ThunderousSmite) PrettyPrint() string {
    return "Thunderous Smite"
}

func (a ThunderousSmite) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thunderous Smite is not implemented yet")
}
