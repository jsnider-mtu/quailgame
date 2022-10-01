package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ThunderousSmite struct {}

func (t ThunderousSmite) Cast(target *npcs.NPC) bool {
    log.Println("The spell Thunderous Smite is not implemented yet")
}

func (t ThunderousSmite) PrettyPrint() string {
    return "Thunderous Smite"
}

func (a ThunderousSmite) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Thunderous Smite is not implemented yet")
}
