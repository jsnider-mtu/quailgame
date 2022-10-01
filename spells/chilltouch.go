package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type ChillTouch struct {}

func (c ChillTouch) Cast(target *npcs.NPC) bool {
    log.Println("The spell Chill Touch is not implemented yet")
}

func (c ChillTouch) PrettyPrint() string {
    return "Chill Touch"
}
