package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Thunderwave struct {}

func (t Thunderwave) Cast(target *npcs.NPC) bool {
    log.Println("The spell Thunderwave is not implemented yet")
}

func (t Thunderwave) PrettyPrint() string {
    return "Thunderwave"
}
