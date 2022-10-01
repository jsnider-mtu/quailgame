package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type TrueStrike struct {}

func (t TrueStrike) Cast(target *npcs.NPC) bool {
    log.Println("The spell True Strike is not implemented yet")
}

func (t TrueStrike) PrettyPrint() string {
    return "True Strike"
}

func (a TrueStrike) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell True Strike is not implemented yet")
}
