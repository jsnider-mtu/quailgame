package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type UnseenServant struct {}

func (u UnseenServant) Cast(target *npcs.NPC) bool {
    log.Println("The spell Unseen Servant is not implemented yet")
}

func (u UnseenServant) PrettyPrint() string {
    return "Unseen Servant"
}

func (a UnseenServant) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Unseen Servant is not implemented yet")
}
