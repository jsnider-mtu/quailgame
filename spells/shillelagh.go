package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Shillelagh struct {}

func (s Shillelagh) Cast(target *npcs.NPC) bool {
    log.Println("The spell Shillelagh is not implemented yet")
}

func (s Shillelagh) PrettyPrint() string {
    return "Shillelagh"
}

func (a Shillelagh) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Shillelagh is not implemented yet")
}
