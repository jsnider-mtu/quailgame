package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type InflictWounds struct {}

func (i InflictWounds) PrettyPrint() string {
    return "Inflict Wounds"
}

func (a InflictWounds) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Inflict Wounds is not implemented yet")
}
