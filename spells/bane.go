package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Bane struct {}

func (b Bane) Cast(target *npcs.NPC) bool {
    log.Println("The spell Bane is not implemented yet")
}

func (b Bane) PrettyPrint() string {
    return "Bane"
}

func (a Bane) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Bane is not implemented yet")
}
