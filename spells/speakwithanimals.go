package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SpeakWithAnimals struct {}

func (s SpeakWithAnimals) PrettyPrint() string {
    return "Speak with Animals"
}

func (s SpeakWithAnimals) GetLevel() int {
    return 1
}

func (a SpeakWithAnimals) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Speak with Animals is not implemented yet")
}
