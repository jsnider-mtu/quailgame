package levels

import (
    "bytes"
    "fmt"
    "image"
//    _ "image/jpeg"
    _ "image/png"
    "log"
    "math/rand"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/quailgame/levels/lvlimages"
    "github.com/jsnider-mtu/quailgame/npcs"
    "github.com/jsnider-mtu/quailgame/npcs/npcimages"
    "github.com/jsnider-mtu/quailgame/player"
)

func lvlTwo(entrance int) *Level {
    npcgirlimage, _, err := image.Decode(bytes.NewReader(npcimages.NPCGirl_PNG))
    if err != nil {
        log.Fatal(err)
    }
    npcGirlImage := ebiten.NewImageFromImage(npcgirlimage)
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.Two_PNG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)

    lvldoors := []*Door{
        &Door{coords: [2]int{48, 96}, NewLvl: []interface{}{"One", 1}},
        &Door{coords: [2]int{144, 0}, NewLvl: []interface{}{"VerticalWall", 1}},
        &Door{coords: [2]int{240, 96}, NewLvl: []interface{}{"One", 1}}}

    NPCs := []*npcs.NPC{npcs.NewNPC(
        "down", [][]string{
            {"Hello there,", "ObiWan Kenobi."},
            {"Seen my dog?", "I swear he was just here...", "Please help me look for him."}},
        240, rand.Intn(60) + 60, &player.Player{
            Name: "Jane Doe", Pos: [2]int{144, 48}, Image: npcGirlImage, Spells: &player.Spells{}, Stats: &player.Stats{
                AC: 10, Str: 10, StrMod: 0, Dex: 10, DexMod: 0, Con: 10, ConMod: 0, Intel: 10, IntelMod: 0, Wis: 10, WisMod: 0,
                Cha: 10, ChaMod: 0, MaxHP: 6, HP: 6, Size: 1}})}

    var pos [2]int

    switch entrance {
    case 0:
        pos = [2]int{-48, -144}
    case 1:
        pos = [2]int{-96, -192}
    case 2:
        pos = [2]int{-144, 0}
    default:
        log.Fatal(fmt.Sprintf("Entrace %d does not exist", entrance))
    }

    return &Level{
        name: "Two", Cutscene: 1, max: [2]int{312, 216}, Pos: pos, Boxes: [][4]int{
            {0, 0, 144, 96},
            {0, 96, 48, 240},
            {96, 96, 144, 192},
            {192, 0, 336, 96},
            {192, 96, 240, 144},
            {288, 96, 336, 240},
            {192, 192, 288, 240}}, Doors: lvldoors, NPCs: NPCs, Image: lvlImg, Anim: func(a *ebiten.Image, l *Level, b, c, d int) {}}
}
