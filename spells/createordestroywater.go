package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CreateOrDestroyWater struct {}

func (c CreateOrDestroyWater) PrettyPrint() string {
    return "Create or Destroy Water"
}

func (c CreateOrDestroyWater) GetLevel() int {
    return 0
}

func (a CreateOrDestroyWater) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Create or Destroy Water is not implemented yet")
}
