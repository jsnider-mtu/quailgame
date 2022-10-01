package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Entangle struct {}

func (e Entangle) Cast(target *npcs.NPC) bool {
    log.Println("The spell Entangle is not implemented yet")
}

func (e Entangle) PrettyPrint() string {
    return "Entangle"
}

func (a Entangle) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Entangle is not implemented yet")
}
