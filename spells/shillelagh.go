package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type Shillelagh struct {}

func (s Shillelagh) PrettyPrint() string {
    return "Shillelagh"
}

func (s Shillelagh) GetLevel() int {
    return 0
}

func (a Shillelagh) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Shillelagh is not implemented yet")
}
