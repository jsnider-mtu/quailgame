package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ShockingGrasp struct {}

func (s ShockingGrasp) Cast(target *npcs.NPC) bool {
    log.Println("The spell Shocking Grasp is not implemented yet")
}

func (s ShockingGrasp) PrettyPrint() string {
    return "Shocking Grasp"
}
