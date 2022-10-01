package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type FeatherFall struct {}

func (f FeatherFall) Cast(target *npcs.NPC) bool {
    log.Println("The spell Feather Fall is not implemented yet")
}

func (f FeatherFall) PrettyPrint() string {
    return "Feather Fall"
}
