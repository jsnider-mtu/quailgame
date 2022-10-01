package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type CreateOrDestroyWater struct {}

func (c CreateOrDestroyWater) Cast(target *npcs.NPC) bool {
    log.Println("The spell Create Or Destroy Water is not implemented yet")
}

func (c CreateOrDestroyWater) PrettyPrint() string {
    return "Create or Destroy Water"
}
