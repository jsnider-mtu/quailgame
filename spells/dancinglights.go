package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type DancingLights struct {}

func (d DancingLights) Cast(target *npcs.NPC) bool {
    log.Println("The spell Dancing Lights is not implemented yet")
}

func (d DancingLights) PrettyPrint() string {
    return "Dancing Lights"
}
