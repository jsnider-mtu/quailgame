package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type DissonantWhispers struct {}

func (d DissonantWhispers) Cast(target *npcs.NPC) bool {
    log.Println("The spell Dissonant Whispers is not implemented yet")
}

func (d DissonantWhispers) PrettyPrint() string {
    return "Dissonant Whispers"
}

func (a DissonantWhispers) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Dissonant Whispers is not implemented yet")
}
