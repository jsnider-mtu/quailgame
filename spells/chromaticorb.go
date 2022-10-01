package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ChromaticOrb struct {}

func (c ChromaticOrb) Cast(target *npcs.NPC) bool {
    log.Println("The spell Chromatic Orb is not implemented yet")
}

func (c ChromaticOrb) PrettyPrint() string {
    return "Chromatic Orb"
}

func (a ChromaticOrb) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Chromatic Orb is not implemented yet")
}
