package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type RayOfSickness struct {}

func (r RayOfSickness) PrettyPrint() string {
    return "Ray of Sickness"
}

func (a RayOfSickness) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Ray of Sickness is not implemented yet")
}
