package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Sanctuary struct {}

func (s Sanctuary) Cast(target *npcs.NPC) bool {
    log.Println("The spell Sanctuary is not implemented yet")
}

func (s Sanctuary) PrettyPrint() string {
    return "Sanctuary"
}

func (a Sanctuary) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Sanctuary is not implemented yet")
}
