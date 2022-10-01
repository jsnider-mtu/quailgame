package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type HailOfThorns struct {}

func (h HailOfThorns) Cast(target *npcs.NPC) bool {
    log.Println("The spell Hail of Thorns is not implemented yet")
}

func (h HailOfThorns) PrettyPrint() string {
    return "Hail of Thorns"
}

func (a HailOfThorns) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Hail of Thorns is not implemented yet")
}
