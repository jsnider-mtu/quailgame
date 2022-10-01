package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type WitchBolt struct {}

func (w WitchBolt) Cast(target *npcs.NPC) bool {
    log.Println("The spell Witch Bolt is not implemented yet")
}

func (w WitchBolt) PrettyPrint() string {
    return "Witch Bolt"
}

func (a WitchBolt) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Witch Bolt is not implemented yet")
}
