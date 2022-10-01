package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type DancingLights struct {}

func (d DancingLights) Cast(target *npcs.NPC) bool {
    log.Println("The spell Dancing Lights is not implemented yet")
}

func (d DancingLights) PrettyPrint() string {
    return "Dancing Lights"
}

func (a DancingLights) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Dancing Lights is not implemented yet")
}
