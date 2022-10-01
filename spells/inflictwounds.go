package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type InflictWounds struct {}

func (i InflictWounds) Cast(target *npcs.NPC) bool {
    log.Println("The spell Inflict Wounds is not implemented yet")
}

func (i InflictWounds) PrettyPrint() string {
    return "Inflict Wounds"
}

func (a InflictWounds) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Inflict Wounds is not implemented yet")
}
