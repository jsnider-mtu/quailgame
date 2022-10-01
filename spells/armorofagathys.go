package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ArmorOfAgathys struct {}

func (a ArmorOfAgathys) Cast(target *npcs.NPC) bool {
    log.Println("The spell Armor of Agathys is not implemented yet")
}

func (a ArmorOfAgathys) PrettyPrint() string {
    return "Armor of Agathys"
}
