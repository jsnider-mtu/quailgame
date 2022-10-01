package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Hex struct {}

func (h Hex) Cast(target *npcs.NPC) bool {
    log.Println("The spell Hex is not implemented yet")
}

func (h Hex) PrettyPrint() string {
    return "Hex"
}
