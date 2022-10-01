package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Light struct {}

func (l Light) Cast(target *npcs.NPC) bool {
    log.Println("The spell Light is not implemented yet")
}

func (l Light) PrettyPrint() string {
    return "Light"
}

func (a Light) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Light is not implemented yet")
}
