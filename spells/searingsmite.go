package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SearingSmite struct {}

func (s SearingSmite) PrettyPrint() string {
    return "Searing Smite"
}

func (s SearingSmite) GetLevel() int {
    return 1
}

func (a SearingSmite) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Searing Smite is not implemented yet")
}
