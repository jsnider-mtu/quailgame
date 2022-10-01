package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Bane struct {}

func (b Bane) Cast(target *npcs.NPC) bool {
    log.Println("The spell Bane is not implemented yet")
}

func (b Bane) PrettyPrint() string {
    return "Bane"
}
