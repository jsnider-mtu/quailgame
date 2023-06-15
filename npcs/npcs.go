package npcs

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"

    "github.com/jsnider-mtu/quailgame/classes"
    "github.com/jsnider-mtu/quailgame/npcs/npcimages"
    "github.com/jsnider-mtu/quailgame/player"
//    "github.com/jsnider-mtu/quailgame/quests"
)

type NPC interface {
    GetName() string
    GetSpeed() int
    GetOffset() int
    Dialog() []string
}

func NewNPC(name string) NPC {
    switch name {
    case "janedoe":
        npcgirlimage, _, err := image.Decode(bytes.NewReader(npcimages.NPCGirl_PNG))
        if err != nil {
            log.Fatal(err)
        }
        npcGirlImage := ebiten.NewImageFromImage(npcgirlimage)
        return &JaneDoe{name: name, speed: 240, offset: rand.Intn(60) + 60, direction: "down", stopped: true, msgs: [][]string{
            {"Hello there,", "ObiWan Kenobi."},
            {"Seen my dog?", "I swear he was just here...", "Please help me look for him."}}, msgcount: 0, PC: &player.Player{
                Name: "Jane Doe", Pos: [2]int{192, 192}, Image: npcGirlImage, Class: &classes.Quail{}}}
    default:
        log.Fatal(fmt.Sprintf("%d is not a valid NPC", name))
    }
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
