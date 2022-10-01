package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HellishRebuke struct {}

func (h HellishRebuke) PrettyPrint() string {
    return "Hellish Rebuke"
}

func (h HellishRebuke) GetLevel() int {
    return 0
}

func (a HellishRebuke) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hellish Rebuke is not implemented yet")
}
