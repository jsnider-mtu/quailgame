package spells

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
)

type CharmPerson struct {}

func (c CharmPerson) PrettyPrint() string {
    return "Charm Person"
}

func (a CharmPerson) Animate(screen *ebiten.Image, targetpos [2]int) {
    log.Println("The animation for the spell Charm Person is not implemented yet")
}
