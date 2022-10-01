package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type GuidingBolt struct {}

func (g GuidingBolt) PrettyPrint() string {
    return "Guiding Bolt"
}

func (g GuidingBolt) GetLevel() int {
    return 1
}

func (a GuidingBolt) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Guiding Bolt is not implemented yet")
}
