package spells

import (
    "github.com/hajimehoshi/ebiten/v2"
)

type Spell interface {
    PrettyPrint() string
    Animate(*ebiten.Image, [2]int)
}
