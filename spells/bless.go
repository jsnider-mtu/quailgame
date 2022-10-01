package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Bless struct {}

func (b Bless) PrettyPrint() string {
    return "Bless"
}

func (b Bless) GetLevel() int {
    return 0
}

func (a Bless) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Bless is not implemented yet")
}
