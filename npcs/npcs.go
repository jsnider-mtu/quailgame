package npcs

import (
    "fmt"
    "log"

//    "github.com/hajimehoshi/ebiten/v2"

//    "github.com/jsnider-mtu/quailgame/classes"
//    "github.com/jsnider-mtu/quailgame/npcs/npcimages"
    "github.com/jsnider-mtu/quailgame/player"
//    "github.com/jsnider-mtu/quailgame/quests"
)

type NPC interface {
    GetPC() *player.Player
    GetName() string
    GetSpeed() int
    GetOffset() int
    GetDirection() string
    GetStopped() bool
    Dialog() []string
    Direction(string)
    Stopped(bool)
}

func NewNPC(name, direction string, speed, offset int, stopped bool, msgs [][]string, pc *player.Player) NPC {
    switch name {
    case "janedoe":
        return &JaneDoe{name: name, speed: speed, offset: offset, direction: direction, stopped: stopped, msgs: msgs, msgcount: 0, PC: pc}
    case "wizard":
        return &Wizard{name: name, speed: speed, offset: offset, direction: direction, stopped: stopped, msgs: msgs, msgcount: 0, PC: pc}
    default:
        log.Fatal(fmt.Sprintf("%d is not a valid NPC", name))
    }
    return nil
}

//type NPC struct {
//    Msgs [][]string
//    MsgCount int
//    speed int
//    offset int
//    Direction string
//    Stopped bool
//    PC *player.Player
//}
//
//func NewNPC(direction string, msgs [][]string, speed, offset int, pc *player.Player) *NPC {
//    return &NPC{Msgs: msgs, MsgCount: 0, speed: speed, offset: offset, Direction: direction, Stopped: true, PC: pc}
//}
//
//func (npc *NPC) GetName() string {
//    return npc.PC.Name
//}
//
//func (npc *NPC) GetSpeed() int {
//    return npc.speed
//}
//
//func (npc *NPC) GetOffset() int {
//    return npc.offset
//}
//
//func (npc *NPC) Dialog() []string {
//    if npc.MsgCount == len(npc.Msgs) {
//        npc.MsgCount = 0
//    }
//    npc.MsgCount++
//    return npc.Msgs[npc.MsgCount - 1]
//}
