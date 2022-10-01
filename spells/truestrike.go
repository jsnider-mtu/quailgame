package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type TrueStrike struct {}

func (t TrueStrike) Cast(target *npcs.NPC) bool {
    log.Println("The spell True Strike is not implemented yet")
}

func (t TrueStrike) PrettyPrint() string {
    return "True Strike"
}
