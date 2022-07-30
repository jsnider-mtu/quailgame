package main

import (
    "bytes"
    "database/sql"
    "fmt"
    "image"
    "image/color"
    _ "image/png"
    "log"
    "math/rand"
    "os"

    "golang.org/x/image/font"
    "golang.org/x/image/font/gofont/gomonobold"

    "github.com/golang/freetype/truetype"

    "github.com/jsnider-mtu/projectx/assets"
    "github.com/jsnider-mtu/projectx/player"
    "github.com/jsnider-mtu/projectx/player/pcimages"
    "github.com/jsnider-mtu/projectx/levels"
    "github.com/jsnider-mtu/projectx/utils"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"

    _ "github.com/mattn/go-sqlite3"
)

var (
    err error
    start bool = false // not yet implemented
    pause bool = false // not yet implemented
    pausesel int = 0
    save bool = false
    load bool = false
    cont bool = false
    name string = "tempname"
    downArrowImage *ebiten.Image
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
    dialogopen bool = false
    dialogstrs []string
    npcname string
    l *levels.Level
    p *player.Player
    fon *truetype.Font
    fo font.Face
    s int = 0
    lvlchange bool = false
    newlvl [2]int
    f int = 0
    fadeImage *ebiten.Image
    dab int = 0
    dialogCount int = 0
)

type Game struct {}

func (g *Game) Update() error {
    if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
        pause = !pause
    }
    if pause {
        if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
            if pausesel > 0 {
                pausesel--
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
            if pausesel < 2 {
                pausesel++
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
            switch pausesel {
            case 0:
                save = true
                pause = false
            case 1:
                load = true
                pause = false
            case 2:
                os.Exit(0)
            }
        }
    } else {
        if save {
            homeDir, err := os.UserHomeDir()
            if err != nil {
                log.Fatal(err)
            }
            db, err := sql.Open("sqlite3", homeDir + "/quailsaves.db")
            if err != nil {
                log.Fatal(err)
            }
            defer db.Close()
            createStmt := `
            create table if not exists saves (name text not null primary key, level text not null, x int not null, y int not null);
            `
            _, err = db.Exec(createStmt)
            if err != nil {
                log.Fatal(fmt.Sprintf("%q: %s\n", err, createStmt))
            }
            saveStmt := `
            insert or replace into saves(name, level, x, y) values(?, ?, ?, ?);
            `
            _, err = db.Exec(saveStmt, name, l.Name, l.Pos[0], l.Pos[1])
            if err != nil {
                log.Fatal(fmt.Sprintf("%q: %s\n", err, saveStmt))
            }
            db.Close()
            save = false
        }
        if load {
            homeDir, err := os.UserHomeDir()
            if err != nil {
                log.Fatal(err)
            }
            db, err := sql.Open("sqlite3", homeDir + "/quailsaves.db")
            if err != nil {
                log.Fatal(err)
            }
            defer db.Close()
            rows, err := db.Query("select * from saves where name = ?", name)
            if err != nil {
                log.Fatal(err)
            }
            defer rows.Close()
            var savename string
            var levelname string
            var x, y int
            for rows.Next() {
                err = rows.Scan(&savename, &levelname, &x, &y)
            }
            err = rows.Err()
            if err != nil {
                log.Fatal(err)
            }
            l = levels.LoadLvl(levelname, x, y)
            p.Pos[0] = -l.Pos[0]
            p.Pos[1] = -l.Pos[1]
            load = false
        }
        if npcCount == 6000 {
            npcCount = 0
        }
        if !dialogopen {
            npcCount++
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyF) {
            if dialogopen {
                s += 2
                if s >= len(dialogstrs) {
                    dialogopen = false
                    s = 0
                }
                return nil
            }
            switch {
            case up:
                for _, npc := range l.NPCs {
                    if npc.PC.Pos[0] >= p.Pos[0] - 24 && npc.PC.Pos[0] <= p.Pos[0] + 24 && npc.PC.Pos[1] + 24 == p.Pos[1] {
                        if !dialogopen {
                            npcname = npc.Name
                            dialogstrs = npc.Dialog()
                            dialogopen = true
                        }
                    }
                }
            case down:
                for _, npc := range l.NPCs {
                    if npc.PC.Pos[0] >= p.Pos[0] - 24 && npc.PC.Pos[0] <= p.Pos[0] + 24 && npc.PC.Pos[1] - 48 == p.Pos[1] {
                        if !dialogopen {
                            npcname = npc.Name
                            dialogstrs = npc.Dialog()
                            dialogopen = true
                        }
                    }
                }
            case left:
                for _, npc := range l.NPCs {
                    if npc.PC.Pos[1] >= p.Pos[1] - 24 && npc.PC.Pos[1] <= p.Pos[1] + 24 && npc.PC.Pos[0] + 24 == p.Pos[0] {
                        if !dialogopen {
                            npcname = npc.Name
                            dialogstrs = npc.Dialog()
                            dialogopen = true
                        }
                    }
                }
            case right:
                for _, npc := range l.NPCs {
                    if npc.PC.Pos[1] >= p.Pos[1] - 24 && npc.PC.Pos[1] <= p.Pos[1] + 24 && npc.PC.Pos[0] - 24 == p.Pos[0] {
                        if !dialogopen {
                            npcname = npc.Name
                            dialogstrs = npc.Dialog()
                            dialogopen = true
                        }
                    }
                }
            }
        }
        if !dialogopen && !lvlchange {
            for _, npc := range l.NPCs {
                if (npcCount + npc.Offset) % npc.Speed == 0 {
                    npc.Stopped = false
                    switch rand.Intn(4) {
                    case 0:
                        npc.Direction = "down"
                        utils.TryUpdatePos(false, npc.PC, l, true, 24, p)
                    case 1:
                        npc.Direction = "up"
                        utils.TryUpdatePos(false, npc.PC, l, true, -24, p)
                    case 2:
                        npc.Direction = "right"
                        utils.TryUpdatePos(false, npc.PC, l, false, 24, p)
                    case 3:
                        npc.Direction = "left"
                        utils.TryUpdatePos(false, npc.PC, l, false, -24, p)
                    }
                } else if !npc.Stopped && (npcCount + npc.Offset - 4) % npc.Speed == 0 {
                    npc.Stopped = true
                }
            }
            if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
                stopped = false
                up = true
                down = false
                left = false
                right = false
                if inpututil.KeyPressDuration(ebiten.KeyW) % 4 == 0 {
                    if utils.TryUpdatePos(true, p, l, true, -24, p) {
                        for _, a := range l.Doors {
                            if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                                newlvl = a.NewLvl
                                lvlchange = true
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
                if inpututil.KeyPressDuration(ebiten.KeyA) % 4 == 0 {
                    if utils.TryUpdatePos(true, p, l, false, -24, p) {
                        for _, a := range l.Doors {
                            if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                                newlvl = a.NewLvl
                                lvlchange = true
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
                if inpututil.KeyPressDuration(ebiten.KeyD) % 4 == 0 {
                    if utils.TryUpdatePos(true, p, l, false, 24, p) {
                        for _, a := range l.Doors {
                            if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                                newlvl = a.NewLvl
                                lvlchange = true
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
                if inpututil.KeyPressDuration(ebiten.KeyS) % 4 == 0 {
                    if utils.TryUpdatePos(true, p, l, true, 24, p) {
                        for _, a := range l.Doors {
                            if p.Pos[0] == a.Coords[0] && p.Pos[1] == a.Coords[1] {
                                newlvl = a.NewLvl
                                lvlchange = true
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
        }
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    w, h := ebiten.WindowSize()
    if ebiten.IsFullscreen() {
        w, h = ebiten.ScreenSizeInFullscreen()
    }
    lgm := ebiten.GeoM{}
    lgm.Translate(float64((w / 2) + l.Pos[0]), float64((h / 2) + l.Pos[1]))
    screen.DrawImage(l.Image, &ebiten.DrawImageOptions{GeoM: lgm})
    mcdrawn := false
    for _, npc := range l.NPCs {
        if npc.PC.Pos[0] == p.Pos[0] && npc.PC.Pos[1] == p.Pos[1] + 24 {
            drawmc(screen, w, h)
            mcdrawn = true
        }
        ngm := ebiten.GeoM{}
        ngm.Scale(0.75, 0.75) // 48x48
        ngm.Translate(float64((w / 2) + l.Pos[0] + npc.PC.Pos[0]), float64((h / 2) + l.Pos[1] + npc.PC.Pos[1]))
        switch npc.Direction {
        case "down":
            if !npc.Stopped {
                sx, sy := pcDownOffsetX + 64, pcDownOffsetY
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            } else {
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            pcDownOffsetX, pcDownOffsetY, pcDownOffsetX + 64, pcDownOffsetY + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            }
        case "up":
            if !npc.Stopped {
                sx, sy := pcUpOffsetX + 64, pcUpOffsetY
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            } else {
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            pcUpOffsetX, pcUpOffsetY, pcUpOffsetX + 64, pcUpOffsetY + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            }
        case "right":
            if !npc.Stopped {
                sx, sy := pcRightOffsetX + 64, pcRightOffsetY
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            } else {
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            pcRightOffsetX, pcRightOffsetY, pcRightOffsetX + 64, pcRightOffsetY + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            }
        case "left":
            if !npc.Stopped {
                sx, sy := pcLeftOffsetX + 64, pcLeftOffsetY
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            sx, sy, sx + 64, sy + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            } else {
                screen.DrawImage(
                    npc.PC.Image.SubImage(
                        image.Rect(
                            pcLeftOffsetX, pcLeftOffsetY, pcLeftOffsetX + 64, pcLeftOffsetY + 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{
                                GeoM: ngm})
            }
        }
    }
    if !mcdrawn {
        drawmc(screen, w, h)
    }
    if dialogopen {
        if dialogCount == 1000 {
            dialogCount = 0
        }
        dialogCount++
        dialoggm := ebiten.GeoM{}
        dialoggm.Translate(float64(128), float64(468))
        dialogimg := ebiten.NewImage(512, 108)
        dialogimg.Fill(color.Black)
        screen.DrawImage(
            dialogimg, &ebiten.DrawImageOptions{
                GeoM: dialoggm})
        dialoggm2 := ebiten.GeoM{}
        dialoggm2.Translate(float64(132), float64(472))
        dialogimg2 := ebiten.NewImage(504, 100)
        dialogimg2.Fill(color.White)
        screen.DrawImage(
            dialogimg2, &ebiten.DrawImageOptions{
                GeoM: dialoggm2})
        r := text.BoundString(fo, dialogstrs[0])
        hei := r.Max.Y - r.Min.Y
        if s < len(dialogstrs) {
            text.Draw(screen, npcname, fo, 140, 500, color.RGBA{200, 36, 121, 255})
            text.Draw(screen, dialogstrs[s], fo, 140, 516 + hei, color.Black)
            if s + 1 < len(dialogstrs) {
                text.Draw(screen, dialogstrs[s + 1], fo, 140, 524 + (hei * 2), color.Black)
                if s + 2 < len(dialogstrs) {
                    dagm := ebiten.GeoM{}
                    dagm.Scale(0.25, 0.25)
                    dagm.Translate(float64(586), float64(522))
                    if dialogCount % 13 == 0 {
                        dab++
                    }
                    if dab == 3 || dab == 5 {
                        dagm.Translate(float64(0), float64(-4))
                    } else if dab == 8 {
                        dab = 0
                    }
                    screen.DrawImage(
                        downArrowImage, &ebiten.DrawImageOptions{
                            GeoM: dagm})
                }
            }
        }
    }
    if pause {
        r := text.BoundString(fo, "> Save game")
        hei := r.Max.Y - r.Min.Y
        wid := r.Max.X - r.Min.X
        pausegm := ebiten.GeoM{}
        pausegm.Translate(float64((w / 2) - (wid / 2) - 8), float64((h / 2) - (3 * hei / 2) - 16))
        pauseimg := ebiten.NewImage(wid + 20, (hei * 3) + 48)
        pauseimg.Fill(color.Black)
        screen.DrawImage(
            pauseimg, &ebiten.DrawImageOptions{
                GeoM: pausegm})
        pausegm2 := ebiten.GeoM{}
        pausegm2.Translate(float64((w / 2) - (wid / 2) - 4), float64((h / 2) - (3 * hei / 2) - 12))
        pauseimg2 := ebiten.NewImage(wid + 12, (hei * 3) + 40)
        pauseimg2.Fill(color.White)
        screen.DrawImage(
            pauseimg2, &ebiten.DrawImageOptions{
                GeoM: pausegm2})
        switch pausesel {
        case 0:
            text.Draw(screen, "> Save game", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            text.Draw(screen, "  Load game", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            text.Draw(screen, "  Quit game", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
        case 1:
            text.Draw(screen, "  Save game", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            text.Draw(screen, "> Load game", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            text.Draw(screen, "  Quit game", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
        case 2:
            text.Draw(screen, "  Save game", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            text.Draw(screen, "  Load game", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            text.Draw(screen, "> Quit game", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
        }
    }
    if lvlchange {
        if npcCount % 13 == 0 {
            f++
        }
        if f == 0 {
            screen.DrawImage(fadeImage, &ebiten.DrawImageOptions{})
        } else if f == 1 {
            op := &ebiten.DrawImageOptions{}
            op.ColorM.Scale(1.0, 1.0, 1.0, 2.0)
            screen.DrawImage(fadeImage, op)
        } else if f == 2 {
            op := &ebiten.DrawImageOptions{}
            op.ColorM.Scale(1.0, 1.0, 1.0, 3.0)
            screen.DrawImage(fadeImage, op)
        } else if f == 3 {
            op := &ebiten.DrawImageOptions{}
            op.ColorM.Scale(1.0, 1.0, 1.0, 4.0)
            screen.DrawImage(fadeImage, op)
        } else if f == 4 {
            f = 0
            lvlchange = false
            l = loadlvl(newlvl)
            p.Pos[0] = -l.Pos[0]
            p.Pos[1] = -l.Pos[1]
        }
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int)  {
    return outsideWidth, outsideHeight
}

func drawmc(screen *ebiten.Image, w, h int) {
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

func loadlvl(lvl [2]int) *levels.Level {
    switch lvl[0] {
    case 1:
        return levels.LvlOne(lvl[1])
    case 2:
        return levels.LvlTwo(lvl[1])
    }
    return levels.LvlOne(lvl[1])
}

func init() {
    fon, err = truetype.Parse(gomonobold.TTF)
    if err != nil {
        log.Fatal(err)
    }
    fo = truetype.NewFace(fon, &truetype.Options{Size: 20})

    downarrowimage, _, err := image.Decode(bytes.NewReader(assets.DownArrow_PNG))
    if err != nil {
        log.Fatal(err)
    }
    downArrowImage = ebiten.NewImageFromImage(downarrowimage)

    pcimage, _, err := image.Decode(bytes.NewReader(pcimages.PC_png))
    if err != nil {
        log.Fatal(err)
    }
    pcImage = ebiten.NewImageFromImage(pcimage)

    pixels := []uint8{}
    for a := 0; a < 442368; a++ {
        pixels = append(pixels, 0x40)
    }

    fadeImage = ebiten.NewImageFromImage(&image.Alpha{
        Pix: pixels,
        Stride: 768,
        Rect: image.Rect(0, 0, 768, 576),
    })

    if cont {
        p = &player.Player{Pos: [2]int{0, 0}, Image: pcImage}
        load = true
    } else {
        l = levels.LvlOne(0)
        p = &player.Player{Pos: [2]int{-l.Pos[0], -l.Pos[1]}, Image: pcImage}
        save = true
    }
}

func main() {
    ebiten.SetWindowSize(768, 576)
    ebiten.SetWindowTitle("CHANGEME")

    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
