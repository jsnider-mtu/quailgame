package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Thunderwave struct {}

func (t Thunderwave) PrettyPrint() string {
    return "Thunderwave"
}

func (t Thunderwave) GetLevel() int {
    return 1
}

func (a Thunderwave) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Thunderwave is not implemented yet")
}
