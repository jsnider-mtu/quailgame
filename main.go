package main

import (
    "bytes"
//    "fmt"
    "image"
    "image/color"
    _ "image/png"
    "log"
    "math/rand"
//    "os"

    "github.com/jsnider-mtu/projectx/player"
    "github.com/jsnider-mtu/projectx/player/pcimages"
    "github.com/jsnider-mtu/projectx/levels"
//    "github.com/jsnider-mtu/projectx/npcs"
    "github.com/jsnider-mtu/projectx/utils"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
    start bool = false
    pause bool = false
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
    npcCount int = 0
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
        if inpututil.KeyPressDuration(ebiten.KeyW) % 3 == 0 {
            if utils.TryUpdatePos(true, p, l, true, -24, p) {
                for _, a := range l.Doors {
                    if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                        l = loadlvl(a.NewLvl)
                        p.Pos[0] = -l.Pos[0]
                        p.Pos[1] = -l.Pos[1]
                    }
                }
            }
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
        stopped = false
        left = true
        up = false
        down = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyA) % 3 == 0 {
            if utils.TryUpdatePos(true, p, l, false, -24, p) {
                for _, a := range l.Doors {
                    if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                        l = loadlvl(a.NewLvl)
                        p.Pos[0] = -l.Pos[0]
                        p.Pos[1] = -l.Pos[1]
                    }
                }
            }
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
        stopped = false
        right = true
        left = false
        up = false
        down = false
        if inpututil.KeyPressDuration(ebiten.KeyD) % 3 == 0 {
            if utils.TryUpdatePos(true, p, l, false, 24, p) {
                for _, a := range l.Doors {
                    if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                        l = loadlvl(a.NewLvl)
                        p.Pos[0] = -l.Pos[0]
                        p.Pos[1] = -l.Pos[1]
                    }
                }
            }
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
        stopped = false
        down = true
        up = false
        left = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyS) % 3 == 0 {
            if utils.TryUpdatePos(true, p, l, true, 24, p) {
                for _, a := range l.Doors {
                    if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                        l = loadlvl(a.NewLvl)
                        p.Pos[0] = -l.Pos[0]
                        p.Pos[1] = -l.Pos[1]
                    }
                }
            }
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
        bgm.Translate(float64(box[0]), float64(box[1]))
        bi := ebiten.NewImage(box[2] - box[0], box[3] - box[1])
        bi.Fill(color.White)
        l.Image.DrawImage(bi, &ebiten.DrawImageOptions{
            GeoM: bgm})
    }
    for _, door := range l.Doors {
        dgm := ebiten.GeoM{}
        dgm.Translate(float64(door.Coords[0]), float64(door.Coords[1]))
        l.Image.DrawImage(door.Image, &ebiten.DrawImageOptions{
            GeoM: dgm})
    }
    lgm := ebiten.GeoM{}
    lgm.Translate(float64((w / 2) + l.Pos[0]), float64((h / 2) + l.Pos[1]))
    screen.DrawImage(l.Image, &ebiten.DrawImageOptions{GeoM: lgm})
    if npcCount == 1000 {
        npcCount = 0
    }
    npcCount++
    for _, npc := range l.NPCs {
        ngm := ebiten.GeoM{}
        ngm.Scale(0.75, 0.75) // 48x48
        ngm.Translate(float64((w / 2) + l.Pos[0] + npc.PC.Pos[0]), float64((h / 2) + l.Pos[1] + npc.PC.Pos[1]))
        if npcCount % npc.Speed == 0 {
            switch rand.Intn(4) {
            case 0:
                if utils.TryUpdatePos(false, npc.PC, l, true, 24, p) {
                    npc.Direction = "down"
                }
            case 1:
                if utils.TryUpdatePos(false, npc.PC, l, true, -24, p) {
                    npc.Direction = "up"
                }
            case 2:
                if utils.TryUpdatePos(false, npc.PC, l, false, 24, p) {
                    npc.Direction = "right"
                }
            case 3:
                if utils.TryUpdatePos(false, npc.PC, l, false, -24, p) {
                    npc.Direction = "left"
                }
            }
        }
        switch npc.Direction {
        case "down":
            screen.DrawImage(
                npc.PC.Image.SubImage(
                    image.Rect(
                        pcDownOffsetX, pcDownOffsetY, pcDownOffsetX + 64, pcDownOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: ngm})
        case "up":
            screen.DrawImage(
                npc.PC.Image.SubImage(
                    image.Rect(
                        pcUpOffsetX, pcUpOffsetY, pcUpOffsetX + 64, pcUpOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: ngm})
        case "right":
            screen.DrawImage(
                npc.PC.Image.SubImage(
                    image.Rect(
                        pcRightOffsetX, pcRightOffsetY, pcRightOffsetX + 64, pcRightOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: ngm})
        case "left":
            screen.DrawImage(
                npc.PC.Image.SubImage(
                    image.Rect(
                        pcLeftOffsetX, pcLeftOffsetY, pcLeftOffsetX + 64, pcLeftOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: ngm})
        }
    }
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

func loadlvl(lvl int) *levels.Level {
    switch lvl {
    case 1:
        return levels.LvlOne()
    case 2:
        return levels.LvlTwo()
    }
    return levels.LvlOne()
}

func init() {
    pcimage, _, err := image.Decode(bytes.NewReader(pcimages.PC_png))
    if err != nil {
        log.Fatal(err)
    }
    pcImage = ebiten.NewImageFromImage(pcimage)

    l = levels.LvlOne()
    p = &player.Player{Pos: [2]int{-l.Pos[0], -l.Pos[1]}, Image: pcImage}
}

func main() {
    ebiten.SetWindowSize(768, 576)
    ebiten.SetWindowTitle("CHANGEME")

    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
