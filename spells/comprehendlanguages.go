package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ComprehendLanguages struct {}

func (c ComprehendLanguages) Cast(target *npcs.NPC) bool {
    log.Println("The spell Comprehend Languages is not implemented yet")
}

func (c ComprehendLanguages) PrettyPrint() string {
    return "Comprehend Languages"
}

func (a ComprehendLanguages) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Comprehend Languages is not implemented yet")
}
