package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Druidcraft struct {}

func (d Druidcraft) Cast(target *npcs.NPC) bool {
    log.Println("The spell Druidcraft is not implemented yet")
}

func (d Druidcraft) PrettyPrint() string {
    return "Druidcraft"
}
