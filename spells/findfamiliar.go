package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type FindFamiliar struct {}

func (f FindFamiliar) Cast(target *npcs.NPC) bool {
    log.Println("The spell Find Familiar is not implemented yet")
}

func (f FindFamiliar) PrettyPrint() string {
    return "Find Familiar"
}
