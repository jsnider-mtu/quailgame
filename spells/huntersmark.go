package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type HuntersMark struct {}

func (h HuntersMark) PrettyPrint() string {
    return "Hunter's Mark"
}

func (h HuntersMark) GetLevel() int {
    return 0
}

func (a HuntersMark) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Hunter's Mark is not implemented yet")
}
