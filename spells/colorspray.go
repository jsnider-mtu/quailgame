package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ColorSpray struct {}

func (c ColorSpray) Cast(target *npcs.NPC) bool {
    log.Println("The spell Color Spray is not implemented yet")
}

func (c ColorSpray) PrettyPrint() string {
    return "Color Spray"
}

func (a ColorSpray) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Color Spray is not implemented yet")
}
