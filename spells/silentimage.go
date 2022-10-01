package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type SilentImage struct {}

func (s SilentImage) Cast(target *npcs.NPC) bool {
    log.Println("The spell Silent Image is not implemented yet")
}

func (s SilentImage) PrettyPrint() string {
    return "Silent Image"
}

func (a SilentImage) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Silent Image is not implemented yet")
}
