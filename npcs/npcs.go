package npcs

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/projectx/levels"
    "github.com/jsnider-mtu/projectx/utils"
)

type NPC struct {
    Name string
    Msgs []string
    Img *ebiten.Image
    Speed int // Speed of 0 means no movement
    Pos [2]int
}

func (npc *NPC) Move(l *levels.Level, direction string) {
    var newpos [2]int
    switch direction {
    case "up":
        newpos = utils.TryUpdatePos(false, npc.Pos, l, true, -24)
    case "down":
        newpos = utils.TryUpdatePos(false, npc.Pos, l, true, 24)
    case "left":
        newpos = utils.TryUpdatePos(false, npc.Pos, l, false, -24)
    case "right":
        newpos = utils.TryUpdatePos(false, npc.Pos, l, false, 24)
    default:
        newpos = npc.Pos
    }
    npc.Pos = newpos
}
