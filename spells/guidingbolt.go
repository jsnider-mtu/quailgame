package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type GuidingBolt struct {}

func (g GuidingBolt) Cast(target *npcs.NPC) bool {
    log.Println("The spell Guiding Bolt is not implemented yet")
}

func (g GuidingBolt) PrettyPrint() string {
    return "Guiding Bolt"
}
