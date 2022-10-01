package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type AcidSplash struct {}

func (a AcidSplash) Cast(target *npcs.NPC) bool {
    log.Println("The spell Acid Splash is not implemented yet")
}

func (a AcidSplash) PrettyPrint() string {
    return "Acid Splash"
}

func (a AcidSplash) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Acid Splash is not implemented yet")
}
