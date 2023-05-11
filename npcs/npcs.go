package npcs

import (
    //"strconv"

    "github.com/jsnider-mtu/quailgame/player"
)

type NPC struct {
    Msgs [][]string
    MsgCount int
    speed int
    offset int
    Direction string
    Stopped bool
    PC *player.Player
}

func NewNPC(direction string, msgs [][]string, speed, offset int, pc *player.Player) *NPC {
    return &NPC{Msgs: msgs, MsgCount: 0, speed: speed, offset: offset, Direction: direction, Stopped: true, PC: pc}
}

func (npc *NPC) GetName() string {
    return npc.PC.Name
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

//func (npc *NPC) SaveHP() string {
//    return npc.GetName() + "=" + strconv.Itoa(npc.PC.Stats.HP) + ";"
//}
