package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type WrathfulSmite struct {}

func (w WrathfulSmite) PrettyPrint() string {
    return "Wrathful Smite"
}

func (w WrathfulSmite) GetLevel() int {
    return 1
}

func (a WrathfulSmite) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Wrathful Smite is not implemented yet")
}
