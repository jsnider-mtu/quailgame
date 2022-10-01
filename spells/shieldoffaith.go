package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ShieldOfFaith struct {}

func (s ShieldOfFaith) PrettyPrint() string {
    return "Shield of Faith"
}

func (a ShieldOfFaith) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shield of Faith is not implemented yet")
}
