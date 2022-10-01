package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Shield struct {}

func (s Shield) Cast(target *npcs.NPC) bool {
    log.Println("The spell Shield is not implemented yet")
}

func (s Shield) PrettyPrint() string {
    return "Shield"
}
