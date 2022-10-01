package spells

import (
    "github.com/hajimehoshi/ebiten/v2"
)

type Spell interface {
    PrettyPrint() string
    GetLevel() int
    Animate(*ebiten.Image, [2]int)
}
