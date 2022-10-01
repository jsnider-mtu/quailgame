package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type PoisonSpray struct {}

func (p PoisonSpray) Cast(target *npcs.NPC) bool {
    log.Println("The spell Poison Spray is not implemented yet")
}

func (p PoisonSpray) PrettyPrint() string {
    return "Poison Spray"
}
