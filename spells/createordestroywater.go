package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type CreateOrDestroyWater struct {}

func (c CreateOrDestroyWater) Cast(target *npcs.NPC) bool {
    log.Println("The spell Create Or Destroy Water is not implemented yet")
}

func (c CreateOrDestroyWater) PrettyPrint() string {
    return "Create or Destroy Water"
}

func (a CreateOrDestroyWater) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Create or Destroy Water is not implemented yet")
}
