package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Prestidigitation struct {}

func (p Prestidigitation) Cast(target *npcs.NPC) bool {
    log.Println("The spell Prestidigitation is not implemented yet")
}

func (p Prestidigitation) PrettyPrint() string {
    return "Prestidigitation"
}

func (a Prestidigitation) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Prestidigitation is not implemented yet")
}
