package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type FaerieFire struct {}

func (f FaerieFire) Cast(target *npcs.NPC) bool {
    log.Println("The spell Faerie Fire is not implemented yet")
}

func (f FaerieFire) PrettyPrint() string {
    return "Faerie Fire"
}
