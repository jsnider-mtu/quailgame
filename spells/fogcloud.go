package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type FogCloud struct {}

func (f FogCloud) Cast(target *npcs.NPC) bool {
    log.Println("The spell Fog Cloud is not implemented yet")
}

func (f FogCloud) PrettyPrint() string {
    return "Fog Cloud"
}
