package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Entangle struct {}

func (e Entangle) Cast(target *npcs.NPC) bool {
    log.Println("The spell Entangle is not implemented yet")
}

func (e Entangle) PrettyPrint() string {
    return "Entangle"
}
