package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type FalseLife struct {}

func (f FalseLife) Cast(target *npcs.NPC) bool {
    log.Println("The spell False Life is not implemented yet")
}

func (f FalseLife) PrettyPrint() string {
    return "False Life"
}
