package npcs

import (
//    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/player"
//    "github.com/jsnider-mtu/quailgame/levels"
//    "github.com/jsnider-mtu/quailgame/utils"
)

type NPC struct {
    name string
    Msgs [][]string
    MsgCount int
    speed int
    offset int
    Direction string
    Stopped bool
    PC *player.Player
}

func NewNPC(name, direction string, msgs [][]string, speed, offset int, pc *player.Player) *NPC {
    return &NPC{name: name, Msgs: msgs, MsgCount: 0, speed: speed, offset: offset, Direction: direction, Stopped: true, PC: pc}
}

func (npc *NPC) GetName() string {
    return npc.name
}

func (npc *NPC) GetSpeed() int {
    return npc.speed
}

func (npc *NPC) GetOffset() int {
    return npc.offset
}

func (npc *NPC) Dialog() []string {
    if npc.MsgCount == len(npc.Msgs) {
        npc.MsgCount = 0
    }
    npc.MsgCount++
    return npc.Msgs[npc.MsgCount - 1]
}
