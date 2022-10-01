package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type CompelledDuel struct {}

func (c CompelledDuel) Cast(target *npcs.NPC) bool {
    log.Println("The spell Compelled Duel is not implemented yet")
}

func (c CompelledDuel) PrettyPrint() string {
    return "Compelled Duel"
}
