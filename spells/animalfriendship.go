package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type AnimalFriendship struct {
    level int
    casttime string
    range int
    components string
    duration int
    concentration bool
}

func (a AnimalFriendship) PrettyPrint() string {
    return "Animal Friendship"
}

func (a AnimalFriendship) GetLevel() int {
    return a.level
}

func (a AnimalFriendship) GetCastTime() string {
    return a.casttime
}

func (a AnimalFriendship) GetRange() int {
    return a.range
}

func (a AnimalFriendship) GetComponents() string {
    return a.components
}

func (a AnimalFriendship) GetDuration() int {
    return a.duration
}

func (a AnimalFriendship) GetConcentration() bool {
    return a.concentration
}

func (a AnimalFriendship) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Animal Friendship is not implemented yet")
}
