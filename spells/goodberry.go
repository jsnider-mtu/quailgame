package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Goodberry struct {}

func (g Goodberry) Cast(target *npcs.NPC) bool {
    log.Println("The spell Goodberry is not implemented yet")
}

func (g Goodberry) PrettyPrint() string {
    return "Goodberry"
}
