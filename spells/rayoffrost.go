package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type RayOfFrost struct {}

func (r RayOfFrost) Cast(target *npcs.NPC) bool {
    log.Println("The spell Ray of Frost is not implemented yet")
}

func (r RayOfFrost) PrettyPrint() string {
    return "Ray of Frost"
}

func (a RayOfFrost) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Ray of Frost is not implemented yet")
}
