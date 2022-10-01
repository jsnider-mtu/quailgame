package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type AcidSplash struct {}

func (a AcidSplash) PrettyPrint() string {
    return "Acid Splash"
}

func (a AcidSplash) GetLevel() int {
    return 0
}

func (a AcidSplash) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Acid Splash is not implemented yet")
}
