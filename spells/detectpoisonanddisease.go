package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type DetectPoisonAndDisease struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (d DetectPoisonAndDisease) PrettyPrint() string {
    return "Detect Poison and Disease"
}

func (d DetectPoisonAndDisease) GetLevel() int {
    return d.level
}

func (d DetectPoisonAndDisease) GetCastTime() string {
    return d.casttime
}

func (d DetectPoisonAndDisease) GetRange() int {
    return d.range
}

func (d DetectPoisonAndDisease) GetComponents() string {
    return d.components
}

func (d DetectPoisonAndDisease) GetDuration() int {
    return d.duration
}

func (d DetectPoisonAndDisease) GetConcentration() bool {
    return d.concentration
}

func (d DetectPoisonAndDisease) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Detect Poison and Disease is not implemented yet")
}
