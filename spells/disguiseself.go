package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type DisguiseSelf struct {}

func (d DisguiseSelf) Cast(target *npcs.NPC) bool {
    log.Println("The spell Disguise Self is not implemented yet")
}

func (d DisguiseSelf) PrettyPrint() string {
    return "Disguise Self"
}
