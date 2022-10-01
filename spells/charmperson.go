package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type CharmPerson struct {}

func (c CharmPerson) Cast(target *npcs.NPC) bool {
    log.Println("The spell Charm Person is not implemented yet")
}

func (c CharmPerson) PrettyPrint() string {
    return "Charm Person"
}

func (a CharmPerson) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Charm Person is not implemented yet")
}
