package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DetectPoisonAndDisease struct {}

func (d DetectPoisonAndDisease) PrettyPrint() string {
    return "Detect Poison and Disease"
}

func (a DetectPoisonAndDisease) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Detect Poison and Disease is not implemented yet")
}
