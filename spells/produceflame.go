package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ProduceFlame struct {}

func (p ProduceFlame) Cast(target *npcs.NPC) bool {
    log.Println("The spell Produce Flame is not implemented yet")
}

func (p ProduceFlame) PrettyPrint() string {
    return "Produce Flame"
}
