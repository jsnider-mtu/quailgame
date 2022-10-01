package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Grease struct {}

func (g Grease) Cast(target *npcs.NPC) bool {
    log.Println("The spell Grease is not implemented yet")
}

func (g Grease) PrettyPrint() string {
    return "Grease"
}
