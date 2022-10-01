package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ThornWhip struct {}

func (t ThornWhip) Cast(target *npcs.NPC) bool {
    log.Println("The spell Thorn Whip is not implemented yet")
}

func (t ThornWhip) PrettyPrint() string {
    return "Thorn Whip"
}

func (a ThornWhip) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Thorn Whip is not implemented yet")
}
