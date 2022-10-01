package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Druidcraft struct {}

func (d Druidcraft) Cast(target *npcs.NPC) bool {
    log.Println("The spell Druidcraft is not implemented yet")
}

func (d Druidcraft) PrettyPrint() string {
    return "Druidcraft"
}

func (a Druidcraft) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Druidcraft is not implemented yet")
}
