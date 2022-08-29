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
    //"github.com/jsnider-mtu/quailgame/player/pcimages"
)

type Door struct {
    Coords [2]int
    NewLvl [2]int
}

type Level struct {
    Name string
    Cutscene int
    Max [2]int
    Pos [2]int
    Boxes [][4]int
    Doors []*Door
    NPCs []*npcs.NPC
    Image *ebiten.Image
    Anim func(*ebiten.Image, *Level, int, int, int)
}

func LoadLvl(name string, x, y int) *Level {
    switch name {
    case "One":
        l := LvlOne(0)
        l.Pos = [2]int{x, y}
        return l
    case "Two":
        l := LvlTwo(0)
        l.Pos = [2]int{x, y}
        return l
    case "VerticalWall":
        l := VerticalWallLvl(0)
        l.Pos = [2]int{x, y}
        return l
    default:
        log.Fatal(fmt.Sprintf("Level %s does not exist"))
    }
    return LvlOne(0)
}

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
        &Door{Coords: [2]int{0, 0}, NewLvl: [2]int{2, 1}},
        &Door{Coords: [2]int{336, 504}, NewLvl: [2]int{2, 2}}}

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

func LvlTwo(entrance int) *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.Two_PNG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)

    lvldoors := []*Door{
        &Door{Coords: [2]int{48, 96}, NewLvl: [2]int{1, 1}},
        &Door{Coords: [2]int{144, 0}, NewLvl: [2]int{3, 1}},
        &Door{Coords: [2]int{240, 96}, NewLvl: [2]int{1, 1}}}

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
        Name: "Two", Cutscene: 1, Max: [2]int{312, 216}, Pos: pos, Boxes: [][4]int{
            {0, 0, 144, 96},
            {0, 96, 48, 240},
            {96, 96, 144, 192},
            {192, 0, 336, 96},
            {192, 96, 240, 144},
            {288, 96, 336, 240},
            {192, 192, 288, 240}}, Doors: lvldoors, NPCs: []*npcs.NPC{}, Image: lvlImg, Anim: func(a *ebiten.Image, l *Level, b, c, d int) {}}
}

func VerticalWallLvl(entrance int) *Level {
    lvlimg, _, err := image.Decode(bytes.NewReader(lvlimages.VerticalWallLvlOne_PNG))
    if err != nil {
        log.Fatal(err)
    }
    lvlImg := ebiten.NewImageFromImage(lvlimg)
    cloudimg, _, err := image.Decode(bytes.NewReader(lvlimages.Clouds_PNG))
    if err != nil {
        log.Fatal(err)
    }
    cloudImage := ebiten.NewImageFromImage(cloudimg)

    lvldoors := []*Door{}

    for x := 0; x < 15; x++ {
        lvldoors = append(lvldoors, &Door{Coords: [2]int{432 + (24 * x), 288}, NewLvl: [2]int{2, 1}})
    }

    for x := 0; x < 15; x++ {
        lvldoors = append(lvldoors, &Door{Coords: [2]int{432 + (24 * x), 5808}, NewLvl: [2]int{2, 2}})
    }

    var pos [2]int

    switch entrance {
    case 0:
        pos = [2]int{-600, -336}
    case 1:
        pos = [2]int{-600, -5760}
    default:
        log.Fatal(fmt.Sprintf("Entrace %d does not exist", entrance))
    }

    return &Level{
        Name: "VerticalWall", Cutscene: -1, Max:[2]int{816, 5856}, Pos: pos, Boxes: [][4]int{
            {0, 0, 432, 6144},
            {816, 0, 1248, 6144}}, Doors: lvldoors, NPCs: []*npcs.NPC{}, Image: lvlImg,
        Anim: func(screen *ebiten.Image, l *Level, count, w, h int) {
            iw, _ := l.Image.Size()
            gm := ebiten.GeoM{}
            gm.Translate(float64(((iw - 96) + l.Pos[0]) - count), float64(l.Pos[1]))
            screen.DrawImage(cloudImage, &ebiten.DrawImageOptions{GeoM: gm})}}
}
