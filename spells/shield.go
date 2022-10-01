package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Shield struct {}

func (s Shield) Cast(target *npcs.NPC) bool {
    log.Println("The spell Shield is not implemented yet")
}

func (s Shield) PrettyPrint() string {
    return "Shield"
}

func (a Shield) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Shield is not implemented yet")
}
