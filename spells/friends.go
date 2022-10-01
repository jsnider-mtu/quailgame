package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Friends struct {}

func (f Friends) Cast(target *npcs.NPC) bool {
    log.Println("The spell Friends is not implemented yet")
}

func (f Friends) PrettyPrint() string {
    return "Friends"
}
