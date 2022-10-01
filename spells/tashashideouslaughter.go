package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type TashasHideousLaughter struct {}

func (t TashasHideousLaughter) PrettyPrint() string {
    return "Tasha's Hideous Laughter"
}

func (t TashasHideousLaughter) GetLevel() int {
    return 1
}

func (a TashasHideousLaughter) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Tasha's Hideous Laughter is not implemented yet")
}
