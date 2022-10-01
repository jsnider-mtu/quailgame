package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Heroism struct {}

func (h Heroism) Cast(target *npcs.NPC) bool {
    log.Println("The spell Heroism is not implemented yet")
}

func (h Heroism) PrettyPrint() string {
    return "Heroism"
}

func (a Heroism) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Heroism is not implemented yet")
}
