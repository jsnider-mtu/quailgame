package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DetectEvilAndGood struct {}

func (d DetectEvilAndGood) PrettyPrint() string {
    return "Detect Evil and Good"
}

func (d DetectEvilAndGood) GetLevel() int {
    return 1
}

func (a DetectEvilAndGood) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Detect Evil and Good is not implemented yet")
}
