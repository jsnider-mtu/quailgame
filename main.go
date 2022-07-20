package main

import (
    "bytes"
//    "fmt"
    "image"
    "image/color"
    _ "image/png"
    "log"
//    "os"

    "github.com/jsnider-mtu/projectx/assets"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
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
    l *Level
    p *Player
)

type Level struct {
    Max [2]float64
    Pos [2]float64
    Boxes [][4]float64
}

type Player struct {
    pos [2]float64
}

func (p *Player) TryUpdatePos(l *Level, vert bool, dist float64) bool {
    if vert {
        // up
        if dist < 0 {
            if p.pos[1] + dist > 0 && p.pos[1] + dist < l.Max[1] {
                for _, a := range l.Boxes {
                    if p.pos[0] + 48 > a[0] && p.pos[1] + dist > a[1] && p.pos[0] < a[2] && p.pos[1] + dist < a[3] {
                        return false
                    }
                }
                p.pos[1] += dist
                l.Pos[1] -= dist
                return true
            }
            return false
        } else {
            // down
            if p.pos[1] + dist > 0 && p.pos[1] + dist < l.Max[1] {
                for _, a := range l.Boxes {
                    if p.pos[0] + 48 > a[0] && p.pos[1] + 48 + dist > a[1] && p.pos[0] < a[2] && p.pos[1] + 48 + dist < a[3] {
                        return false
                    }
                }
                p.pos[1] += dist
                l.Pos[1] -= dist
                return true
            }
            return false
        }
    } else {
        // left
        if dist < 0 {
            if p.pos[0] + dist > 0 && p.pos[0] + dist < l.Max[0] {
                for _, b := range l.Boxes {
                    if p.pos[0] + dist > b[0] && p.pos[1] + 48 > b[1] && p.pos[0] + dist < b[2] && p.pos[1] < b[3] {
                        return false
                    }
                }
                p.pos[0] += dist
                l.Pos[0] -= dist
                return true
            }
            return false
        } else {
            // right
            if p.pos[0] + dist > 0 && p.pos[0] + dist < l.Max[0] {
                for _, b := range l.Boxes {
                    if p.pos[0] + 48 + dist > b[0] && p.pos[1] + 48 > b[1] && p.pos[0] + 48 + dist < b[2] && p.pos[1] < b[3] {
                        return false
                    }
                }
                p.pos[0] += dist
                l.Pos[0] -= dist
                return true
            }
            return false
        }
    }
}

type Game struct {}

func (g *Game) Update() error {
    if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
        stopped = false
        up = true
        down = false
        left = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyW) % 5 == 0 {
            p.TryUpdatePos(l, true, -float64(48))
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
        stopped = false
        left = true
        up = false
        down = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyA) % 5 == 0 {
            p.TryUpdatePos(l, false, -float64(48))
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
        stopped = false
        right = true
        left = false
        up = false
        down = false
        if inpututil.KeyPressDuration(ebiten.KeyD) % 5 == 0 {
            p.TryUpdatePos(l, false, float64(48))
        }
        count++
    }
    if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
        stopped = false
        down = true
        up = false
        left = false
        right = false
        if inpututil.KeyPressDuration(ebiten.KeyS) % 5 == 0 {
            p.TryUpdatePos(l, true, float64(48))
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
        bgm.Translate(float64(w / 2) + box[0], float64(h / 2) + box[1])
        bi := ebiten.NewImage(int(box[2] - box[0]), int(box[3] - box[1]))
        bi.Fill(color.Black)
        levelImage.DrawImage(bi, &ebiten.DrawImageOptions{
            GeoM: bgm})
    }
    lgm := ebiten.GeoM{}
    lgm.Translate(l.Pos[0], l.Pos[1])
    screen.DrawImage(levelImage, &ebiten.DrawImageOptions{GeoM: lgm})
    gm := ebiten.GeoM{}
    gm.Scale(0.75, 0.75) // 48x48
    gm.Translate(float64(w / 2), float64(h / 2))
    switch {
    case up:
        if stopped {
            screen.DrawImage(
                pcImage.SubImage(
                    image.Rect(
                        pcUpOffsetX, pcUpOffsetY, pcUpOffsetX + 64, pcUpOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 10) % 4
            sx, sy := pcUpOffsetX + (i * 64), pcUpOffsetY
            screen.DrawImage(
                pcImage.SubImage(
                    image.Rect(
                        sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        }
    case left:
        if stopped {
            screen.DrawImage(
                pcImage.SubImage(
                    image.Rect(
                        pcLeftOffsetX, pcLeftOffsetY, pcLeftOffsetX + 64, pcLeftOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 10) % 4
            sx, sy := pcLeftOffsetX + (i * 64), pcLeftOffsetY
            screen.DrawImage(
                pcImage.SubImage(
                    image.Rect(
                        sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        }
    case right:
        if stopped {
            screen.DrawImage(
                pcImage.SubImage(
                    image.Rect(
                        pcRightOffsetX, pcRightOffsetY, pcRightOffsetX + 64, pcRightOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 10) % 4
            sx, sy := pcRightOffsetX + (i * 64), pcRightOffsetY
            screen.DrawImage(
                pcImage.SubImage(
                    image.Rect(
                        sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        }
    case down:
        if stopped {
            screen.DrawImage(
                pcImage.SubImage(
                    image.Rect(
                        pcDownOffsetX, pcDownOffsetY, pcDownOffsetX + 64, pcDownOffsetY + 64)).(*ebiten.Image),
                        &ebiten.DrawImageOptions{
                            GeoM: gm})
        } else {
            i := (count / 10) % 4
            sx, sy := pcDownOffsetX + (i * 64), pcDownOffsetY
            screen.DrawImage(
                pcImage.SubImage(
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
    pcimage, _, err := image.Decode(bytes.NewReader(assets.PC_png))
    if err != nil {
        log.Fatal(err)
    }
    pcImage = ebiten.NewImageFromImage(pcimage)

    levelimage, _, err := image.Decode(bytes.NewReader(platformer.Background_png))
    if err != nil {
        log.Fatal(err)
    }
    levelImage = ebiten.NewImageFromImage(levelimage)

    lw, lh := levelImage.Size()
    //l = &Level{Max: [2]float64{float64(lw - 768), float64(lh - 576)}, Pos: [2]float64{0, 0}, Boxes: [][4]float64{{0, 0, 48, 48}}}
    l = &Level{Max: [2]float64{float64(lw - 768), float64(lh - 576)}, Pos: [2]float64{0, 0}, Boxes: [][4]float64{{576, 336, 672, 432}}}

    p = &Player{pos: l.Pos}
    //p = &Player{pos: [2]float64{float64(0), float64(0)}}
}

func main() {
    ebiten.SetWindowSize(768, 576)
    ebiten.SetWindowTitle("CHANGEME")

    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
