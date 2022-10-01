package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type EldritchBlast struct {}

func (e EldritchBlast) PrettyPrint() string {
    return "Eldritch Blast"
}

func (a EldritchBlast) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Eldritch Blast is not implemented yet")
}
