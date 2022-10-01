package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ShockingGrasp struct {}

func (s ShockingGrasp) Cast(target *npcs.NPC) bool {
    log.Println("The spell Shocking Grasp is not implemented yet")
}

func (s ShockingGrasp) PrettyPrint() string {
    return "Shocking Grasp"
}

func (a ShockingGrasp) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Shocking Grasp is not implemented yet")
}
