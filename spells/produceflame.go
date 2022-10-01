package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ProduceFlame struct {}

func (p ProduceFlame) Cast(target *npcs.NPC) bool {
    log.Println("The spell Produce Flame is not implemented yet")
}

func (p ProduceFlame) PrettyPrint() string {
    return "Produce Flame"
}

func (a ProduceFlame) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Produce Flame is not implemented yet")
}
