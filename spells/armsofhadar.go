package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ArmsOfHadar struct {}

func (a ArmsOfHadar) Cast(target *npcs.NPC) bool {
    log.Println("The spell Arms of Hadar is not implemented yet")
}

func (a ArmsOfHadar) PrettyPrint() string {
    return "Arms of Hadar"
}

func (a ArmsOfHadar) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Arms of Hadar is not implemented yet")
}
