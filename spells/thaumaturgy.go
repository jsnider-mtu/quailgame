package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Thaumaturgy struct {}

func (t Thaumaturgy) Cast(target *npcs.NPC) bool {
    log.Println("The spell Thaumaturgy is not implemented yet")
}

func (t Thaumaturgy) PrettyPrint() string {
    return "Thaumaturgy"
}

func (a Thaumaturgy) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Thaumaturgy is not implemented yet")
}
