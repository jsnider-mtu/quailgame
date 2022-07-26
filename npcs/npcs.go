package npcs

import (
//    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/projectx/player"
//    "github.com/jsnider-mtu/projectx/levels"
//    "github.com/jsnider-mtu/projectx/utils"
)

type NPC struct {
    Name string
    Msgs [][]string
    MsgCount int
    Speed int // Speed of 0 means no movement
    Direction string
    PC *player.Player
}

func (npc *NPC) Dialog() []string {
    if npc.MsgCount == len(npc.Msgs) {
        npc.MsgCount = 0
    }
    npc.MsgCount++
    return npc.Msgs[npc.MsgCount - 1]
}
