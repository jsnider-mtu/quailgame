package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Identify struct {}

func (i Identify) Cast(target *npcs.NPC) bool {
    log.Println("The spell Identify is not implemented yet")
}

func (i Identify) PrettyPrint() string {
    return "Identify"
}

func (a Identify) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Identify is not implemented yet")
}
