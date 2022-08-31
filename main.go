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
    "reflect"
    "strconv"
    "strings"

    "golang.org/x/image/font"
    "golang.org/x/image/font/gofont/gomonobold"

    "github.com/golang/freetype/truetype"

    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/cutscenes"
    "github.com/jsnider-mtu/quailgame/inventory"
    "github.com/jsnider-mtu/quailgame/inventory/items"
    "github.com/jsnider-mtu/quailgame/levels"
    "github.com/jsnider-mtu/quailgame/player"
    "github.com/jsnider-mtu/quailgame/player/pcimages"
    "github.com/jsnider-mtu/quailgame/utils"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"

    _ "github.com/mattn/go-sqlite3"
)

var (
    sb strings.Builder
    err error
    start bool = true
    startanimdone bool = false
    startsel int = 0
    selload bool = false
    loads [][2]string
    loadsel int = 0
    findloads bool = true
    pause bool = false
    pausesel int = 0
    save bool = false
    firstsave bool = false
    load bool = false
    cont bool = false
    name string = "tempname"
    downArrowImage *ebiten.Image
    pcImage *ebiten.Image
    startImage *ebiten.Image
    lightningImage *ebiten.Image
    rainImage *ebiten.Image
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
    newlvl []interface{}
    f int = 0
    fadeImage *ebiten.Image
    dab int = 0
    dialogCount int = 0
    overwritewarning bool = false
    overwritesel int = 0
    y int = 0
    z int = 0
    a int = 0
    loadsfound bool = false
    cutscene bool = false
    csCount int = 0
    curCS int = 0
    csDone []int
    fadeScreen *ebiten.Image
    savesTableSchema []string
    schemaRowsCount int = 0
    copyColsCount int = 0
    colsStr string
    animCount int = 0
)

type Game struct {}

func (g *Game) Update() error {
    if start {
        if startanimdone {
            if findloads {
                homeDir, err := os.UserHomeDir()
                if err != nil {
                    log.Fatal(err)
                }
                db, err := sql.Open("sqlite3", homeDir + "/quailsaves.db")
                if err != nil {
                    log.Fatal(err)
                }
                defer db.Close()
                rows, err := db.Query("select name, level from saves")
                if err != nil {
                    log.Fatal(err)
                }
                defer rows.Close()
                for rows.Next() {
                    var savename string
                    var levelname string
                    err = rows.Scan(&savename, &levelname)
                    loads = append(loads, [2]string{savename, levelname})
                }
                err = rows.Err()
                if err != nil {
                    log.Fatal(err)
                }
                findloads = false
                loadsfound = true
                if len(loads) == 0 {
                    selload = false
                }
            }
            if selload {
                if len(loads) == 0 {
                    selload = false
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                    if loadsel > 0 {
                        loadsel--
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                    if loadsel < len(loads) - 1 {
                        loadsel++
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyD) {
                    homeDir, err := os.UserHomeDir()
                    if err != nil {
                        log.Fatal(err)
                    }
                    db, err := sql.Open("sqlite3", homeDir + "/quailsaves.db")
                    if err != nil {
                        log.Fatal(err)
                    }
                    defer db.Close()
                    _, err = db.Exec("delete from saves where name = ?", loads[loadsel][0])
                    if err != nil {
                        log.Fatal(err)
                    }
                    loads = [][2]string{}
                    loadsfound = false
                    findloads = true
                    return nil
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
                    selload = false
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                    name = loads[loadsel][0]
                    p = &player.Player{Pos: [2]int{0, 0}, Inv: &inventory.Inv{}, Image: pcImage}
                    load = true
                    selload = false
                    start = false
                }
            } else if overwritewarning {
                if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
                    overwritesel = 0
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
                    overwritesel = 1
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
                    overwritewarning = false
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                    if overwritesel == 0 {
                        sb.Reset()
                        firstsave = false
                        start = false
                        save = true
                        overwritewarning = false
                    } else {
                        overwritewarning = false
                    }
                }
            } else if firstsave {
                if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
                    firstsave = false
                }
                if len(sb.String()) < 24 {
                    switch {
                    case inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.KeyPressDuration(ebiten.KeyA) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('A')
                        } else {
                            err = sb.WriteByte('a')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyB) || inpututil.KeyPressDuration(ebiten.KeyB) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('B')
                        } else {
                            err = sb.WriteByte('b')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyC) || inpututil.KeyPressDuration(ebiten.KeyC) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('C')
                        } else {
                            err = sb.WriteByte('c')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.KeyPressDuration(ebiten.KeyD) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('D')
                        } else {
                            err = sb.WriteByte('d')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyE) || inpututil.KeyPressDuration(ebiten.KeyE) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('E')
                        } else {
                            err = sb.WriteByte('e')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.KeyPressDuration(ebiten.KeyF) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('F')
                        } else {
                            err = sb.WriteByte('f')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyG) || inpututil.KeyPressDuration(ebiten.KeyG) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('G')
                        } else {
                            err = sb.WriteByte('g')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyH) || inpututil.KeyPressDuration(ebiten.KeyH) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('H')
                        } else {
                            err = sb.WriteByte('h')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyI) || inpututil.KeyPressDuration(ebiten.KeyI) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('I')
                        } else {
                            err = sb.WriteByte('i')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyJ) || inpututil.KeyPressDuration(ebiten.KeyJ) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('J')
                        } else {
                            err = sb.WriteByte('j')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyK) || inpututil.KeyPressDuration(ebiten.KeyK) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('K')
                        } else {
                            err = sb.WriteByte('k')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyL) || inpututil.KeyPressDuration(ebiten.KeyL) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('L')
                        } else {
                            err = sb.WriteByte('l')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyM) || inpututil.KeyPressDuration(ebiten.KeyM) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('M')
                        } else {
                            err = sb.WriteByte('m')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyN) || inpututil.KeyPressDuration(ebiten.KeyN) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('N')
                        } else {
                            err = sb.WriteByte('n')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyO) || inpututil.KeyPressDuration(ebiten.KeyO) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('O')
                        } else {
                            err = sb.WriteByte('o')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.KeyPressDuration(ebiten.KeyP) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('P')
                        } else {
                            err = sb.WriteByte('p')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.KeyPressDuration(ebiten.KeyQ) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('Q')
                        } else {
                            err = sb.WriteByte('q')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyR) || inpututil.KeyPressDuration(ebiten.KeyR) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('R')
                        } else {
                            err = sb.WriteByte('r')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.KeyPressDuration(ebiten.KeyS) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('S')
                        } else {
                            err = sb.WriteByte('s')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyT) || inpututil.KeyPressDuration(ebiten.KeyT) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('T')
                        } else {
                            err = sb.WriteByte('t')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyU) || inpututil.KeyPressDuration(ebiten.KeyU) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('U')
                        } else {
                            err = sb.WriteByte('u')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyV) || inpututil.KeyPressDuration(ebiten.KeyV) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('V')
                        } else {
                            err = sb.WriteByte('v')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.KeyPressDuration(ebiten.KeyW) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('W')
                        } else {
                            err = sb.WriteByte('w')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.KeyPressDuration(ebiten.KeyX) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('X')
                        } else {
                            err = sb.WriteByte('x')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyY) || inpututil.KeyPressDuration(ebiten.KeyY) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('Y')
                        } else {
                            err = sb.WriteByte('y')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.KeyZ) || inpututil.KeyPressDuration(ebiten.KeyZ) > 20:
                        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 {
                            err = sb.WriteByte('Z')
                        } else {
                            err = sb.WriteByte('z')
                        }
                    case inpututil.IsKeyJustPressed(ebiten.Key0) || inpututil.KeyPressDuration(ebiten.Key0) > 20:
                        err = sb.WriteByte('0')
                    case inpututil.IsKeyJustPressed(ebiten.Key1) || inpututil.KeyPressDuration(ebiten.Key1) > 20:
                        err = sb.WriteByte('1')
                    case inpututil.IsKeyJustPressed(ebiten.Key2) || inpututil.KeyPressDuration(ebiten.Key2) > 20:
                        err = sb.WriteByte('2')
                    case inpututil.IsKeyJustPressed(ebiten.Key3) || inpututil.KeyPressDuration(ebiten.Key3) > 20:
                        err = sb.WriteByte('3')
                    case inpututil.IsKeyJustPressed(ebiten.Key4) || inpututil.KeyPressDuration(ebiten.Key4) > 20:
                        err = sb.WriteByte('4')
                    case inpututil.IsKeyJustPressed(ebiten.Key5) || inpututil.KeyPressDuration(ebiten.Key5) > 20:
                        err = sb.WriteByte('5')
                    case inpututil.IsKeyJustPressed(ebiten.Key6) || inpututil.KeyPressDuration(ebiten.Key6) > 20:
                        err = sb.WriteByte('6')
                    case inpututil.IsKeyJustPressed(ebiten.Key7) || inpututil.KeyPressDuration(ebiten.Key7) > 20:
                        err = sb.WriteByte('7')
                    case inpututil.IsKeyJustPressed(ebiten.Key8) || inpututil.KeyPressDuration(ebiten.Key8) > 20:
                        err = sb.WriteByte('8')
                    case inpututil.IsKeyJustPressed(ebiten.Key9) || inpututil.KeyPressDuration(ebiten.Key9) > 20:
                        err = sb.WriteByte('9')
                    case inpututil.IsKeyJustPressed(ebiten.KeyBackspace) || inpututil.KeyPressDuration(ebiten.KeyBackspace) > 20:
                        str := sb.String()
                        if len(str) > 0 {
                            str = str[:len(str) - 1]
                            sb.Reset()
                            _, err = sb.WriteString(str)
                        }
                    }
                } else {
                    if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
                        str := sb.String()
                        if len(str) > 0 {
                            str = str[:len(str) - 1]
                            sb.Reset()
                            _, err = sb.WriteString(str)
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                    if len(sb.String()) > 0 {
                        name = sb.String()
                        homeDir, err := os.UserHomeDir()
                        if err != nil {
                            log.Fatal(err)
                        }
                        db, err := sql.Open("sqlite3", homeDir + "/quailsaves.db")
                        if err != nil {
                            log.Fatal(err)
                        }
                        defer db.Close()
                        rows, err := db.Query("select name from saves")
                        if err != nil {
                            log.Fatal(err)
                        }
                        defer rows.Close()
                        var savename string
                        for rows.Next() {
                            err = rows.Scan(&savename)
                            if name == savename {
                                overwritewarning = true
                                return nil
                            }
                        }
                        err = rows.Err()
                        if err != nil {
                            log.Fatal(err)
                        }
                        sb.Reset()
                        l = levels.LoadLvl("One", 0)
                        p.Pos[0] = -l.Pos[0]
                        p.Pos[1] = -l.Pos[1]
                        down = true
                        up = false
                        left = false
                        right = false
                        firstsave = false
                        start = false
                        save = true
                        curCS = 0
                        cutscene = true
                    }
                }
            } else {
                if !loadsfound {
                    findloads = true
                    return nil
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                    if startsel > 0 {
                        if len(loads) > 0 {
                            startsel--
                        } else {
                            startsel -= 2
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                    if startsel < 2 {
                        if len(loads) > 0 {
                            startsel++
                        } else {
                            startsel += 2
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                    switch startsel {
                    case 0:
                        if l == nil {
                            l = levels.LvlOne(0)
                            p = &player.Player{Pos: [2]int{-l.Pos[0], -l.Pos[1]}, Inv: &inventory.Inv{Cap: 8}, Image: pcImage}
                        }
                        firstsave = true
                    case 1:
                        if loadsfound {
                            if len(loads) != 0 {
                                selload = true
                            }
                        }
                    case 2:
                        os.Exit(0)
                    }
                }
            }
        } else {
            if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
                startanimdone = true
            }
        }
        return nil
    } else {
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
                if pausesel < 3 {
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
                    start = true
                    loads = [][2]string{}
                    loadsfound = false
                    findloads = true
                    pause = false
                case 3:
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
                qMarks := "?" + strings.Repeat(", ?", len(savesTableSchema) - 1)
                saveStmt := "insert or replace into saves("
                for cind, col := range savesTableSchema {
                    colArr := strings.Split(col, ",")
                    if cind == len(savesTableSchema) - 1 {
                        saveStmt += colArr[0] + ") values(" + qMarks + ");"
                    } else {
                        saveStmt += colArr[0] + ", "
                    }
                }
                var csdonestr string
                for csdoneind, csdoneval := range csDone {
                    if csdoneind == len(csDone) - 1 {
                        csdonestr += strconv.Itoa(csdoneval)
                    } else {
                        csdonestr += strconv.Itoa(csdoneval) + ","
                    }
                }
                var invstr string
                for itemind, item := range p.Inv.Items {
                    if itemind == len(p.Inv.Items) - 1 {
                        invstr += item.Save()
                    } else {
                        invstr += item.Save() + ";"
                    }
                }
                _, err = db.Exec(saveStmt, name, l.Name, l.Pos[0], l.Pos[1], p.Inv.Cap, csdonestr, invstr)
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
                var invcap int
                var csdonestr string
                var invstr string
                for rows.Next() {
                    err = rows.Scan(&savename, &levelname, &x, &y, &invcap, &csdonestr, &invstr)
                }
                err = rows.Err()
                if err != nil {
                    log.Fatal(err)
                }
                csdonestrarr := strings.Split(csdonestr, ",")
                csDone = []int{}
                for _, numstr := range csdonestrarr {
                    if numstr == "" {
                        break
                    }
                    numint, err := strconv.Atoi(numstr)
                    if err != nil {
                        log.Fatal(err)
                    }
                    csDone = append(csDone, numint)
                }
                invstrarr := strings.Split(invstr, ";")
                for _, item := range invstrarr {
                    if item == "" {
                        break
                    }
                    itemprops := strings.Split(item, ",")
                    p.Inv.Items = append(p.Inv.Items, items.LoadItem(itemprops[0], itemprops[1], itemprops[2]))
                }
                l = levels.LoadLvl(levelname, 0, x, y)
                p.Pos = [2]int{-l.Pos[0], -l.Pos[1]}
                p.Inv.Cap = invcap
                load = false
            }
            if cutscene {
                csCount++
            } else if csCount > 0 {
                csCount = 0
            }
            if npcCount == 6000 {
                npcCount = 0
            }
            animCount++
            if animCount == 4000 {
                animCount = 0
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
                                npc.Direction = "down"
                                npcname = npc.Name
                                dialogstrs = npc.Dialog()
                                dialogopen = true
                            }
                        }
                    }
                case down:
                    for _, npc := range l.NPCs {
                        if npc.PC.Pos[0] >= p.Pos[0] - 24 && npc.PC.Pos[0] <= p.Pos[0] + 24 && npc.PC.Pos[1] - 24 == p.Pos[1] {
                            if !dialogopen {
                                npc.Direction = "up"
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
                                npc.Direction = "right"
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
                                npc.Direction = "left"
                                npcname = npc.Name
                                dialogstrs = npc.Dialog()
                                dialogopen = true
                            }
                        }
                    }
                }
            }
            if !dialogopen && !lvlchange && !start {
                for _, npc := range l.NPCs {
                    if npc.Speed > 0 && (npcCount + npc.Offset) % npc.Speed == 0 {
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
                dirarr := [4]int{inpututil.KeyPressDuration(ebiten.KeyW), inpututil.KeyPressDuration(ebiten.KeyA),
                                 inpututil.KeyPressDuration(ebiten.KeyD), inpututil.KeyPressDuration(ebiten.KeyS)}
                var smallestnum int = 0
                var smallestind int = 4
                for smind, smnum := range dirarr {
                    if smnum > 0 {
                        smallestnum = smnum
                        smallestind = smind
                        break
                    }
                }
                if smallestnum > 0 {
                    for sind, snum := range dirarr {
                        if snum > 0 {
                            if snum < smallestnum {
                                smallestind = sind
                            }
                        }
                    }
                }
                switch smallestind {
                case 0:
                    stopped = false
                    up = true
                    down = false
                    left = false
                    right = false
                    if smallestnum % 4 == 0 {
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
                case 1:
                    stopped = false
                    left = true
                    up = false
                    down = false
                    right = false
                    if smallestnum % 4 == 0 {
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
                case 2:
                    stopped = false
                    right = true
                    left = false
                    up = false
                    down = false
                    if smallestnum % 4 == 0 {
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
                case 3:
                    stopped = false
                    down = true
                    up = false
                    left = false
                    right = false
                    if smallestnum % 4 == 0 {
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
                case 4:
                    stopped = true
                    count = 0
                }
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
    mcdrawn := false
    if !startanimdone {
        y++
        if y <= 65 {
            animgm := ebiten.GeoM{}
            animgm.Translate(float64(0), float64(h - (9 * y)))
            screen.DrawImage(
                startImage, &ebiten.DrawImageOptions{
                    GeoM: animgm})
        } else {
            screen.DrawImage(startImage, nil)
            y = 0
            startanimdone = true
        }
    } else if start {
        screen.DrawImage(startImage, nil)
        animop := &ebiten.DrawImageOptions{}
        animop.GeoM.Scale(4.0, 4.0)
        animop.GeoM.Translate(float64((w / 2) - 40), float64(0))
        animop.ColorM.Scale(1.0, 1.0, 1.0, 0.60)
        z++
        if z >= 20 && z < 40 {
            a := (z / 4) % 5
            if a == 0 || a == 4 {
                animop.ColorM.Scale(1.0, 1.0, 1.0, 0.50)
            }
            screen.DrawImage(
                lightningImage.SubImage(
                    image.Rect(
                        (a % 5) * 80, 0, ((a % 5) + 1) * 80, 96)).(*ebiten.Image),
                        animop)
        } else if z == 300 {
                z = 0
                a = 0
        }
        for b := 0; b < 9; b++ {
            for c := 0; c < 12; c++ {
                raingm := ebiten.GeoM{}
                raingm.Reset()
                raingm.Translate(float64(c * 64), float64(b * 64))
                screen.DrawImage(
                    rainImage.SubImage(
                        image.Rect(
                            (z % 16) * 64, 0, ((z % 16) + 1) * 64, 64)).(*ebiten.Image),
                            &ebiten.DrawImageOptions{GeoM: raingm})
            }
        }
        if selload {
            r := text.BoundString(fo, fmt.Sprint("> aaaaaaaaaaaaaaaaaaaaaaaa -- Level: aaaaaaaaaaaa"))
            hei := r.Max.Y - r.Min.Y
            wid := r.Max.X - r.Min.X
            for ind, lo := range loads {
                savesuffix := 24 - len(lo[0])
                if loadsel == len(loads) {
                    loadsel--
                }
                if loadsel > 15 {
                    if loadsel == ind {
                        text.Draw(screen, fmt.Sprintf("> %s -- Level: %s", lo[0] + strings.Repeat(" ", savesuffix), lo[1]), fo, (w / 2) - (wid / 2), (hei * 2 * (ind - (ind - 16))), color.White)
                    } else {
                        text.Draw(screen, fmt.Sprintf("  %s -- Level: %s", lo[0] + strings.Repeat(" ", savesuffix), lo[1]), fo, (w / 2) - (wid / 2), (hei * 2 * (16 - (loadsel - ind))), color.White)
                    }
                } else {
                    if loadsel == ind {
                        text.Draw(screen, fmt.Sprintf("> %s -- Level: %s", lo[0] + strings.Repeat(" ", savesuffix), lo[1]), fo, (w / 2) - (wid / 2), (hei * 2 * (ind + 1)), color.White)
                    } else {
                        text.Draw(screen, fmt.Sprintf("  %s -- Level: %s", lo[0] + strings.Repeat(" ", savesuffix), lo[1]), fo, (w / 2) - (wid / 2), (hei * 2 * (ind + 1)), color.White)
                    }
                }
            }
        } else if overwritewarning {
            warning := "           WARNING!!!\n\nYou will overwrite a previous save\n\n           Continue??\n"
            selection := "      > Yes <         > No <"
            r := text.BoundString(fo, warning + selection)
            hei := r.Max.Y - r.Min.Y
            wid := r.Max.X - r.Min.X
            warninggm := ebiten.GeoM{}
            warninggm.Translate(float64((w / 2) - (wid / 2) - 8), float64((h / 2) - (hei / 2) - 24))
            warningimg := ebiten.NewImage(wid + 16, (2 * hei) + 16)
            warningimg.Fill(color.Black)
            screen.DrawImage(
                warningimg, &ebiten.DrawImageOptions{
                    GeoM: warninggm})
            text.Draw(screen, warning, fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2), color.White)
            if overwritesel == 0 {
                text.Draw(screen, "      > Yes <           No  ", fo, (w / 2) - (wid / 2), (h / 2) + hei, color.White)
            } else {
                text.Draw(screen, "        Yes           > No <", fo, (w / 2) - (wid / 2), (h / 2) + hei, color.White)
            }
        } else if firstsave {
            r := text.BoundString(fo, "aaaaaaaaaaaaaaaaaaaaaaaa")
            hei := r.Max.Y - r.Min.Y
            wid := r.Max.X - r.Min.X
            inputgm := ebiten.GeoM{}
            inputgm.Translate(float64((w / 2) - (wid / 2) - 8), float64((h / 2) - (hei / 2) - 16))
            inputimg := ebiten.NewImage(wid + 8, hei + 16)
            inputimg.Fill(color.Black)
            screen.DrawImage(
                inputimg, &ebiten.DrawImageOptions{
                    GeoM: inputgm})
            text.Draw(screen, sb.String(), fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + (3 * hei / 8), color.White)
            r2 := text.BoundString(fo, "Name")
            hei2 := r2.Max.Y - r2.Min.Y
            wid2 := r2.Max.X - r2.Min.X
            text.Draw(screen, "Name", fo, (w / 2) - (wid2 / 2), (h / 2) - (hei2 * 2), color.White)
        } else {
            r := text.BoundString(fo, "> New Game <")
            hei := r.Max.Y - r.Min.Y
            wid := r.Max.X - r.Min.X
            switch startsel {
            case 0:
                text.Draw(screen, "> New Game <", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 4), color.White)
                if len(loads) == 0 {
                    text.Draw(screen, "  Continue  ", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 6), color.Gray16{0x8000})
                } else {
                    text.Draw(screen, "  Continue  ", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 6), color.White)
                }
                text.Draw(screen, "  Quit Game  ", fo, (w / 2) - (wid / 2) - (wid / 24), (h / 2) + (hei * 8), color.White)
            case 1:
                if len(loads) == 0 {
                    startsel = 0
                }
                text.Draw(screen, "  New Game  ", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 4), color.White)
                text.Draw(screen, "> Continue <", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 6), color.White)
                text.Draw(screen, "  Quit Game  ", fo, (w / 2) - (wid / 2) - (wid / 24), (h / 2) + (hei * 8), color.White)
            case 2:
                text.Draw(screen, "  New Game  ", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 4), color.White)
                if len(loads) == 0 {
                    text.Draw(screen, "  Continue  ", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 6), color.Gray16{0x8000})
                } else {
                    text.Draw(screen, "  Continue  ", fo, (w / 2) - (wid / 2), (h / 2) + (hei * 6), color.White)
                }
                text.Draw(screen, "> Quit Game <", fo, (w / 2) - (wid / 2) - (wid / 24), (h / 2) + (hei * 8), color.White)
            }
        }
    } else if cutscene {
        for _, csval := range csDone {
            if csval == curCS {
                cutscene = false
                break
            }
        }
        if cutscene {
            done, rstCount := cutscenes.CutScene(screen, curCS, csCount, &fo)
            if rstCount {
                csCount = 0
            }
            if done {
                csDone = append(csDone, curCS)
                cutscene = false
            }
        }
    } else if l != nil {
        lgm := ebiten.GeoM{}
        lgm.Translate(float64((w / 2) + l.Pos[0]), float64((h / 2) + l.Pos[1]))
        screen.DrawImage(l.Image, &ebiten.DrawImageOptions{GeoM: lgm})
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
        if !mcdrawn && !start && !cutscene {
            drawmc(screen, w, h)
        }
        l.Anim(screen, l, animCount, w, h)
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
        pauseimg := ebiten.NewImage(wid + 28, (hei * 5) + 64)
        pauseimg.Fill(color.Black)
        screen.DrawImage(
            pauseimg, &ebiten.DrawImageOptions{
                GeoM: pausegm})
        pausegm2 := ebiten.GeoM{}
        pausegm2.Translate(float64((w / 2) - (wid / 2) - 4), float64((h / 2) - (3 * hei / 2) - 12))
        pauseimg2 := ebiten.NewImage(wid + 20, (hei * 5) + 56)
        pauseimg2.Fill(color.White)
        screen.DrawImage(
            pauseimg2, &ebiten.DrawImageOptions{
                GeoM: pausegm2})
        switch pausesel {
        case 0:
            text.Draw(screen, "> Save game", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            text.Draw(screen, "  Load last", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            text.Draw(screen, "  save", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
            text.Draw(screen, "  Main menu", fo, (w / 2) - (wid / 2), (h / 2) + (3 * hei / 2) + 40, color.Black)
            text.Draw(screen, "  Quit game", fo, (w / 2) - (wid / 2), (h / 2) + (5 * hei / 2) + 48, color.Black)
        case 1:
            text.Draw(screen, "  Save game", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            text.Draw(screen, "> Load last", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            text.Draw(screen, "  save", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
            text.Draw(screen, "  Main menu", fo, (w / 2) - (wid / 2), (h / 2) + (3 * hei / 2) + 40, color.Black)
            text.Draw(screen, "  Quit game", fo, (w / 2) - (wid / 2), (h / 2) + (5 * hei / 2) + 48, color.Black)
        case 2:
            text.Draw(screen, "  Save game", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            text.Draw(screen, "  Load last", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            text.Draw(screen, "  save", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
            text.Draw(screen, "> Main menu", fo, (w / 2) - (wid / 2), (h / 2) + (3 * hei / 2) + 40, color.Black)
            text.Draw(screen, "  Quit game", fo, (w / 2) - (wid / 2), (h / 2) + (5 * hei / 2) + 48, color.Black)
        case 3:
            text.Draw(screen, "  Save game", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            text.Draw(screen, "  Load last", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            text.Draw(screen, "  save", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
            text.Draw(screen, "  Main menu", fo, (w / 2) - (wid / 2), (h / 2) + (3 * hei / 2) + 40, color.Black)
            text.Draw(screen, "> Quit game", fo, (w / 2) - (wid / 2), (h / 2) + (5 * hei / 2) + 48, color.Black)
        }
    }
    if lvlchange {
        op := &ebiten.DrawImageOptions{}
        fadeScreen = ebiten.NewImage(768, 576)
        fadeScreen.Fill(color.Black)
        if npcCount % 10 == 0 {
            f++
        }
        if f == 0 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.0)
            screen.DrawImage(fadeImage, op)
        } else if f == 1 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.2)
            screen.DrawImage(fadeScreen, op)
        } else if f == 2 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.4)
            screen.DrawImage(fadeScreen, op)
        } else if f == 3 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.6)
            screen.DrawImage(fadeScreen, op)
        } else if f == 4 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.8)
            screen.DrawImage(fadeScreen, op)
        } else if f == 5 {
            screen.DrawImage(fadeScreen, nil)
            f = 0
            lvlchange = false
            l = levels.LoadLvl(newlvl...)
            p.Pos[0] = -l.Pos[0]
            p.Pos[1] = -l.Pos[1]
            if l.Cutscene > 0 {
                curCS = l.Cutscene
                cutscene = true
            }
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
            i := (count / 10) % 4
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
            i := (count / 10) % 4
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
            i := (count / 10) % 4
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
            i := (count / 10) % 4
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

func init() {
    fon, err = truetype.Parse(gomonobold.TTF)
    if err != nil {
        log.Fatal(err)
    }
    fo = truetype.NewFace(fon, &truetype.Options{Size: 20})

    startimage, _, err := image.Decode(bytes.NewReader(assets.Start_PNG))
    if err != nil {
        log.Fatal(err)
    }
    startImage = ebiten.NewImageFromImage(startimage)

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
        pixels = append(pixels, 0x33)
    }

    fadeImage = ebiten.NewImageFromImage(&image.Alpha{
        Pix: pixels,
        Stride: 768,
        Rect: image.Rect(0, 0, 768, 576),
    })

    lightningimage, _, err := image.Decode(bytes.NewReader(assets.Lightning_PNG))
    if err != nil {
        log.Fatal(err)
    }
    lightningImage = ebiten.NewImageFromImage(lightningimage)

    rainimage, _, err := image.Decode(bytes.NewReader(assets.RainHeavy_PNG))
    if err != nil {
        log.Fatal(err)
    }
    rainImage = ebiten.NewImageFromImage(rainimage)

    savesTableSchema = []string{"name,TEXT,1,null,1", "level,TEXT,1,\"One\",0", "x,INT,1,null,0", "y,INT,1,null,0", "invcap,INT,1,8,0", "csdone,TEXT,0,null,0", "inventory,TEXT,0,null,0"}
    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    db, err := sql.Open("sqlite3", homeDir + "/quailsaves.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    var createStmt string = "create table if not exists saves ("
    for cind, col := range savesTableSchema {
        colArr := strings.Split(col, ",")
        createStmt += colArr[0] + " " + colArr[1]
        if colArr[2] == "1" {
            createStmt += " not null"
        }
        if colArr[3] != "null" {
            createStmt += " default " + colArr[3]
        }
        if colArr[4] == "1" {
            createStmt += " primary key"
        }
        if cind == len(savesTableSchema) - 1 {
            createStmt += ");"
        } else {
            createStmt += ", "
        }
    }
    _, err = db.Exec(createStmt)
    if err != nil {
        log.Fatal(fmt.Sprintf("%q: %s\n", err, createStmt))
    }
    schemaRows, err := db.Query("PRAGMA table_info(saves)")
    if err != nil {
        log.Fatal(err)
    }
    defer schemaRows.Close()
    fixSchema := false
    for schemaRows.Next() {
        schemaRowsCount++
        var schemaRowsIndex int
        var schemaRowsName string
        var schemaRowsType string
        var schemaRowsNotNull int
        var schemaRowsDefault interface{}
        var schemaRowsPk int
        err = schemaRows.Scan(&schemaRowsIndex, &schemaRowsName, &schemaRowsType, &schemaRowsNotNull, &schemaRowsDefault, &schemaRowsPk)
        if schemaRowsDefault == nil {
            schemaRowsDefault = "null"
        }
        if savesTableSchema[schemaRowsIndex] != schemaRowsName + "," + schemaRowsType + "," + strconv.Itoa(schemaRowsNotNull) + "," + fmt.Sprint(schemaRowsDefault) + "," + strconv.Itoa(schemaRowsPk) {
            fixSchema = true
        }
    }
    err = schemaRows.Err()
    if err != nil {
        log.Fatal(err)
    }
    if fixSchema {
        copyStmt := `
        create table copied as select * from saves;
        `
        _, err = db.Exec(copyStmt)
        if err != nil {
            log.Fatal(fmt.Sprintf("%q: %s\n", err, copyStmt))
        }
        dropStmt := `
        drop table saves;
        `
        _, err = db.Exec(dropStmt)
        if err != nil {
            log.Fatal(fmt.Sprintf("%q: %s\n", err, dropStmt))
        }
        _, err = db.Exec(createStmt)
        if err != nil {
            log.Fatal(fmt.Sprintf("%q: %s\n", err, createStmt))
        }
        copyCols, err := db.Query("PRAGMA table_info(copied)")
        if err != nil {
            log.Fatal(err)
        }
        defer copyCols.Close()
        for copyCols.Next() {
            copyColsCount++
        }
        for copyColsIndex := 0; copyColsIndex < copyColsCount; copyColsIndex++ {
            colsarr := strings.Split(savesTableSchema[copyColsIndex], ",")
            if copyColsIndex == copyColsCount - 1 {
                colsStr += colsarr[0]
            } else {
                colsStr += colsarr[0] + ", "
            }
        }
        copyRows, err := db.Query("select " + colsStr + " from copied")
        if err != nil {
            log.Fatal(err)
        }
        defer copyRows.Close()
        var insertStmts = make([]string, 0)
        var copyRowsPtrs = make([]interface{}, copyColsCount)
        var copyRowsArr = make([]interface{}, copyColsCount)
        for i, _ := range copyRowsPtrs {
            copyRowsPtrs[i] = &copyRowsArr[i]
        }
        var nullStr string
        var newColsInd int = len(savesTableSchema) - (len(savesTableSchema) - copyColsCount)
        for copyRows.Next() {
            err = copyRows.Scan(copyRowsPtrs...)
            insertStmt := "insert into saves ("
            for cind, col := range savesTableSchema {
                insStmtDone := false
                skip := false
                colArr := strings.Split(col, ",")
                if cind >= newColsInd {
                    if colArr[2] != "1" {
                        nullStr += ", null"
                    } else if colArr[3] != "null" {
                        if cind == len(savesTableSchema) - 1 {
                            insertStmt += ") values("
                            insStmtDone = true
                        } else {
                            skip = true
                        }
                    } else {
                        log.Fatal(fmt.Sprintf("%s is NOT NULL but DEFAULT is \"null\"", colArr[0]))
                    }
                }
                if cind == len(savesTableSchema) - 1 {
                    if !insStmtDone {
                        insertStmt += colArr[0] + ") values("
                    }
                } else {
                    if !skip {
                        insertStmt += colArr[0] + ", "
                    }
                }
            }
            for whatever, whateverPtr := range copyRowsArr {
                switch reflect.TypeOf(whateverPtr).String() {
                case "string":
                    if whatever == len(copyRowsArr) - 1 {
                        insertStmt += "\"" + fmt.Sprint(whateverPtr) + "\"" + nullStr + ");"
                    } else {
                        insertStmt += "\"" + fmt.Sprint(whateverPtr) + "\", "
                    }
                case "int64":
                    if whatever == len(copyRowsArr) - 1 {
                        insertStmt += fmt.Sprint(whateverPtr) + nullStr + ");"
                    } else {
                        insertStmt += fmt.Sprint(whateverPtr) + ", "
                    }
                }
            }
            insertStmts = append(insertStmts, insertStmt)
        }
        for _, insStmt := range insertStmts {
            _, err = db.Exec(insStmt)
            if err != nil {
                log.Fatal(fmt.Sprintf("%q: %s\n", err, insStmt))
            }
        }
        err = copyRows.Err()
        if err != nil {
            log.Fatal(err)
        }
        copyDropStmt := `
        drop table copied;
        `
        _, err = db.Exec(copyDropStmt)
        if err != nil {
            log.Fatal(fmt.Sprintf("%q: %s\n", err, copyDropStmt))
        }
    }
    if schemaRowsCount < len(savesTableSchema) && !fixSchema {
        for colMisIndex := len(savesTableSchema) - schemaRowsCount; colMisIndex > 0; colMisIndex-- {
            var colopts string
            colarr := strings.Split(savesTableSchema[len(savesTableSchema) - colMisIndex], ",")
            if colarr[2] == "1" {
                colopts += " not null"
            }
            if colarr[3] == "1" {
                colopts += " primary key"
            }
            alterStmt := "alter table saves add column " + colarr[0] + " " + colarr[1] + colopts + ";"
            _, err = db.Exec(alterStmt)
            if err != nil {
                log.Fatal(err)
            }
        }
    }
}

func main() {
    ebiten.SetWindowSize(768, 576)
    ebiten.SetWindowTitle("CHANGEME")

    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
