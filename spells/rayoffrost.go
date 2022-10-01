package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type RayOfFrost struct {}

func (r RayOfFrost) PrettyPrint() string {
    return "Ray of Frost"
}

func (r RayOfFrost) GetLevel() int {
    return 0
}

func (a RayOfFrost) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Ray of Frost is not implemented yet")
}
