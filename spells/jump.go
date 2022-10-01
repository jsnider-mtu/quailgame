package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Jump struct {}

func (j Jump) Cast(target *npcs.NPC) bool {
    log.Println("The spell Jump is not implemented yet")
}

func (j Jump) PrettyPrint() string {
    return "Jump"
}
