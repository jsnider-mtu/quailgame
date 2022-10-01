package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type IllusoryScript struct {}

func (i IllusoryScript) PrettyPrint() string {
    return "Illusory Script"
}

func (i IllusoryScript) GetLevel() int {
    return 1
}

func (a IllusoryScript) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Illusory Script is not implemented yet")
}
