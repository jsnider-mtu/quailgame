package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ComprehendLanguages struct {}

func (c ComprehendLanguages) PrettyPrint() string {
    return "Comprehend Languages"
}

func (c ComprehendLanguages) GetLevel() int {
    return 1
}

func (a ComprehendLanguages) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Comprehend Languages is not implemented yet")
}
