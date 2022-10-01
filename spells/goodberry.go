package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Goodberry struct {}

func (g Goodberry) Cast(target *npcs.NPC) bool {
    log.Println("The spell Goodberry is not implemented yet")
}

func (g Goodberry) PrettyPrint() string {
    return "Goodberry"
}

func (a Goodberry) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Goodberry is not implemented yet")
}
