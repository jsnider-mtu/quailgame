package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type AnimalFriendship struct {}

func (a AnimalFriendship) PrettyPrint() string {
    return "Animal Friendship"
}

func (a AnimalFriendship) GetLevel() int {
    return 1
}

func (a AnimalFriendship) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Animal Friendship is not implemented yet")
}
