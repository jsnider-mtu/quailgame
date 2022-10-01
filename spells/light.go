package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Light struct {}

func (l Light) PrettyPrint() string {
    return "Light"
}

func (l Light) GetLevel() int {
    return 0
}

func (a Light) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Light is not implemented yet")
}
