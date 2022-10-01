package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Identify struct {}

func (i Identify) Cast(target *npcs.NPC) bool {
    log.Println("The spell Identify is not implemented yet")
}

func (i Identify) PrettyPrint() string {
    return "Identify"
}
