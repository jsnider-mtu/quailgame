package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type WitchBolt struct {}

func (w WitchBolt) PrettyPrint() string {
    return "Witch Bolt"
}

func (w WitchBolt) GetLevel() int {
    return 1
}

func (a WitchBolt) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Witch Bolt is not implemented yet")
}
