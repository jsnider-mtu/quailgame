package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Shillelagh struct {}

func (s Shillelagh) Cast(target *npcs.NPC) bool {
    log.Println("The spell Shillelagh is not implemented yet")
}

func (s Shillelagh) PrettyPrint() string {
    return "Shillelagh"
}
