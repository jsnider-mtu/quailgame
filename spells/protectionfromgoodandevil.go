package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ProtectionFromGoodAndEvil struct {}

func (p ProtectionFromGoodAndEvil) PrettyPrint() string {
    return "Protection from Good and Evil"
}

func (a ProtectionFromGoodAndEvil) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Protection from Good and Evil is not implemented yet")
}
