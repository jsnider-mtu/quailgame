package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type TashasHideousLaughter struct {}

func (t TashasHideousLaughter) Cast(target *npcs.NPC) bool {
    log.Println("The spell Tasha's Hideous Laughter is not implemented yet")
}

func (t TashasHideousLaughter) PrettyPrint() string {
    return "Tasha's Hideous Laughter"
}

func (a TashasHideousLaughter) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Tasha's Hideous Laughter is not implemented yet")
}
