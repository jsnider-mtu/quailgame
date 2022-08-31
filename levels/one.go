package levels

import (
    "bytes"
    "fmt"
    "image"
    _ "image/jpeg"
    _ "image/png"
    "log"
    "math/rand"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/levels/lvlimages"
    "github.com/jsnider-mtu/quailgame/npcs"
    "github.com/jsnider-mtu/quailgame/npcs/npcimages"
    "github.com/jsnider-mtu/quailgame/player"
)

func LvlOne(entrance int) *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.One_JPEG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)
    npcgirlimage, _, err := image.Decode(bytes.NewReader(npcimages.NPCGirl_PNG))
    if err != nil {
        log.Fatal(err)
    }
    npcGirlImage := ebiten.NewImageFromImage(npcgirlimage)
    wizardimage, _, err := image.Decode(bytes.NewReader(npcimages.Wizard_PNG))
    if err != nil {
        log.Fatal(err)
    }
    wizardImage := ebiten.NewImageFromImage(wizardimage)

    lvldoors := []*Door{
        &Door{Coords: [2]int{0, 0}, NewLvl: []interface{}{"Two", 1}},
        &Door{Coords: [2]int{336, 504}, NewLvl: []interface{}{"Two", 2}}}

    NPCs := []*npcs.NPC{&npcs.NPC{
        Name: "Jane Doe", Msgs: [][]string{
            {"Hello there,", "ObiWan Kenobi."},
            {"Seen my dog?", "I swear he was just here...", "Please help me look for him."}},
        MsgCount: 0, Speed: 240, Offset: rand.Intn(60) + 60, Direction: "down", Stopped: true, PC: &player.Player{
            Pos: [2]int{192, 192}, Image: npcGirlImage}},
        &npcs.NPC{Name: "Wizard", Msgs: [][]string{
            {"I'm a wizard, Harry!"},
            {"The great and terrible Lord Adrian", "has invaded the peaceful Quail Kingdom,",
             "stoking a rebellion from within.", "", "Your quest is simple,",
             "quell the rebellion and defeat", "Lord Adrian!"}},
        MsgCount: 0, Speed: 0, Offset: rand.Intn(60) + 60, Direction: "down", Stopped: true, PC: &player.Player{
            Pos: [2]int{288, 288}, Image: wizardImage}}}

    var pos [2]int

    switch entrance {
    case 0:
        pos = [2]int{-48, -144}
    case 1:
        pos = [2]int{-336, -504}
    default:
        log.Fatal(fmt.Sprintf("Entrace %d does not exist", entrance))
    }

    return &Level{
        Name: "One", Cutscene: -1, Max: [2]int{720, 528}, Pos: pos, Boxes: [][4]int{
            {48, 48, 96, 96}}, Doors: lvldoors, NPCs: NPCs, Image: lvlImg, Anim: func(a *ebiten.Image, l *Level, b, c, d int) {}}
}