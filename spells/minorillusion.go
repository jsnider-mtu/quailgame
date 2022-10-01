package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type MinorIllusion struct {}

func (m MinorIllusion) Cast(target *npcs.NPC) bool {
    log.Println("The spell Minor Illusion is not implemented yet")
}

func (m MinorIllusion) PrettyPrint() string {
    return "Minor Illusion"
}
