package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ArmorOfAgathys struct {}

func (a ArmorOfAgathys) PrettyPrint() string {
    return "Armor of Agathys"
}

func (a ArmorOfAgathys) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Armor of Agathys is not implemented yet")
}
