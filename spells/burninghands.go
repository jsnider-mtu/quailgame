package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type BurningHands struct {}

func (b BurningHands) Cast(target *npcs.NPC) bool {
    log.Println("The spell Burning Hands is not implemented yet")
}

func (b BurningHands) PrettyPrint() string {
    return "Burning Hands"
}

func (a BurningHands) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Burning Hands is not implemented yet")
}
