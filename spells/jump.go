package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Jump struct {}

func (j Jump) Cast(target *npcs.NPC) bool {
    log.Println("The spell Jump is not implemented yet")
}

func (j Jump) PrettyPrint() string {
    return "Jump"
}

func (a Jump) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Jump is not implemented yet")
}
