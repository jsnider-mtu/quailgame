package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type HellishRebuke struct {}

func (h HellishRebuke) Cast(target *npcs.NPC) bool {
    log.Println("The spell Hellish Rebuke is not implemented yet")
}

func (h HellishRebuke) PrettyPrint() string {
    return "Hellish Rebuke"
}

func (a HellishRebuke) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Hellish Rebuke is not implemented yet")
}
