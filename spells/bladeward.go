package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type BladeWard struct {}

func (b BladeWard) Cast(target *npcs.NPC) bool {
    log.Println("The spell Blade Ward is not implemented yet")
}

func (b BladeWard) PrettyPrint() string {
    return "Blade Ward"
}

func (a BladeWard) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Blade Ward is not implemented yet")
}
