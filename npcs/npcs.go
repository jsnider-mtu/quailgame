package npcs

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/projectx/levels"
)

type NPC struct {
    Name string
    Msgs []string
    Img *ebiten.Image
    Speed int // Speed of 0 means no movement
    Pos [2]int
}

func (npc *NPC) Move(l *levels.Level) {
    levels.TryUpdatePos(false, npc.Pos, l, true, 24)
}
