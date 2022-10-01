package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type WrathfulSmite struct {}

func (w WrathfulSmite) Cast(target *npcs.NPC) bool {
    log.Println("The spell Wrathful Smite is not implemented yet")
}

func (w WrathfulSmite) PrettyPrint() string {
    return "Wrathful Smite"
}

func (a WrathfulSmite) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Wrathful Smite is not implemented yet")
}
