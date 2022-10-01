package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type CompelledDuel struct {}

func (c CompelledDuel) Cast(target *npcs.NPC) bool {
    log.Println("The spell Compelled Duel is not implemented yet")
}

func (c CompelledDuel) PrettyPrint() string {
    return "Compelled Duel"
}

func (a CompelledDuel) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Compelled Duel is not implemented yet")
}
