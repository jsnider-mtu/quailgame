package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Druidcraft struct {}

func (d Druidcraft) PrettyPrint() string {
    return "Druidcraft"
}

func (d Druidcraft) GetLevel() int {
    return 0
}

func (a Druidcraft) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Druidcraft is not implemented yet")
}
