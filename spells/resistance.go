package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Resistance struct {}

func (r Resistance) Cast(target *npcs.NPC) bool {
    log.Println("The spell Resistance is not implemented yet")
}

func (r Resistance) PrettyPrint() string {
    return "Resistance"
}
