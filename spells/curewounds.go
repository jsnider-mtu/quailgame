package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type CureWounds struct {}

func (c CureWounds) Cast(target *npcs.NPC) bool {
    log.Println("The spell Cure Wounds is not implemented yet")
}

func (c CureWounds) PrettyPrint() string {
    return "Cure Wounds"
}

func (a CureWounds) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Cure Wounds is not implemented yet")
}
