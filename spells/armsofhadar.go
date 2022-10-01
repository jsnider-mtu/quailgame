package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ArmsOfHadar struct {}

func (a ArmsOfHadar) PrettyPrint() string {
    return "Arms of Hadar"
}

func (a ArmsOfHadar) GetLevel() int {
    return 1
}

func (a ArmsOfHadar) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Arms of Hadar is not implemented yet")
}
