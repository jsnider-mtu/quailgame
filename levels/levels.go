package levels

import (
    "bytes"
    "image"
    _ "image/jpeg"
    _ "image/png"
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/projectx/levels/lvlimages"
    "github.com/jsnider-mtu/projectx/npcs"
    "github.com/jsnider-mtu/projectx/player"
    "github.com/jsnider-mtu/projectx/player/pcimages"
)

type Level struct {
    Max [2]int
    Pos [2]int
    Boxes [][4]int
    Doors []*Door
    NPCs []*npcs.NPC
    Image *ebiten.Image
}

type Door struct {
    Coords [2]int
    NewLvl int
}

func LvlOne() *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.One_JPEG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)
    pcimage, _, err := image.Decode(bytes.NewReader(pcimages.PC_png))
    if err != nil {
        log.Fatal(err)
    }
    pcImage := ebiten.NewImageFromImage(pcimage)

    lvldoors := []*Door{&Door{Coords: [2]int{0, 0}, NewLvl: 2}}

    NPCs := []*npcs.NPC{&npcs.NPC{Name: "FirstNPC", Msgs: [][]string{{"Hello there,", "ObiWan Kenobi."}, {"Seen my dog?", "I swear he was just here...", "Please help me look for him."}}, MsgCount: 0, Speed: 200, Direction: "down", PC: &player.Player{Pos: [2]int{192, 192}, Image: pcImage}}}

    return &Level{Max: [2]int{720, 528}, Pos: [2]int{-48, -144}, Boxes: [][4]int{{48, 48, 96, 96}}, Doors: lvldoors, NPCs: NPCs, Image: lvlImg}
}

func LvlTwo() *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.Two_PNG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)

    lvldoors := []*Door{&Door{Coords: [2]int{48, 96}, NewLvl: 1}}

    return &Level{Max: [2]int{120, 216}, Pos: [2]int{-48, -144}, Boxes: [][4]int{{0, 0, 144, 96}, {0, 96, 48, 240}, {96, 96, 144, 192}}, Doors: lvldoors, NPCs: []*npcs.NPC{}, Image: lvlImg}
    //return &Level{Max: [2]int{144, 240}, Pos: [2]int{-48, -144}, Boxes: [][4]int{{0, 0, 144, 96}, {0, 96, 48, 240}, {96, 96, 144, 192}}, Doors: lvldoors, NPCs: []*npcs.NPC{}, Image: lvlImg}
}
