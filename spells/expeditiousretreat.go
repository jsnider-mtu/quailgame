package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ExpeditiousRetreat struct {}

func (e ExpeditiousRetreat) Cast(target *npcs.NPC) bool {
    log.Println("The spell Expeditious Retreat is not implemented yet")
}

func (e ExpeditiousRetreat) PrettyPrint() string {
    return "Expeditious Retreat"
}
