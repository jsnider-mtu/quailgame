package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ExpeditiousRetreat struct {}

func (e ExpeditiousRetreat) Cast(target *npcs.NPC) bool {
    log.Println("The spell Expeditious Retreat is not implemented yet")
}

func (e ExpeditiousRetreat) PrettyPrint() string {
    return "Expeditious Retreat"
}

func (a ExpeditiousRetreat) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Expeditious Retreat is not implemented yet")
}
