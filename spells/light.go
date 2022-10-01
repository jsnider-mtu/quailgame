package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Light struct {}

func (l Light) Cast(target *npcs.NPC) bool {
    log.Println("The spell Light is not implemented yet")
}

func (l Light) PrettyPrint() string {
    return "Light"
}
