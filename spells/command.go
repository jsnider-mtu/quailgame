package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Command struct {}

func (c Command) Cast(target *npcs.NPC) bool {
    log.Println("The spell Command is not implemented yet")
}

func (c Command) PrettyPrint() string {
    return "Command"
}

func (a Command) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Command is not implemented yet")
}
