package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SilentImage struct {}

func (s SilentImage) PrettyPrint() string {
    return "Silent Image"
}

func (a SilentImage) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Silent Image is not implemented yet")
}
