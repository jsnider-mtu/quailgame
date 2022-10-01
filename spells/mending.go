package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Mending struct {}

func (m Mending) Cast(target *npcs.NPC) bool {
    log.Println("The spell Mending is not implemented yet")
}

func (m Mending) PrettyPrint() string {
    return "Mending"
}
