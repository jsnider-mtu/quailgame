package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type FogCloud struct {}

func (f FogCloud) Cast(target *npcs.NPC) bool {
    log.Println("The spell Fog Cloud is not implemented yet")
}

func (f FogCloud) PrettyPrint() string {
    return "Fog Cloud"
}

func (a FogCloud) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Fog Cloud is not implemented yet")
}
