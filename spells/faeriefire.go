package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type FaerieFire struct {}

func (f FaerieFire) Cast(target *npcs.NPC) bool {
    log.Println("The spell Faerie Fire is not implemented yet")
}

func (f FaerieFire) PrettyPrint() string {
    return "Faerie Fire"
}

func (a FaerieFire) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Faerie Fire is not implemented yet")
}
