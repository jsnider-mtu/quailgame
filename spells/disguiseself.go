package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DisguiseSelf struct {}

func (d DisguiseSelf) PrettyPrint() string {
    return "Disguise Self"
}

func (d DisguiseSelf) GetLevel() int {
    return 1
}

func (a DisguiseSelf) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Disguise Self is not implemented yet")
}
