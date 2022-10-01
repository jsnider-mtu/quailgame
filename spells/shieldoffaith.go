package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type ShieldOfFaith struct {}

func (s ShieldOfFaith) Cast(target *npcs.NPC) bool {
    log.Println("The spell Shield of Faith is not implemented yet")
}

func (s ShieldOfFaith) PrettyPrint() string {
    return "Shield of Faith"
}

func (a ShieldOfFaith) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Shield of Faith is not implemented yet")
}
