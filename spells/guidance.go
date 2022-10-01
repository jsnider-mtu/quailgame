package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Guidance struct {}

func (g Guidance) Cast(target *npcs.NPC) bool {
    log.Println("The spell Guidance is not implemented yet")
}

func (g Guidance) PrettyPrint() string {
    return "Guidance"
}
