package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type Longstrider struct {}

func (l Longstrider) Cast(target *npcs.NPC) bool {
    log.Println("The spell Longstrider is not implemented yet")
}

func (l Longstrider) PrettyPrint() string {
    return "Longstrider"
}
