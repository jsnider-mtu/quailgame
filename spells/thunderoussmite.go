package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ThunderousSmite struct {}

func (t ThunderousSmite) Cast(target *npcs.NPC) bool {
    log.Println("The spell Thunderous Smite is not implemented yet")
}

func (t ThunderousSmite) PrettyPrint() string {
    return "Thunderous Smite"
}
