package main

import (
    "bytes"
//    "fmt"
    "image"
    "image/color"
    _ "image/png"
    "log"
//    "os"

    "github.com/jsnider-mtu/projectx/player"
    "github.com/jsnider-mtu/projectx/player/pcimages"
    "github.com/jsnider-mtu/projectx/levels"
    "github.com/jsnider-mtu/projectx/levels/lvlone"
//    "github.com/jsnider-mtu/projectx/npcs"
    "github.com/jsnider-mtu/projectx/utils"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
//    "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
)

var (
    start bool = false
    pause bool = false
    levelImage *ebiten.Image
    pcImage *ebiten.Image
    pcDownOffsetX int = 0
    pcDownOffsetY int = 0
    pcLeftOffsetX int = 0
    pcLeftOffsetY int = 64
    pcRightOffsetX int = 0
    pcRightOffsetY int = 128
    pcUpOffsetX int = 0
    pcUpOffsetY int = 192
    down bool = true
    up bool = false
    left bool = false
    right bool = false
    stopped bool = true
    count int = 0
    lastCount int = 0
    lw, lh int
    l *levels.Level
    p *player.Player
)

type Game struct {}

func (g *Game) Update() error {
    if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
        stopped = false
        up = true
        down = false
        left = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyW) % 2 == 0 {
            utils.TryUpdatePos(true, p, l, true, -24)
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
        stopped = false
        left = true
        up = false
        down = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyA) % 2 == 0 {
            utils.TryUpdatePos(true, p, l, false, -24)
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
        stopped = false
        right = true
        left = false
        up = false
        down = false
        if inpututil.KeyPressDuration(ebiten.KeyD) % 2 == 0 {
            utils.TryUpdatePos(true, p, l, false, 24)
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
        stopped = false
        down = true
        up = false
        left = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyS) % 2 == 0 {
            utils.TryUpdatePos(true, p, l, true, 24)
        }
        count++
    }
    if count == lastCount {
        stopped = true
        count = 0
        lastCount = 0
    } else {
        lastCount = count
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    w, h := ebiten.WindowSize()
    if ebiten.IsFullscreen() {
        w, h = ebiten.ScreenSizeInFullscreen()
    }
    for _, box := range l.Boxes {
        bgm := ebiten.GeoM{}
        bgm.Translate(float64((w / 2) + box[0]), float64((h / 2) + box[1]))
        bi := ebiten.NewImage(box[2] - box[0], box[3] - box[1])
        bi.Fill(color.Black)
        l.Image.DrawImage(bi, &ebiten.DrawImageOptions{
            GeoM: bgm})
    }
    for _, door := range l.Doors {
        dgm := ebiten.GeoM{}
        dgm.Translate(float64((w / 2) + door.Coords[0]), float64((h / 2) + door.Coords[1]))
        l.Image.DrawImage(door.Image, &ebiten.DrawImageOptions{
            GeoM: dgm})
    }
    lgm := ebiten.GeoM{}
    lgm.Translate(float64(l.Pos[0]), float64(l.Pos[1]))
    screen.DrawImage(l.Image, &ebiten.DrawImageOptions{GeoM: lgm})
    gm := ebiten.GeoM{}
    gm.Scale(0.75, 0.75) // 48x48
    gm.Translate(float64(w / 2), float64(h / 2))
    switch {
    case up:
        if stopped {
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        pcUpOffsetX, pcUpOffsetY, pcUpOffsetX + 64, pcUpOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 5) % 4
            sx, sy := pcUpOffsetX + (i * 64), pcUpOffsetY
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        }
    case left:
        if stopped {
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        pcLeftOffsetX, pcLeftOffsetY, pcLeftOffsetX + 64, pcLeftOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 5) % 4
            sx, sy := pcLeftOffsetX + (i * 64), pcLeftOffsetY
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        }
    case right:
        if stopped {
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        pcRightOffsetX, pcRightOffsetY, pcRightOffsetX + 64, pcRightOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 5) % 4
            sx, sy := pcRightOffsetX + (i * 64), pcRightOffsetY
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        }
    case down:
        if stopped {
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        pcDownOffsetX, pcDownOffsetY, pcDownOffsetX + 64, pcDownOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 5) % 4
            sx, sy := pcDownOffsetX + (i * 64), pcDownOffsetY
            screen.DrawImage(
                p.Image.SubImage(
                    image.Rect(
                        sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        }
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int)  {
    return outsideWidth, outsideHeight
}

func init() {
    pcimage, _, err := image.Decode(bytes.NewReader(pcimages.PC_png))
    if err != nil {
        log.Fatal(err)
    }
    pcImage = ebiten.NewImageFromImage(pcimage)

//    levelimage, _, err := image.Decode(bytes.NewReader(platformer.Background_png))
//    if err != nil {
//        log.Fatal(err)
//    }
//    levelImage = ebiten.NewImageFromImage(levelimage)

//    lw, lh := levelImage.Size()
//    lvldoors := []*levels.Door{&levels.Door{Coords: [4]int{96, 96, 192, 144}, Direction: "up", Image: ebiten.NewImage(192-96, 144-96)}}
//    for _, ld := range lvldoors {
//        ld.Image.Fill(color.Black)
//    }
    l = lvlone.Setup()
//    l = &levels.Level{Max: [2]int{lw - 768, lh - 576}, Pos: [2]int{0, 0}, Boxes: [][4]int{{576, 336, 672, 432}}, Doors: lvldoors, Image: levelImage}
//    player.P.StartLoc(l, l.Pos)
    p = &player.Player{Pos: l.Pos, Image: pcImage}
}

func main() {
    ebiten.SetWindowSize(768, 576)
    ebiten.SetWindowTitle("CHANGEME")

    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
