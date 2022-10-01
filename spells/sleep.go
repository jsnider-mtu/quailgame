package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Sleep struct {}

func (s Sleep) Cast(target *npcs.NPC) bool {
    log.Println("The spell Sleep is not implemented yet")
}

func (s Sleep) PrettyPrint() string {
    return "Sleep"
}

func (a Sleep) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Sleep is not implemented yet")
}
