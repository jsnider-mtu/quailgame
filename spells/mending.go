package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Mending struct {}

func (m Mending) Cast(target *npcs.NPC) bool {
    log.Println("The spell Mending is not implemented yet")
}

func (m Mending) PrettyPrint() string {
    return "Mending"
}

func (a Mending) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Mending is not implemented yet")
}
