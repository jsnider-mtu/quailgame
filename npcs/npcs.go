package npcs

import (
//    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/player"
//    "github.com/jsnider-mtu/quailgame/levels"
//    "github.com/jsnider-mtu/quailgame/utils"
)

type NPC struct {
    Name string
    Msgs [][]string
    MsgCount int
    Speed int
    Offset int
    Direction string
    Stopped bool
    PC *player.Player
}

func (npc *NPC) Dialog() []string {
    if npc.MsgCount == len(npc.Msgs) {
        npc.MsgCount = 0
    }
    npc.MsgCount++
    return npc.Msgs[npc.MsgCount - 1]
}
