package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Heroism struct {}

func (h Heroism) PrettyPrint() string {
    return "Heroism"
}

func (h Heroism) GetLevel() int {
    return 0
}

func (a Heroism) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Heroism is not implemented yet")
}
