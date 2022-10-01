package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type DetectMagic struct {}

func (d DetectMagic) Cast(target *npcs.NPC) bool {
    log.Println("The spell Detect Magic is not implemented yet")
}

func (d DetectMagic) PrettyPrint() string {
    return "Detect Magic"
}

func (a DetectMagic) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Detect Magic is not implemented yet")
}
