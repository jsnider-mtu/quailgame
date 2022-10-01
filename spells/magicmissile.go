package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type MagicMissile struct {}

func (m MagicMissile) PrettyPrint() string {
    return "Magic Missile"
}

func (m MagicMissile) GetLevel() int {
    return 0
}

func (a MagicMissile) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Magic Missile is not implemented yet")
}
