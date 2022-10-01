package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type SacredFlame struct {}

func (s SacredFlame) Cast(target *npcs.NPC) bool {
    log.Println("The spell Sacred Flame is not implemented yet")
}

func (s SacredFlame) PrettyPrint() string {
    return "Sacred Flame"
}
