package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FogCloud struct {}

func (f FogCloud) PrettyPrint() string {
    return "Fog Cloud"
}

func (f FogCloud) GetLevel() int {
    return 0
}

func (a FogCloud) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Fog Cloud is not implemented yet")
}
