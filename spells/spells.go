package spells

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/npcs"
)

type Spell interface {
    Cast(*npcs.NPC) bool
    PrettyPrint() string
    Animate(*ebiten.Image, *npcs.NPC)
}
