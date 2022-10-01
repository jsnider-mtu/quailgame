package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type BurningHands struct {}

func (b BurningHands) Cast(target *npcs.NPC) bool {
    log.Println("The spell Burning Hands is not implemented yet")
}

func (b BurningHands) PrettyPrint() string {
    return "Burning Hands"
}
