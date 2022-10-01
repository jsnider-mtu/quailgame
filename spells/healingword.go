package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type HealingWord struct {}

func (h HealingWord) Cast(target *npcs.NPC) bool {
    log.Println("The spell Healing Word is not implemented yet")
}

func (h HealingWord) PrettyPrint() string {
    return "Healing Word"
}
