package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Friends struct {}

func (f Friends) Cast(target *npcs.NPC) bool {
    log.Println("The spell Friends is not implemented yet")
}

func (f Friends) PrettyPrint() string {
    return "Friends"
}

func (a Friends) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Friends is not implemented yet")
}
