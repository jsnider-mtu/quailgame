package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Bless struct {}

func (b Bless) Cast(target *npcs.NPC) bool {
    log.Println("The spell Bless is not implemented yet")
}

func (b Bless) PrettyPrint() string {
    return "Bless"
}

func (a Bless) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Bless is not implemented yet")
}
