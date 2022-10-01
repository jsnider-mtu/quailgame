package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type WitchBolt struct {}

func (w WitchBolt) Cast(target *npcs.NPC) bool {
    log.Println("The spell Witch Bolt is not implemented yet")
}

func (w WitchBolt) PrettyPrint() string {
    return "Witch Bolt"
}
