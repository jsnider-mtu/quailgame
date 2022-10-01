package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type DetectPoisonAndDisease struct {}

func (d DetectPoisonAndDisease) Cast(target *npcs.NPC) bool {
    log.Println("The spell Detect Poison and Disease is not implemented yet")
}

func (d DetectPoisonAndDisease) PrettyPrint() string {
    return "Detect Poison and Disease"
}
