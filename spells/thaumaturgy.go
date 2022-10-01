package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Thaumaturgy struct {}

func (t Thaumaturgy) Cast(target *npcs.NPC) bool {
    log.Println("The spell Thaumaturgy is not implemented yet")
}

func (t Thaumaturgy) PrettyPrint() string {
    return "Thaumaturgy"
}
