package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type RayOfSickness struct {}

func (r RayOfSickness) Cast(target *npcs.NPC) bool {
    log.Println("The spell Ray of Sickness is not implemented yet")
}

func (r RayOfSickness) PrettyPrint() string {
    return "Ray of Sickness"
}

func (a RayOfSickness) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Ray of Sickness is not implemented yet")
}
