package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Guidance struct {}

func (g Guidance) Cast(target *npcs.NPC) bool {
    log.Println("The spell Guidance is not implemented yet")
}

func (g Guidance) PrettyPrint() string {
    return "Guidance"
}

func (a Guidance) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Guidance is not implemented yet")
}
