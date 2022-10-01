package spells

import "github.com/jsnider-mtu/quailgame/npcs"

type RayOfFrost struct {}

func (r RayOfFrost) Cast(target *npcs.NPC) bool {
    log.Println("The spell Ray of Frost is not implemented yet")
}

func (r RayOfFrost) PrettyPrint() string {
    return "Ray of Frost"
}
