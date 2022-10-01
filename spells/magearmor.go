package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type MageArmor struct {}

func (m MageArmor) Cast(target *npcs.NPC) bool {
    log.Println("The spell Mage Armor is not implemented yet")
}

func (m MageArmor) PrettyPrint() string {
    return "Mage Armor"
}
