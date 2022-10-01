package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type FireBolt struct {}

func (f FireBolt) PrettyPrint() string {
    return "Fire Bolt"
}

func (a FireBolt) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Fire Bolt is not implemented yet")
}
