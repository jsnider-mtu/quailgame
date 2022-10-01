package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type MagicMissile struct {}

func (m MagicMissile) Cast(target *npcs.NPC) bool {
    log.Println("The spell Magic Missile is not implemented yet")
}

func (m MagicMissile) PrettyPrint() string {
    return "Magic Missile"
}

func (a MagicMissile) Animate(screen *ebiten.Image, target *npcs.NPC) {
    log.Println("The animation for the spell Magic Missile is not implemented yet")
}
