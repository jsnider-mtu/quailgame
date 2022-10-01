package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Prestidigitation struct {}

func (p Prestidigitation) Cast(target *npcs.NPC) bool {
    log.Println("The spell Prestidigitation is not implemented yet")
}

func (p Prestidigitation) PrettyPrint() string {
    return "Prestidigitation"
}
