package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type SpareTheDying struct {}

func (s SpareTheDying) Cast(target *npcs.NPC) bool {
    log.Println("The spell Spare the Dying is not implemented yet")
}

func (s SpareTheDying) PrettyPrint() string {
    return "Spare the Dying"
}

func (a SpareTheDying) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Spare the Dying is not implemented yet")
}
