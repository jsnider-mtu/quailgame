package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Bless struct {}

func (b Bless) Cast(target *npcs.NPC) bool {
    log.Println("The spell Bless is not implemented yet")
}

func (b Bless) PrettyPrint() string {
    return "Bless"
}
