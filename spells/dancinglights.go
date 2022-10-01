package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DancingLights struct {}

func (d DancingLights) PrettyPrint() string {
    return "Dancing Lights"
}

func (d DancingLights) GetLevel() int {
    return 0
}

func (a DancingLights) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Dancing Lights is not implemented yet")
}
