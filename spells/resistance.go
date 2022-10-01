package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Resistance struct {}

func (r Resistance) Cast(target *npcs.NPC) bool {
    log.Println("The spell Resistance is not implemented yet")
}

func (r Resistance) PrettyPrint() string {
    return "Resistance"
}

func (a Resistance) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Resistance is not implemented yet")
}
