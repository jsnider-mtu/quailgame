package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type DetectPoisonAndDisease struct {}

func (d DetectPoisonAndDisease) Cast(target *npcs.NPC) bool {
    log.Println("The spell Detect Poison and Disease is not implemented yet")
}

func (d DetectPoisonAndDisease) PrettyPrint() string {
    return "Detect Poison and Disease"
}

func (a DetectPoisonAndDisease) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Detect Poison and Disease is not implemented yet")
}
