package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Hex struct {}

func (h Hex) Cast(target *npcs.NPC) bool {
    log.Println("The spell Hex is not implemented yet")
}

func (h Hex) PrettyPrint() string {
    return "Hex"
}

func (a Hex) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Hex is not implemented yet")
}
