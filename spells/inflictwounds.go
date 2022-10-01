package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type InflictWounds struct {}

func (i InflictWounds) Cast(target *npcs.NPC) bool {
    log.Println("The spell Inflict Wounds is not implemented yet")
}

func (i InflictWounds) PrettyPrint() string {
    return "Inflict Wounds"
}
