package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ArmsOfHadar struct {}

func (a ArmsOfHadar) Cast(target *npcs.NPC) bool {
    log.Println("The spell Arms of Hadar is not implemented yet")
}

func (a ArmsOfHadar) PrettyPrint() string {
    return "Arms of Hadar"
}
