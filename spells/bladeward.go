package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type BladeWard struct {}

func (b BladeWard) PrettyPrint() string {
    return "Blade Ward"
}

func (b BladeWard) GetLevel() int {
    return 0
}

func (a BladeWard) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Blade Ward is not implemented yet")
}
