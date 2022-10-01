package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type EnsnaringStrike struct {}

func (e EnsnaringStrike) Cast(target *npcs.NPC) bool {
    log.Println("The spell Ensnaring Strike is not implemented yet")
}

func (e EnsnaringStrike) PrettyPrint() string {
    return "Ensnaring Strike"
}

func (a EnsnaringStrike) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Ensnaring Strike is not implemented yet")
}
