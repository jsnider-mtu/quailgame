package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type UnseenServant struct {}

func (u UnseenServant) Cast(target *npcs.NPC) bool {
    log.Println("The spell Unseen Servant is not implemented yet")
}

func (u UnseenServant) PrettyPrint() string {
    return "Unseen Servant"
}
