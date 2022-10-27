package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type SpeakWithAnimals struct {
    level int
    casttime string
    spellrange int
    components string
    duration int
    concentration bool
}

func (s SpeakWithAnimals) PrettyPrint() string {
    return "Speak with Animals"
}

func (s SpeakWithAnimals) GetLevel() int {
    return s.level
}

func (s SpeakWithAnimals) GetCastTime() string {
    return s.casttime
}

func (s SpeakWithAnimals) GetRange() float64 {
    return float64(s.spellrange) * 4.8
}

func (s SpeakWithAnimals) GetComponents() string {
    return s.components
}

func (s SpeakWithAnimals) GetDuration() int {
    return s.duration
}

func (s SpeakWithAnimals) GetConcentration() bool {
    return s.concentration
}

func (s SpeakWithAnimals) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Speak with Animals is not implemented yet")
}
