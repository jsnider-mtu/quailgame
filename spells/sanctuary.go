package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Sanctuary struct {}

func (s Sanctuary) Cast(target *npcs.NPC) bool {
    log.Println("The spell Sanctuary is not implemented yet")
}

func (s Sanctuary) PrettyPrint() string {
    return "Sanctuary"
}
