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

func lvlOne(entrance int, npchps string) *Level {
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
        &Door{coords: [2]int{0, 0}, NewLvl: []interface{}{"Two", 1, npchps}},
        &Door{coords: [2]int{336, 504}, NewLvl: []interface{}{"Two", 2, npchps}}}

    NPCs := []*npcs.NPC{npcs.NewNPC(
        "down", [][]string{
            {"Hello there,", "ObiWan Kenobi."},
            {"Seen my dog?", "I swear he was just here...", "Please help me look for him."}},
        240, rand.Intn(60) + 60, &player.Player{
            Name: "Jane Doe", Pos: [2]int{192, 192}, Image: npcGirlImage, Spells: &player.Spells{}, Stats: &player.Stats{
                AC: 10, Str: 10, StrMod: 0, Dex: 10, DexMod: 0, Con: 10, ConMod: 0, Intel: 10, IntelMod: 0, Wis: 10, WisMod: 0,
                Cha: 10, ChaMod: 0, MaxHP: 6, HP: 6, Size: 1}}),
        npcs.NewNPC("down", [][]string{
            {"I'm a wizard, Harry!"},
            {"The great and terrible Lord Adrian", "has invaded the peaceful Quail Kingdom,",
             "stoking a rebellion from within.", "", "Your quest is simple,",
             "quell the rebellion and defeat", "Lord Adrian!"}},
        0, rand.Intn(60) + 60, &player.Player{
            Name: "Wizard", Pos: [2]int{288, 288}, Image: wizardImage, Spells: &player.Spells{}, Stats: &player.Stats{
                AC: 10, Str: 10, StrMod: 0, Dex: 10, DexMod: 0, Con: 10, ConMod: 0, Intel: 10, IntelMod: 0, Wis: 10, WisMod: 0,
                Cha: 10, ChaMod: 0, MaxHP: 6, HP: 6, Size: 1}})}

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
        name: "One", Cutscene: -1, max: [2]int{720, 528}, Pos: pos, Boxes: [][4]int{
            {48, 48, 96, 96}}, Doors: lvldoors, NPCs: NPCs, grasses: [][2]int{
            {336, 384}, {336, 432}}, Image: lvlImg,
        Anim: func(screen *ebiten.Image, l *Level, count, w, h int) {
            l.GrassAnim(screen, w, h)}}
}
