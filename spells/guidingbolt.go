package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type GuidingBolt struct {}

func (g GuidingBolt) Cast(target *npcs.NPC) bool {
    log.Println("The spell Guiding Bolt is not implemented yet")
}

func (g GuidingBolt) PrettyPrint() string {
    return "Guiding Bolt"
}

func (a GuidingBolt) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Guiding Bolt is not implemented yet")
}
