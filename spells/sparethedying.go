package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SpareTheDying struct {}

func (s SpareTheDying) PrettyPrint() string {
    return "Spare the Dying"
}

func (s SpareTheDying) GetLevel() int {
    return 0
}

func (a SpareTheDying) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Spare the Dying is not implemented yet")
}
