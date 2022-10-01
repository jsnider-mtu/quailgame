package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Heroism struct {}

func (h Heroism) Cast(target *npcs.NPC) bool {
    log.Println("The spell Heroism is not implemented yet")
}

func (h Heroism) PrettyPrint() string {
    return "Heroism"
}
