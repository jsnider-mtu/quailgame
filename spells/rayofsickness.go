package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type RayOfSickness struct {}

func (r RayOfSickness) Cast(target *npcs.NPC) bool {
    log.Println("The spell Ray of Sickness is not implemented yet")
}

func (r RayOfSickness) PrettyPrint() string {
    return "Ray of Sickness"
}
