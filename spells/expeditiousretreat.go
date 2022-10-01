package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type ExpeditiousRetreat struct {}

func (e ExpeditiousRetreat) PrettyPrint() string {
    return "Expeditious Retreat"
}

func (a ExpeditiousRetreat) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Expeditious Retreat is not implemented yet")
}
