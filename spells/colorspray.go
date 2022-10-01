package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ColorSpray struct {}

func (c ColorSpray) PrettyPrint() string {
    return "Color Spray"
}

func (a ColorSpray) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Color Spray is not implemented yet")
}
