package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type HuntersMark struct {}

func (h HuntersMark) Cast(target *npcs.NPC) bool {
    log.Println("The spell Hunter's Mark is not implemented yet")
}

func (h HuntersMark) PrettyPrint() string {
    return "Hunter's Mark"
}
