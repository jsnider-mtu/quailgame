package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type CureWounds struct {}

func (c CureWounds) Cast(target *npcs.NPC) bool {
    log.Println("The spell Cure Wounds is not implemented yet")
}

func (c CureWounds) PrettyPrint() string {
    return "Cure Wounds"
}
