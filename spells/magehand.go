package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type MageHand struct {}

func (m MageHand) Cast(target *npcs.NPC) bool {
    log.Println("The spell Mage Hand is not implemented yet")
}

func (m MageHand) PrettyPrint() string {
    return "Mage Hand"
}
