package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type MinorIllusion struct {}

func (m MinorIllusion) Cast(target *npcs.NPC) bool {
    log.Println("The spell Minor Illusion is not implemented yet")
}

func (m MinorIllusion) PrettyPrint() string {
    return "Minor Illusion"
}

func (a MinorIllusion) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Minor Illusion is not implemented yet")
}
