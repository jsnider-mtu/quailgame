package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DetectMagic struct {}

func (d DetectMagic) PrettyPrint() string {
    return "Detect Magic"
}

func (d DetectMagic) GetLevel() int {
    return 0
}

func (a DetectMagic) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Detect Magic is not implemented yet")
}
