package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Grease struct {}

func (g Grease) Cast(target *npcs.NPC) bool {
    log.Println("The spell Grease is not implemented yet")
}

func (g Grease) PrettyPrint() string {
    return "Grease"
}

func (a Grease) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Grease is not implemented yet")
}
