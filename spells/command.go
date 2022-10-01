package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Command struct {}

func (c Command) Cast(target *npcs.NPC) bool {
    log.Println("The spell Command is not implemented yet")
}

func (c Command) PrettyPrint() string {
    return "Command"
}
