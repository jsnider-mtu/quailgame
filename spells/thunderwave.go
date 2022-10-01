package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Thunderwave struct {}

func (t Thunderwave) Cast(target *npcs.NPC) bool {
    log.Println("The spell Thunderwave is not implemented yet")
}

func (t Thunderwave) PrettyPrint() string {
    return "Thunderwave"
}

func (a Thunderwave) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Thunderwave is not implemented yet")
}
