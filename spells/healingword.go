package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HealingWord struct {}

func (h HealingWord) PrettyPrint() string {
    return "Healing Word"
}

func (h HealingWord) GetLevel() int {
    return 0
}

func (a HealingWord) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Healing Word is not implemented yet")
}
