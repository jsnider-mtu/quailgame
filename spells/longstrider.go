package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Longstrider struct {}

func (l Longstrider) Cast(target *npcs.NPC) bool {
    log.Println("The spell Longstrider is not implemented yet")
}

func (l Longstrider) PrettyPrint() string {
    return "Longstrider"
}

func (a Longstrider) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Longstrider is not implemented yet")
}
