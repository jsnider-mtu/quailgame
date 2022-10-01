package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type HealingWord struct {}

func (h HealingWord) Cast(target *npcs.NPC) bool {
    log.Println("The spell Healing Word is not implemented yet")
}

func (h HealingWord) PrettyPrint() string {
    return "Healing Word"
}

func (a HealingWord) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Healing Word is not implemented yet")
}
