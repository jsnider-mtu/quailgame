package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ViciousMockery struct {}

func (v ViciousMockery) Cast(target *npcs.NPC) bool {
    log.Println("The spell Vicious Mockery is not implemented yet")
}

func (v ViciousMockery) PrettyPrint() string {
    return "Vicious Mockery"
}

func (a ViciousMockery) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Vicious Mockery is not implemented yet")
}
