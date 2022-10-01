package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type FalseLife struct {}

func (f FalseLife) Cast(target *npcs.NPC) bool {
    log.Println("The spell False Life is not implemented yet")
}

func (f FalseLife) PrettyPrint() string {
    return "False Life"
}

func (a FalseLife) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell False Life is not implemented yet")
}
