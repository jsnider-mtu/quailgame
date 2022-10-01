package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type SearingSmite struct {}

func (s SearingSmite) Cast(target *npcs.NPC) bool {
    log.Println("The spell Searing Smite is not implemented yet")
}

func (s SearingSmite) PrettyPrint() string {
    return "Searing Smite"
}

func (a SearingSmite) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Searing Smite is not implemented yet")
}
