package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Sleep struct {}

func (s Sleep) Cast(target *npcs.NPC) bool {
    log.Println("The spell Sleep is not implemented yet")
}

func (s Sleep) PrettyPrint() string {
    return "Sleep"
}
