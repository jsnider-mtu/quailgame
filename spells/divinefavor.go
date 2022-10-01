package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type DivineFavor struct {}

func (d DivineFavor) Cast(target *npcs.NPC) bool {
    log.Println("The spell Divine Favor is not implemented yet")
}

func (d DivineFavor) PrettyPrint() string {
    return "Divine Favor"
}
