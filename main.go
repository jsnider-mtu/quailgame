package main

import (
    "bytes"
    "database/sql"
    "errors"
    "fmt"
    "image"
    "image/color"
    _ "image/png"
    "log"
    "math/rand"
    "os"
    "reflect"
    "sort"
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
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
    overworld bool = false
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
    overworldImage *ebiten.Image
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
    colsStr string
    animCount int = 0
    icon16img image.Image
    icon32img image.Image
    icon48img image.Image
    creation bool = false
    creationsel int = 0
    creationpage [4]int
    racesel int = 0
    classsel int = 0
    backgroundsel int = 0
    equipmentsel int = 0
    choices bool = false
    racechoices bool = false
    dupwarning bool = false
    option0 int = 0
    option1 int = 0
    option2 int = 0
    option3 int = 0
    option4 int = 0
    option5 int = 0
    option6 int = 0
    option7 int = 0
    option8 int = 0
    raceopt0 int = 0
    raceopt1 int = 0
    raceopt2 int = 0
    raceopt3 int = 0
    raceopt4 int = 0
    ac int
    str int
    dex int
    con int
    intel int
    wis int
    cha int
    pb int
    hp int
    hd string
    speed int
    size int // 0: Small, 1: Medium, 2: Large
    darkvision bool = false
    lucky bool = false
    nimbleness bool = false
    brave bool = false
    ancestry string
)

var racemap = make(map[int]string)
var classmap = make(map[int]string)
//var backgroundmap = make(map[int]string)
var equipmentmap = make(map[int]string)
var abilities = make([]int, 6)
var savingthrows = make(map[string]int)
var languages = make([]string, 0)
var proficiencies = make([]string, 0)
var resistances = make([]string, 0)

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
                        creation = true
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
                            l = levels.LoadLvl("One", 0)
                            p = &player.Player{Pos: [2]int{-l.Pos[0], -l.Pos[1]}, Inv: &inventory.Inv{}, Image: pcImage}
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
    } else if creation {
        if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
            creationsel = 0
            choices = false
            creation = false
            start = true
        }
        // character creation
        if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
            creationsel--
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
            switch creationsel {
            case 0:
                racesel--
                if racesel < 0 {
                    racesel = 0
                }
            case 1:
                classsel--
                if classsel < 0 {
                    classsel = 0
                }
            default:
                return errors.New("Out of bounds (577)")
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
            switch creationsel {
            case 0:
                racesel++
                if racesel > 8 {
                    racesel = 8
                }
            case 1:
                classsel++
                if classsel > 11 {
                    classsel = 11
                }
            default:
                return errors.New("Out of bounds (Update)")
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
            creationsel++
        }
        if creationsel < 0 {
            creationsel = 0
        } else if creationsel > 1 {
            creationsel = 0
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
            // Save player info; gen stats
            // Roll ability scores; sort then assign based on class
            onescore := make([]int, 4)
            for x := 0; x < 6; x++ {
                for a := 0; a < 4; a++ {
                    onescore[a] = rand.Intn(6) + 1
                }
                // sort onescore
                sort.Slice(onescore, func(i, j int) bool {
                    return onescore[i] > onescore[j]
                })
                score := onescore[0] + onescore[1] + onescore[2]
                abilities[x] = score
            }
            sort.Slice(abilities, func(i, j int) bool {
                return abilities[i] > abilities[j]
            })
            switch classsel {
            case 0:
                str = abilities[0]
                con = abilities[1]
                dex = abilities[2]
                intel = abilities[3]
                wis = abilities[4]
                cha = abilities[5]
                pb = 2
                hp = 12 + con
                hd = "1d12"
                savingthrows["str"] = ((str - 10) / 2) + pb
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = ((con - 10) / 2) + pb
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = (wis - 10) / 2
                savingthrows["cha"] = (cha - 10) / 2
                proficiencies = append(proficiencies,
                    "light armor", "medium armor", "shields",
                    "simple weapons", "martial weapons")
            case 1:
                cha = abilities[0]
                dex = abilities[1]
                con = abilities[2]
                intel = abilities[3]
                wis = abilities[4]
                str = abilities[5]
                pb = 2
                hp = 8 + con
                hd = "1d8"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = ((dex - 10) / 2) + pb
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = (wis - 10) / 2
                savingthrows["cha"] = ((cha - 10) / 2) + pb
                proficiencies = append(proficiencies,
                    "light armor", "simple weapons", "hand crossbows",
                    "longswords", "rapiers", "shortswords") // 3 instruments
            case 2:
                wis = abilities[0]
                con = abilities[1]
                str = abilities[2]
                wis = abilities[3]
                intel = abilities[4]
                cha = abilities[5]
                pb = 2
                hp = 8 + con
                hd = "1d8"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = ((wis - 10) / 2) + pb
                savingthrows["cha"] = ((cha - 10) / 2) + pb
                proficiencies = append(proficiencies,
                    "light armor", "medium armor", "shields",
                    "simple weapons")
            case 3:
                wis = abilities[0]
                con = abilities[1]
                dex = abilities[2]
                intel = abilities[3]
                str = abilities[4]
                cha = abilities[5]
                pb = 2
                hp = 8 + con
                hd = "1d8"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = ((intel - 10) / 2) + pb
                savingthrows["wis"] = ((wis - 10) / 2) + pb
                savingthrows["cha"] = (cha - 10) / 2
                proficiencies = append(proficiencies,
                    "light armor", "medium armor", "shields",
                    "clubs", "daggers", "darts", "javelins", "maces",
                    "quarterstaffs", "scimitars", "sickles", "slings",
                    "spears", "herbalism kit")
            case 4:
                str = abilities[0]
                con = abilities[1]
                dex = abilities[2]
                intel = abilities[3]
                wis = abilities[4]
                cha = abilities[5]
                pb = 2
                hp = 10 + con
                hd = "1d10"
                savingthrows["str"] = ((str - 10) / 2) + pb
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = ((con - 10) / 2) + pb
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = (wis - 10) / 2
                savingthrows["cha"] = (cha - 10) / 2
                proficiencies = append(proficiencies,
                    "all armor", "shields", "simple weapons", "martial weapons")
            case 5:
                dex = abilities[0]
                wis = abilities[1]
                str= abilities[2]
                con = abilities[3]
                intel = abilities[4]
                cha = abilities[5]
                pb = 2
                hp = 8 + con
                hd = "1d8"
                savingthrows["str"] = ((str - 10) / 2) + pb
                savingthrows["dex"] = ((dex - 10) / 2) + pb
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = (wis - 10) / 2
                savingthrows["cha"] = (cha - 10) / 2
                proficiencies = append(proficiencies,
                    "simple weapons", "shortswords") // one artisan tools or one instrument
            case 6:
                str = abilities[0]
                cha = abilities[1]
                con = abilities[2]
                wis = abilities[3]
                dex = abilities[4]
                intel = abilities[5]
                pb = 2
                hp = 10 + con
                hd = "1d10"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = ((wis - 10) / 2) + pb
                savingthrows["cha"] = ((cha - 10) / 2) + pb
                proficiencies = append(proficiencies,
                    "all armor", "shields", "simple weapons", "martial weapons")
            case 7:
                dex = abilities[0]
                wis = abilities[1]
                con = abilities[2]
                intel = abilities[3]
                str = abilities[4]
                cha = abilities[5]
                pb = 2
                hp = 10 + con
                hd = "1d10"
                savingthrows["str"] = ((str - 10) / 2) + pb
                savingthrows["dex"] = ((dex - 10) / 2) + pb
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = (wis - 10) / 2
                savingthrows["cha"] = (cha - 10) / 2
                proficiencies = append(proficiencies,
                    "light armor", "medium armor", "shields", "simple weapons",
                    "martial weapons")
            case 8:
                dex = abilities[0]
                cha = abilities[1]
                con = abilities[2]
                intel = abilities[3]
                wis = abilities[4]
                str = abilities[5]
                pb = 2
                hp = 8 + con
                hd = "1d8"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = ((dex - 10) / 2) + pb
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = ((intel - 10) / 2) + pb
                savingthrows["wis"] = (wis - 10) / 2
                savingthrows["cha"] = (cha - 10) / 2
                proficiencies = append(proficiencies,
                    "light armor", "simple weapons", "hand crossbows", "longswords",
                    "rapiers", "shortswords", "thieves tools")
            case 9:
                cha = abilities[0]
                con = abilities[1]
                intel = abilities[2]
                dex = abilities[3]
                wis = abilities[4]
                str = abilities[5]
                pb = 2
                hp = 6 + con
                hd = "1d6"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = ((con - 10) / 2) + pb
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = (wis - 10) / 2
                savingthrows["cha"] = ((cha - 10) / 2) + pb
                proficiencies = append(proficiencies,
                    "daggers", "darts", "slings", "quarterstaffs",
                    "light crossbows")
            case 10:
                cha = abilities[0]
                con = abilities[1]
                dex = abilities[2]
                intel = abilities[3]
                wis = abilities[4]
                str = abilities[5]
                pb = 2
                hp = 8 + con
                hd = "1d8"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = (intel - 10) / 2
                savingthrows["wis"] = ((wis - 10) / 2) + pb
                savingthrows["cha"] = ((cha - 10) / 2) + pb
                proficiencies = append(proficiencies,
                    "light armor", "simple weapons")
            case 11:
                intel = abilities[0]
                con = abilities[1]
                dex = abilities[2]
                cha = abilities[3]
                wis = abilities[4]
                str = abilities[5]
                pb = 2
                hp = 6 + con
                hd = "1d6"
                savingthrows["str"] = (str - 10) / 2
                savingthrows["dex"] = (dex - 10) / 2
                savingthrows["con"] = (con - 10) / 2
                savingthrows["intel"] = ((intel - 10) / 2) + pb
                savingthrows["wis"] = ((wis - 10) / 2) + pb
                savingthrows["cha"] = (cha - 10) / 2
                proficiencies = append(proficiencies,
                    "daggers", "darts", "slings", "quarterstaffs", "light crossbows")
            default:
                return errors.New("Invalid value for classsel")
            }
            switch racesel {
            case 0:
                con += 2
                hp += 2
                speed = 25
                size = 1
                languages = append(languages, "Common", "Dwarvish")
                proficiencies = append(proficiencies,
                    "battleaxe", "handaxe", "light hammer", "warhammer",
                    "smith tools", "brewer supplies", "mason tools")
                resistances = append(resistances, "poison")
                darkvision = true
            case 1:
                dex += 2
                speed = 30
                size = 1
                languages = append(languages, "Common", "Elvish")
                proficiencies = append(proficiencies, "perception")
                resistances = append(resistances, "sleep")
                darkvision = true
            case 2:
                dex += 2
                speed = 25
                size = 0
                languages = append(languages, "Common", "Halfling")
                darkvision = false
                lucky = true
                nimbleness = true
                brave = true
            case 3:
                str++
                dex++
                con++
                intel++
                wis++
                cha++
                hp++
                speed = 30
                size = 1
                languages = append(languages, "Common") // 1 more language
                darkvision = false
            case 4:
                str += 2
                cha++
                speed = 30
                size = 1
                // draconic ancestry
                languages = append(languages, "Common", "Draconic")
                darkvision = false
            case 5:
                intel += 2
                speed = 25
                size = 0
                languages = append(languages, "Common", "Gnomish")
                darkvision = true
            case 6:
                cha += 2
                // two abilities +1
                speed = 30
                size = 1
                // proficiency in 2 skills
                languages = append(languages, "Common", "Elvish") // +1 language
                resistances = append(resistances, "sleep")
                darkvision = true
            case 7:
                str += 2
                con++
                hp++
                speed = 30
                size = 1
                proficiencies = append(proficiencies, "intimidation")
                languages = append(languages, "Common", "Orc")
                darkvision = true
            case 8:
                intel++
                cha += 2
                speed = 30
                size = 1
                resistances = append(resistances, "fire")
                languages = append(languages, "Common", "Infernal")
                darkvision = true
            default:
                return errors.New("Invalid value for racesel")
            }
            creationsel = 0
            creation = false
            switch racesel {
            case 3, 4, 6:
                racechoices = true
            default:
                choices = true
            }
        }
    } else if racechoices {
        if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
            proficiencies = make([]string, 0)
            resistances = make([]string, 0)
            languages = make([]string, 0)
            darkvision = false
            lucky = false
            nimbleness = false
            creationsel = 0
            option0 = 0
            option1 = 0
            option2 = 0
            option3 = 0
            option4 = 0
            option5 = 0
            option6 = 0
            option7 = 0
            option8 = 0
            racechoices = false
            creation = true
        }
        switch racesel {
        case 3:
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                raceopt0--
                if raceopt0 < 0 {
                    raceopt0 = 0
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                raceopt0++
                if raceopt0 > 14 {
                    raceopt0 = 14
                }
            }
        case 4:
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                raceopt0--
                if raceopt0 < 0 {
                    raceopt0 = 0
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                raceopt0++
                if raceopt0 > 9 {
                    raceopt0 = 9
                }
            }
        case 6:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    raceopt0--
                    if raceopt0 < 0 {
                        raceopt0 = 0
                    }
                case 1:
                    raceopt1--
                    if raceopt1 < 0 {
                        raceopt1 = 0
                    }
                case 2:
                    raceopt2--
                    if raceopt2 < 0 {
                        raceopt2 = 0
                    }
                case 3:
                    raceopt3--
                    if raceopt3 < 0 {
                        raceopt3 = 0
                    }
                case 4:
                    raceopt4--
                    if raceopt4 < 0 {
                        raceopt4 = 0
                    }
                default:
                    return errors.New("Out of bounds (1019)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    raceopt0++
                    if raceopt0 > 5 {
                        raceopt0 = 5
                    }
                case 1:
                    raceopt1++
                    if raceopt1 > 5 {
                        raceopt1 = 5
                    }
                case 2:
                    raceopt2++
                    if raceopt2 > 17 {
                        raceopt2 = 17
                    }
                case 3:
                    raceopt3++
                    if raceopt3 > 17 {
                        raceopt3 = 17
                    }
                case 4:
                    raceopt4++
                    if raceopt4 > 13 {
                        raceopt4 = 13
                    }
                default:
                    return errors.New("Out of bounds (1055)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 4 {
                creationsel = 4
            }
            if raceopt0 == raceopt1 || raceopt2 == raceopt3 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        default:
            return errors.New("Invalid value for racesel")
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !dupwarning {
            switch racesel {
            case 3:
                switch raceopt0 {
                case 0:
                    languages = append(languages, "Dwarvish")
                case 1:
                    languages = append(languages, "Elvish")
                case 2:
                    languages = append(languages, "Giant")
                case 3:
                    languages = append(languages, "Gnomish")
                case 4:
                    languages = append(languages, "Goblin")
                case 5:
                    languages = append(languages, "Halfling")
                case 6:
                    languages = append(languages, "Orc")
                case 7:
                    languages = append(languages, "Abyssal")
                case 8:
                    languages = append(languages, "Celestial")
                case 9:
                    languages = append(languages, "Draconic")
                case 10:
                    languages = append(languages, "Deep Speech")
                case 11:
                    languages = append(languages, "Infernal")
                case 12:
                    languages = append(languages, "Primordial")
                case 13:
                    languages = append(languages, "Sylvan")
                case 14:
                    languages = append(languages, "Undercommon")
                default:
                    return errors.New("Invalid value for raceopt0 (case 3)")
                }
            case 4:
                switch raceopt0 {
                case 0:
                    ancestry = "Black"
                case 1:
                    ancestry = "Blue"
                case 2:
                    ancestry = "Brass"
                case 3:
                    ancestry = "Bronze"
                case 4:
                    ancestry = "Copper"
                case 5:
                    ancestry = "Gold"
                case 6:
                    ancestry = "Green"
                case 7:
                    ancestry = "Red"
                case 8:
                    ancestry = "Silver"
                case 9:
                    ancestry = "White"
                default:
                    return errors.New("Invalid value for raceopt0 (case 4)")
                }
            case 6:
                switch raceopt0 {
                case 0:
                    str++
                case 1:
                    dex++
                case 2:
                    con++
                case 3:
                    intel++
                case 4:
                    wis++
                case 5:
                    cha++
                default:
                    return errors.New("Invalid value for raceopt0 (case 6)")
                }
                switch raceopt1 {
                case 0:
                    str++
                case 1:
                    dex++
                case 2:
                    con++
                case 3:
                    intel++
                case 4:
                    wis++
                case 5:
                    cha++
                default:
                    return errors.New("Invalid value for raceopt1 (case 6)")
                }
                switch raceopt2 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "arcana")
                case 3:
                    proficiencies = append(proficiencies, "athletics")
                case 4:
                    proficiencies = append(proficiencies, "deception")
                case 5:
                    proficiencies = append(proficiencies, "history")
                case 6:
                    proficiencies = append(proficiencies, "insight")
                case 7:
                    proficiencies = append(proficiencies, "intimidation")
                case 8:
                    proficiencies = append(proficiencies, "investigation")
                case 9:
                    proficiencies = append(proficiencies, "medicine")
                case 10:
                    proficiencies = append(proficiencies, "nature")
                case 11:
                    proficiencies = append(proficiencies, "perception")
                case 12:
                    proficiencies = append(proficiencies, "performance")
                case 13:
                    proficiencies = append(proficiencies, "persuasion")
                case 14:
                    proficiencies = append(proficiencies, "religion")
                case 15:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 16:
                    proficiencies = append(proficiencies, "stealth")
                case 17:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for raceopt2 (case 6)")
                }
                switch raceopt3 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "arcana")
                case 3:
                    proficiencies = append(proficiencies, "athletics")
                case 4:
                    proficiencies = append(proficiencies, "deception")
                case 5:
                    proficiencies = append(proficiencies, "history")
                case 6:
                    proficiencies = append(proficiencies, "insight")
                case 7:
                    proficiencies = append(proficiencies, "intimidation")
                case 8:
                    proficiencies = append(proficiencies, "investigation")
                case 9:
                    proficiencies = append(proficiencies, "medicine")
                case 10:
                    proficiencies = append(proficiencies, "nature")
                case 11:
                    proficiencies = append(proficiencies, "perception")
                case 12:
                    proficiencies = append(proficiencies, "performance")
                case 13:
                    proficiencies = append(proficiencies, "persuasion")
                case 14:
                    proficiencies = append(proficiencies, "religion")
                case 15:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 16:
                    proficiencies = append(proficiencies, "stealth")
                case 17:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for raceopt3 (case 6)")
                }
                switch raceopt4 {
                case 0:
                    languages = append(languages, "Dwarvish")
                case 1:
                    languages = append(languages, "Undercommon")
                case 2:
                    languages = append(languages, "Giant")
                case 3:
                    languages = append(languages, "Gnomish")
                case 4:
                    languages = append(languages, "Goblin")
                case 5:
                    languages = append(languages, "Halfling")
                case 6:
                    languages = append(languages, "Orc")
                case 7:
                    languages = append(languages, "Abyssal")
                case 8:
                    languages = append(languages, "Celestial")
                case 9:
                    languages = append(languages, "Draconic")
                case 10:
                    languages = append(languages, "Deep Speech")
                case 11:
                    languages = append(languages, "Infernal")
                case 12:
                    languages = append(languages, "Primordial")
                case 13:
                    languages = append(languages, "Sylvan")
                default:
                    return errors.New("Invalid value for raceopt4 (case 6)")
                }
            default:
                return errors.New("Invalid value for racesel")
            }
            racechoices = false
            choices = true
            return nil
        }
    } else if choices {
        if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
            proficiencies = make([]string, 0)
            resistances = make([]string, 0)
            languages = make([]string, 0)
            darkvision = false
            lucky = false
            nimbleness = false
            creationsel = 0
            option0 = 0
            option1 = 0
            option2 = 0
            option3 = 0
            option4 = 0
            option5 = 0
            option6 = 0
            option7 = 0
            option8 = 0
            choices = false
            creation = true
        }
        switch classsel {
        case 0:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                default:
                    return errors.New("Out of bounds (973)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 5 {
                        option0 = 5
                    }
                case 1:
                    option1++
                    if option1 > 5 {
                        option1 = 5
                    }
                case 2:
                    option2++
                    if option2 > 17 {
                        option2 = 17
                    }
                case 3:
                    option3++
                    if option3 > 13 {
                        option3 = 13
                    }
                default:
                    return errors.New("Out of bounds (999)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 3 {
                creationsel = 3
            }
            if option0 == option1 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 1:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 {
                        option5 = 0
                    }
                case 6:
                    option6--
                    if option6 < 0 {
                        option6 = 0
                    }
                case 7:
                    option7--
                    if option7 < 0 {
                        option7 = 0
                    }
                case 8:
                    option8--
                    if option8 < 0 {
                        option8 = 0
                    }
                default:
                    return errors.New("Out of bounds (1067)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 9 {
                        option0 = 9
                    }
                case 1:
                    option1++
                    if option1 > 9 {
                        option1 = 9
                    }
                case 2:
                    option2++
                    if option2 > 9 {
                        option2 = 9
                    }
                case 3:
                    option3++
                    if option3 > 17 {
                        option3 = 17
                    }
                case 4:
                    option4++
                    if option4 > 17 {
                        option4 = 17
                    }
                case 5:
                    option5++
                    if option5 > 17 {
                        option5 = 17
                    }
                case 6:
                    option6++
                    if option6 > 15 {
                        option6 = 15
                    }
                case 7:
                    option7++
                    if option7 > 1 {
                        option7 = 1
                    }
                case 8:
                    option8++
                    if option8 > 9 {
                        option8 = 9
                    }
                default:
                    return errors.New("Out of bounds (1118)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 8 {
                creationsel = 8
            }
            if option0 == option1 || option1 == option2 || option0 == option2 || option3 == option4 || option4 == option5 || option3 == option5 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 2:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 {
                        option5 = 0
                    }
                default:
                    return errors.New("Out of bounds (1171)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 4 {
                        option0 = 4
                    }
                case 1:
                    option1++
                    if option1 > 4 {
                        option1 = 4
                    }
                case 2:
                    option2++
                    for _, prof := range proficiencies {
                        fmt.Println(prof)
                        if prof == "warhammer" || prof == "martial weapons" {
                            if option2 > 1 {
                                option2 = 1
                            }
                            return nil
                        }
//                        } else {
//                            if option2 > 0 {
//                                option2 = 0
//                            }
//                        }
                    }
                    if option2 > 0 {
                        option2 = 0
                    }
                case 3:
                    option3++
                    for _, prof := range proficiencies {
                        if prof == "heavy armor" || prof == "all armor" {
                            if option3 > 2 {
                                option3 = 2
                            }
                            return nil
                        }
//                        } else {
//                            if option3 > 1 {
//                                option3 = 1
//                            }
//                        }
                    }
                    if option3 > 1 {
                        option3 = 1
                    }
                case 4:
                    option4++
                    if option4 > 13 {
                        option4 = 13
                    }
                case 5:
                    option5++
                    if option5 > 1 {
                        option5 = 1
                    }
                default:
                    return errors.New("Out of bounds (1223)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 5 {
                creationsel = 5
            }
            if option0 == option1 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 3:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1271)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 7 {
                        option0 = 7
                    }
                case 1:
                    option1++
                    if option1 > 7 {
                        option1 = 7
                    }
                case 2:
                    option2++
                    if option2 > 14 {
                        option2 = 14
                    }
                case 3:
                    option3++
                    if option3 > 10 {
                        option3 = 10
                    }
                default:
                    log.Fatal("Out of bounds (1297)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 3 {
                creationsel = 3
            }
            if option0 == option1 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 4:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 {
                        option5 = 0
                    }
                case 6:
                    option6--
                    if option6 < 0 {
                        option6 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1350)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 7 {
                        option0 = 7
                    }
                case 1:
                    option1++
                    if option1 > 7 {
                        option1 = 7
                    }
                case 2:
                    option2++
                    if option2 > 1 {
                        option2 = 1
                    }
                case 3:
                    option3++
                    if option3 > 22 {
                        option3 = 22
                    }
                case 4:
                    option4++
                    if option4 > 23 {
                        option4 = 23
                    }
                case 5:
                    option5++
                    if option5 > 1 {
                        option5 = 1
                    }
                case 6:
                    option6++
                    if option6 > 1 {
                        option6 = 1
                    }
                default:
                    log.Fatal("Out of bounds (1391)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 6 {
                creationsel = 6
            }
            if option0 == option1 || option3 + 1 == option4 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 5:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1443)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 5 {
                        option0 = 5
                    }
                case 1:
                    option1++
                    if option1 > 5 {
                        option1 = 5
                    }
                case 2:
                    option2++
                    if option2 > 14 {
                        option2 = 14
                    }
                case 3:
                    option3++
                    if option3 > 1 {
                        option3 = 1
                    }
                default:
                    log.Fatal("Out of bounds (1469)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 3 {
                creationsel = 3
            }
            if option0 == option1 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 6:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 {
                        option5 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1517)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 5 {
                        option0 = 5
                    }
                case 1:
                    option1++
                    if option1 > 5 {
                        option1 = 5
                    }
                case 2:
                    option2++
                    if option2 > 22 {
                        option2 = 22
                    }
                case 3:
                    option3++
                    if option3 > 23 {
                        option3 = 23
                    }
                case 4:
                    option4++
                    if option4 > 10 {
                        option4 = 10
                    }
                case 5:
                    option5++
                    if option5 > 1 {
                        option5 = 1
                    }
                default:
                    log.Fatal("Out of bounds (1553)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 5 {
                creationsel = 5
            }
            if option0 == option1 || option2 + 1 == option3 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 7:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 == 0 {
                        option5 = 0
                    }
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 {
                        option5 = 0
                    }
                case 6:
                    option6--
                    if option6 < 0 {
                        option6 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1606)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 7 {
                        option0 = 7
                    }
                case 1:
                    option1++
                    if option1 > 7 {
                        option1 = 7
                    }
                case 2:
                    option2++
                    if option2 > 7 {
                        option2 = 7
                    }
                case 3:
                    option3++
                    if option3 > 1 {
                        option3 = 1
                    }
                case 4:
                    option4++
                    if option4 > 10 {
                        option4 = 10
                    }
                case 5:
                    option5++
                    if option4 == 0 {
                        if option5 > 0 {
                            option5 = 0
                        }
                    } else {
                        if option5 > 10 {
                            option5 = 10
                        }
                    }
                case 6:
                    option6++
                    if option6 > 1 {
                        option6 = 1
                    }
                default:
                    log.Fatal("Out of bounds (1653)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 6 {
                creationsel = 6
            }
            if option0 == option1 || option1 == option2 || option0 == option2 || (option4 == option5 && option4 != 0) {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 8:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 {
                        option5 = 0
                    }
                case 6:
                    option6--
                    if option6 < 0 {
                        option6 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1711)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 10 {
                        option0 = 10
                    }
                case 1:
                    option1++
                    if option1 > 10 {
                        option1 = 10
                    }
                case 2:
                    option2++
                    if option2 > 10 {
                        option2 = 10
                    }
                case 3:
                    option3++
                    if option3 > 10 {
                        option3 = 10
                    }
                case 4:
                    option4++
                    if option4 > 1 {
                        option4 = 1
                    }
                case 5:
                    option5++
                    if option5 > 1 {
                        option5 = 1
                    }
                case 6:
                    option6++
                    if option6 > 2 {
                        option6 = 2
                    }
                default:
                    log.Fatal("Out of bounds (1752)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 6 {
                creationsel = 6
            }
            if option0 == option1 || option0 == option2 || option0 == option3 || option1 == option2 || option1 == option3 || option2 == option3 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 9:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1800)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 5 {
                        option0 = 5
                    }
                case 1:
                    option1++
                    if option1 > 5 {
                        option1 = 5
                    }
                case 2:
                    option2++
                    if option2 > 13 {
                        option2 = 13
                    }
                case 3:
                    option3++
                    if option3 > 1 {
                        option3 = 1
                    }
                case 4:
                    option4++
                    if option4 > 1 {
                        option4 = 1
                    }
                default:
                    log.Fatal("Out of bounds (1831)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 4 {
                creationsel = 4
            }
            if option0 == option1 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 10:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1879)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 6 {
                        option0 = 6
                    }
                case 1:
                    option1++
                    if option1 > 6 {
                        option1 = 6
                    }
                case 2:
                    option2++
                    if option2 > 13 {
                        option2 = 13
                    }
                case 3:
                    option3++
                    if option3 > 1 {
                        option3 = 1
                    }
                case 4:
                    option4++
                    if option4 > 1 {
                        option4 = 1
                    }
                default:
                    log.Fatal("Out of bounds (1910)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 4 {
                creationsel = 4
            }
            if option0 == option1 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        case 11:
            if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                creationsel--
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                creationsel++
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                switch creationsel {
                case 0:
                    option0--
                    if option0 < 0 {
                        option0 = 0
                    }
                case 1:
                    option1--
                    if option1 < 0 {
                        option1 = 0
                    }
                case 2:
                    option2--
                    if option2 < 0 {
                        option2 = 0
                    }
                case 3:
                    option3--
                    if option3 < 0 {
                        option3 = 0
                    }
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                default:
                    log.Fatal("Out of bounds (1958)")
                }
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                switch creationsel {
                case 0:
                    option0++
                    if option0 > 5 {
                        option0 = 5
                    }
                case 1:
                    option1++
                    if option1 > 5 {
                        option1 = 5
                    }
                case 2:
                    option2++
                    if option2 > 1 {
                        option2 = 1
                    }
                case 3:
                    option3++
                    if option3 > 1 {
                        option3 = 1
                    }
                case 4:
                    option4++
                    if option4 > 1 {
                        option4 = 1
                    }
                default:
                    log.Fatal("Out of bounds (1989)")
                }
            }
            if creationsel < 0 {
                creationsel = 0
            } else if creationsel > 4 {
                creationsel = 4
            }
            if option0 == option1 {
                dupwarning = true
            } else {
                dupwarning = false
            }
        default:
            return errors.New("Invalid value for classsel")
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !dupwarning {
            switch classsel {
            case 0:
                ac = 10 + ((dex - 10) / 2) + ((con - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "animal handling")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "intimidation")
                case 3:
                    proficiencies = append(proficiencies, "nature")
                case 4:
                    proficiencies = append(proficiencies, "perception")
                case 5:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option0 (case 0)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "animal handling")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "intimidation")
                case 3:
                    proficiencies = append(proficiencies, "nature")
                case 4:
                    proficiencies = append(proficiencies, "perception")
                case 5:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option1 (case 0)")
                }
                switch option2 {
                case 0:
                    var greataxe items.Greataxe
                    err = p.Inv.Add(greataxe)
                    if err != nil {
                        return errors.New("Failed to add greataxe to inv")
                    }
                case 1:
                    var battleaxe items.Battleaxe
                    err = p.Inv.Add(battleaxe)
                    if err != nil {
                        return errors.New("Failed to add battleaxe to inv")
                    }
                case 2:
                    var flail items.Flail
                    err = p.Inv.Add(flail)
                    if err != nil {
                        return errors.New("Failed to add flail to inv")
                    }
                case 3:
                    var glaive items.Glaive
                    err = p.Inv.Add(glaive)
                    if err != nil {
                        return errors.New("Failed to add glaive to inv")
                    }
                case 4:
                    var greatsword items.Greatsword
                    err = p.Inv.Add(greatsword)
                    if err != nil {
                        return errors.New("Failed to add greatsword to inv")
                    }
                case 5:
                    var halberd items.Halberd
                    err = p.Inv.Add(halberd)
                    if err != nil {
                        return errors.New("Failed to add halberd to inv")
                    }
                case 6:
                    var lance items.Lance
                    err = p.Inv.Add(lance)
                    if err != nil {
                        return errors.New("Failed to add lance to inv")
                    }
                case 7:
                    var longsword items.Longsword
                    err = p.Inv.Add(longsword)
                    if err != nil {
                        return errors.New("Failed to add longsword to inv")
                    }
                case 8:
                    var maul items.Maul
                    err = p.Inv.Add(maul)
                    if err != nil {
                        return errors.New("Failed to add maul to inv")
                    }
                case 9:
                    var morningstar items.Morningstar
                    err = p.Inv.Add(morningstar)
                    if err != nil {
                        return errors.New("Failed to add morningstar to inv")
                    }
                case 10:
                    var pike items.Pike
                    err = p.Inv.Add(pike)
                    if err != nil {
                        return errors.New("Failed to add pike to inv")
                    }
                case 11:
                    var rapier items.Rapier
                    err = p.Inv.Add(rapier)
                    if err != nil {
                        return errors.New("Failed to add rapier to inv")
                    }
                case 12:
                    var scimitar items.Scimitar
                    err = p.Inv.Add(scimitar)
                    if err != nil {
                        return errors.New("Failed to add scimitar to inv")
                    }
                case 13:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 14:
                    var trident items.Trident
                    err = p.Inv.Add(trident)
                    if err != nil {
                        return errors.New("Failed to add trident to inv")
                    }
                case 15:
                    var warpick items.Warpick
                    err = p.Inv.Add(warpick)
                    if err != nil {
                        return errors.New("Failed to add warpick to inv")
                    }
                case 16:
                    var warhammer items.Warhammer
                    err = p.Inv.Add(warhammer)
                    if err != nil {
                        return errors.New("Failed to add warhammer to inv")
                    }
                case 17:
                    var whip items.Whip
                    err = p.Inv.Add(whip)
                    if err != nil {
                        return errors.New("Failed to add whip to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 0)")
                }
                switch option3 {
                case 0:
                    var handaxe0 items.Handaxe
                    var handaxe1 items.Handaxe
                    err = p.Inv.Add(handaxe0)
                    if err != nil {
                        return errors.New("Failed to add handaxe0 to inv")
                    }
                    err = p.Inv.Add(handaxe1)
                    if err != nil {
                        return errors.New("Failed to add handaxe1 to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 5:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 6:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 7:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 8:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 9:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                case 10:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 11:
                    var dart items.Dart
                    err = p.Inv.Add(dart)
                    if err != nil {
                        return errors.New("Failed to add dart to inv")
                    }
                case 12:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 13:
                    var sling items.Sling
                    err = p.Inv.Add(sling)
                    if err != nil {
                        return errors.New("Failed to add sling to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 0)")
                }
            case 1:
                ac = 11 + ((dex - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "bagpipes")
                case 1:
                    proficiencies = append(proficiencies, "drum")
                case 2:
                    proficiencies = append(proficiencies, "dulcimer")
                case 3:
                    proficiencies = append(proficiencies, "flute")
                case 4:
                    proficiencies = append(proficiencies, "lute")
                case 5:
                    proficiencies = append(proficiencies, "lyre")
                case 6:
                    proficiencies = append(proficiencies, "horn")
                case 7:
                    proficiencies = append(proficiencies, "pan flute")
                case 8:
                    proficiencies = append(proficiencies, "shawm")
                case 9:
                    proficiencies = append(proficiencies, "viol")
                default:
                    return errors.New("Invalid value for option0 (case 1)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "bagpipes")
                case 1:
                    proficiencies = append(proficiencies, "drum")
                case 2:
                    proficiencies = append(proficiencies, "dulcimer")
                case 3:
                    proficiencies = append(proficiencies, "flute")
                case 4:
                    proficiencies = append(proficiencies, "lute")
                case 5:
                    proficiencies = append(proficiencies, "lyre")
                case 6:
                    proficiencies = append(proficiencies, "horn")
                case 7:
                    proficiencies = append(proficiencies, "pan flute")
                case 8:
                    proficiencies = append(proficiencies, "shawm")
                case 9:
                    proficiencies = append(proficiencies, "viol")
                default:
                    return errors.New("Invalid value for option1 (case 1)")
                }
                switch option2 {
                case 0:
                    proficiencies = append(proficiencies, "bagpipes")
                case 1:
                    proficiencies = append(proficiencies, "drum")
                case 2:
                    proficiencies = append(proficiencies, "dulcimer")
                case 3:
                    proficiencies = append(proficiencies, "flute")
                case 4:
                    proficiencies = append(proficiencies, "lute")
                case 5:
                    proficiencies = append(proficiencies, "lyre")
                case 6:
                    proficiencies = append(proficiencies, "horn")
                case 7:
                    proficiencies = append(proficiencies, "pan flute")
                case 8:
                    proficiencies = append(proficiencies, "shawm")
                case 9:
                    proficiencies = append(proficiencies, "viol")
                default:
                    return errors.New("Invalid value for option2 (case 1)")
                }
                switch option3 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "arcana")
                case 3:
                    proficiencies = append(proficiencies, "athletics")
                case 4:
                    proficiencies = append(proficiencies, "deception")
                case 5:
                    proficiencies = append(proficiencies, "history")
                case 6:
                    proficiencies = append(proficiencies, "insight")
                case 7:
                    proficiencies = append(proficiencies, "intimidation")
                case 8:
                    proficiencies = append(proficiencies, "investigation")
                case 9:
                    proficiencies = append(proficiencies, "medicine")
                case 10:
                    proficiencies = append(proficiencies, "nature")
                case 11:
                    proficiencies = append(proficiencies, "perception")
                case 12:
                    proficiencies = append(proficiencies, "performance")
                case 13:
                    proficiencies = append(proficiencies, "persuasion")
                case 14:
                    proficiencies = append(proficiencies, "religion")
                case 15:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 16:
                    proficiencies = append(proficiencies, "stealth")
                case 17:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option3 (case 1)")
                }
                switch option4 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "arcana")
                case 3:
                    proficiencies = append(proficiencies, "athletics")
                case 4:
                    proficiencies = append(proficiencies, "deception")
                case 5:
                    proficiencies = append(proficiencies, "history")
                case 6:
                    proficiencies = append(proficiencies, "insight")
                case 7:
                    proficiencies = append(proficiencies, "intimidation")
                case 8:
                    proficiencies = append(proficiencies, "investigation")
                case 9:
                    proficiencies = append(proficiencies, "medicine")
                case 10:
                    proficiencies = append(proficiencies, "nature")
                case 11:
                    proficiencies = append(proficiencies, "perception")
                case 12:
                    proficiencies = append(proficiencies, "performance")
                case 13:
                    proficiencies = append(proficiencies, "persuasion")
                case 14:
                    proficiencies = append(proficiencies, "religion")
                case 15:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 16:
                    proficiencies = append(proficiencies, "stealth")
                case 17:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option4 (case 1)")
                }
                switch option5 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "arcana")
                case 3:
                    proficiencies = append(proficiencies, "athletics")
                case 4:
                    proficiencies = append(proficiencies, "deception")
                case 5:
                    proficiencies = append(proficiencies, "history")
                case 6:
                    proficiencies = append(proficiencies, "insight")
                case 7:
                    proficiencies = append(proficiencies, "intimidation")
                case 8:
                    proficiencies = append(proficiencies, "investigation")
                case 9:
                    proficiencies = append(proficiencies, "medicine")
                case 10:
                    proficiencies = append(proficiencies, "nature")
                case 11:
                    proficiencies = append(proficiencies, "perception")
                case 12:
                    proficiencies = append(proficiencies, "performance")
                case 13:
                    proficiencies = append(proficiencies, "persuasion")
                case 14:
                    proficiencies = append(proficiencies, "religion")
                case 15:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 16:
                    proficiencies = append(proficiencies, "stealth")
                case 17:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option5 (case 1)")
                }
                switch option6 {
                case 0:
                    var rapier items.Rapier
                    err = p.Inv.Add(rapier)
                    if err != nil {
                        return errors.New("Failed to add rapier to inv")
                    }
                case 1:
                    var longsword items.Longsword
                    err = p.Inv.Add(longsword)
                    if err != nil {
                        return errors.New("Failed to add longsword to inv")
                    }
                case 2:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 3:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 4:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 5:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 6:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 7:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 8:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 9:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 10:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 11:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                case 12:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 13:
                    var dart items.Dart
                    err = p.Inv.Add(dart)
                    if err != nil {
                        return errors.New("Failed to add dart to inv")
                    }
                case 14:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 15:
                    var sling items.Sling
                    err = p.Inv.Add(sling)
                    if err != nil {
                        return errors.New("Failed to add sling to inv")
                    }
                default:
                    return errors.New("Invalid value for option6 (case 1)")
                }
                switch option7 {
                case 0:
                    var clothes = items.Clothes{Quality: "Fine"}
                    var inkbottle items.Inkbottle
                    var inkpen items.Inkpen
                    var lamp items.Lamp
                    var oilflask = items.Oilflask{Quantity: 2}
                    var paper = items.Paper{Quantity: 5}
                    var perfume items.Perfume
                    var sealingwax items.Sealingwax
                    var soap items.Soap
                    err = p.Inv.Add(clothes)
                    if err != nil {
                        return errors.New("Failed to add clothes to inv")
                    }
                    err = p.Inv.Add(inkbottle)
                    if err != nil {
                        return errors.New("Failed to add inkbottle to inv")
                    }
                    err = p.Inv.Add(inkpen)
                    if err != nil {
                        return errors.New("Failed to add inkpen to inv")
                    }
                    err = p.Inv.Add(lamp)
                    if err != nil {
                        return errors.New("Failed to add lamp to inv")
                    }
                    err = p.Inv.Add(oilflask)
                    if err != nil {
                        return errors.New("Failed to add oilflask to inv")
                    }
                    err = p.Inv.Add(paper)
                    if err != nil {
                        return errors.New("Failed to add paper to inv")
                    }
                    err = p.Inv.Add(perfume)
                    if err != nil {
                        return errors.New("Failed to add perfume to inv")
                    }
                    err = p.Inv.Add(sealingwax)
                    if err != nil {
                        return errors.New("Failed to add sealingwax to inv")
                    }
                    err = p.Inv.Add(soap)
                    if err != nil {
                        return errors.New("Failed to add soap to inv")
                    }
                case 1:
                    var clothes = items.Clothes{Quality: "Costume"}
                    var candles = items.Candles{Quantity: 5}
                    var disguisekit items.Disguisekit
                    err = p.Inv.Add(clothes)
                    if err != nil {
                        return errors.New("Failed to add clothes to inv")
                    }
                    err = p.Inv.Add(candles)
                    if err != nil {
                        return errors.New("Failed to add candles to inv")
                    }
                    err = p.Inv.Add(disguisekit)
                    if err != nil {
                        return errors.New("Failed to add disguisekit to inv")
                    }
                default:
                    return errors.New("Invalid value for option7 (case 1)")
                }
                switch option8 {
                case 0:
                    var bagpipes items.Bagpipes
                    err = p.Inv.Add(bagpipes)
                    if err != nil {
                        return errors.New("Failed to add bagpipes to inv")
                    }
                case 1:
                    var drum items.Drum
                    err = p.Inv.Add(drum)
                    if err != nil {
                        return errors.New("Failed to add drum to inv")
                    }
                case 2:
                    var dulcimer items.Dulcimer
                    err = p.Inv.Add(dulcimer)
                    if err != nil {
                        return errors.New("Failed to add dulcimer to inv")
                    }
                case 3:
                    var flute items.Flute
                    err = p.Inv.Add(flute)
                    if err != nil {
                        return errors.New("Failed to add flute to inv")
                    }
                case 4:
                    var lute items.Lute
                    err = p.Inv.Add(lute)
                    if err != nil {
                        return errors.New("Failed to add lute to inv")
                    }
                case 5:
                    var lyre items.Lyre
                    err = p.Inv.Add(lyre)
                    if err != nil {
                        return errors.New("Failed to add lyre to inv")
                    }
                case 6:
                    var horn items.Horn
                    err = p.Inv.Add(horn)
                    if err != nil {
                        return errors.New("Failed to add horn to inv")
                    }
                case 7:
                    var panflute items.Panflute
                    err = p.Inv.Add(panflute)
                    if err != nil {
                        return errors.New("Failed to add panflute to inv")
                    }
                case 8:
                    var shawm items.Shawm
                    err = p.Inv.Add(shawm)
                    if err != nil {
                        return errors.New("Failed to add shawm to inv")
                    }
                case 9:
                    var viol items.Viol
                    err = p.Inv.Add(viol)
                    if err != nil {
                        return errors.New("Failed to add viol to inv")
                    }
                default:
                    return errors.New("Invalid value for option8 (case 1)")
                }
            case 2:
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "history")
                case 1:
                    proficiencies = append(proficiencies, "insight")
                case 2:
                    proficiencies = append(proficiencies, "medicine")
                case 3:
                    proficiencies = append(proficiencies, "persuasion")
                case 4:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option0 (case 2)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "history")
                case 1:
                    proficiencies = append(proficiencies, "insight")
                case 2:
                    proficiencies = append(proficiencies, "medicine")
                case 3:
                    proficiencies = append(proficiencies, "persuasion")
                case 4:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option1 (case 2)")
                }
                switch option2 {
                case 0:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 1:
                    var warhammer items.Warhammer
                    err = p.Inv.Add(warhammer)
                    if err != nil {
                        return errors.New("Failed to add warhammer to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 2)")
                }
                switch option3 {
                case 0:
                    var scalemail items.Scalemail
                    err = p.Inv.Add(scalemail)
                    if err != nil {
                        return errors.New("Failed to add scalemail to inv")
                    }
                    if ((dex - 10) / 2) > 2 {
                        ac = 18
                    } else {
                        ac = 16 + ((dex - 10) / 2)
                    }
                case 1:
                    var leatherarmor items.Leatherarmor
                    err = p.Inv.Add(leatherarmor)
                    if err != nil {
                        return errors.New("Failed to add leatherarmor to inv")
                    }
                    ac = 13 + ((dex - 10) / 2)
                case 2:
                    var chainmail items.Chainmail
                    err = p.Inv.Add(chainmail)
                    if err != nil {
                        return errors.New("Failed to add chainmail to inv")
                    }
                    ac = 18
                default:
                    return errors.New("Invalid value for option3 (case 2)")
                }
                switch option4 {
                case 0:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                case 11:
                    var dart items.Dart
                    err = p.Inv.Add(dart)
                    if err != nil {
                        return errors.New("Failed to add dart to inv")
                    }
                case 12:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 13:
                    var sling items.Sling
                    err = p.Inv.Add(sling)
                    if err != nil {
                        return errors.New("Failed to add sling to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 2)")
                }
                switch option5 {
                case 0:
                    var candles = items.Candles{Quantity: 10}
                    var tinderbox items.Tinderbox
                    err = p.Inv.Add(candles)
                    if err != nil {
                        return errors.New("Failed to add candles to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                case 1:
                    var tinderbox items.Tinderbox
                    var torches = items.Torches{Quantity: 10}
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option5 (case 2)")
                }
            case 3:
                ac = 11 + ((dex - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "medicine")
                case 4:
                    proficiencies = append(proficiencies, "nature")
                case 5:
                    proficiencies = append(proficiencies, "perception")
                case 6:
                    proficiencies = append(proficiencies, "religion")
                case 7:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option0 (case 3)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "medicine")
                case 4:
                    proficiencies = append(proficiencies, "nature")
                case 5:
                    proficiencies = append(proficiencies, "perception")
                case 6:
                    proficiencies = append(proficiencies, "religion")
                case 7:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option1 (case 3)")
                }
                switch option2 {
                case 0:
                    var shield items.Shield
                    err = p.Inv.Add(shield)
                    if err != nil {
                        return errors.New("Failed to add shield to inv")
                    }
                    ac += 2
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                case 11:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 12:
                    var dart items.Dart
                    err = p.Inv.Add(dart)
                    if err != nil {
                        return errors.New("Failed to add dart to inv")
                    }
                case 13:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 14:
                    var sling items.Sling
                    err = p.Inv.Add(sling)
                    if err != nil {
                        return errors.New("Failed to add sling to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 3)")
                }
                switch option3 {
                case 0:
                    var scimitar items.Scimitar
                    err = p.Inv.Add(scimitar)
                    if err != nil {
                        return errors.New("Failed to add scimitar to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 3)")
                }
            case 4:
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "athletics")
                case 3:
                    proficiencies = append(proficiencies, "history")
                case 4:
                    proficiencies = append(proficiencies, "insight")
                case 5:
                    proficiencies = append(proficiencies, "intimidation")
                case 6:
                    proficiencies = append(proficiencies, "perception")
                case 7:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option0 (case 4)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "animal handling")
                case 2:
                    proficiencies = append(proficiencies, "athletics")
                case 3:
                    proficiencies = append(proficiencies, "history")
                case 4:
                    proficiencies = append(proficiencies, "insight")
                case 5:
                    proficiencies = append(proficiencies, "intimidation")
                case 6:
                    proficiencies = append(proficiencies, "perception")
                case 7:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option1 (case 4)")
                }
                switch option2 {
                case 0:
                    var chainmail items.Chainmail
                    err = p.Inv.Add(chainmail)
                    if err != nil {
                        return errors.New("Failed to add chainmail to inv")
                    }
                    ac = 16
                case 1:
                    var leatherarmor items.Leatherarmor
                    var longbow items.Longbow
                    err = p.Inv.Add(leatherarmor)
                    if err != nil {
                        return errors.New("Failed to add leatherarmor to inv")
                    }
                    err = p.Inv.Add(longbow)
                    if err != nil {
                        return errors.New("Failed to add longbow to inv")
                    }
                    ac = 11 + ((dex - 10) / 2)
                default:
                    return errors.New("Invalid value for option2 (case 4)")
                }
                switch option3 {
                case 0:
                    var battleaxe items.Battleaxe
                    err = p.Inv.Add(battleaxe)
                    if err != nil {
                        return errors.New("Failed to add battleaxe to inv")
                    }
                case 1:
                    var flail items.Flail
                    err = p.Inv.Add(flail)
                    if err != nil {
                        return errors.New("Failed to add flail to inv")
                    }
                case 2:
                    var glaive items.Glaive
                    err = p.Inv.Add(glaive)
                    if err != nil {
                        return errors.New("Failed to add glaive to inv")
                    }
                case 3:
                    var greataxe items.Greataxe
                    err = p.Inv.Add(greataxe)
                    if err != nil {
                        return errors.New("Failed to add greataxe to inv")
                    }
                case 4:
                    var greatsword items.Greatsword
                    err = p.Inv.Add(greatsword)
                    if err != nil {
                        return errors.New("Failed to add greatsword to inv")
                    }
                case 5:
                    var halberd items.Halberd
                    err = p.Inv.Add(halberd)
                    if err != nil {
                        return errors.New("Failed to add halberd to inv")
                    }
                case 6:
                    var lance items.Lance
                    err = p.Inv.Add(lance)
                    if err != nil {
                        return errors.New("Failed to add lance to inv")
                    }
                case 7:
                    var longsword items.Longsword
                    err = p.Inv.Add(longsword)
                    if err != nil {
                        return errors.New("Failed to add longsword to inv")
                    }
                case 8:
                    var maul items.Maul
                    err = p.Inv.Add(maul)
                    if err != nil {
                        return errors.New("Failed to add maul to inv")
                    }
                case 9:
                    var morningstar items.Morningstar
                    err = p.Inv.Add(morningstar)
                    if err != nil {
                        return errors.New("Failed to add morningstar to inv")
                    }
                case 10:
                    var pike items.Pike
                    err = p.Inv.Add(pike)
                    if err != nil {
                        return errors.New("Failed to add pike to inv")
                    }
                case 11:
                    var rapier items.Rapier
                    err = p.Inv.Add(rapier)
                    if err != nil {
                        return errors.New("Failed to add rapier to inv")
                    }
                case 12:
                    var scimitar items.Scimitar
                    err = p.Inv.Add(scimitar)
                    if err != nil {
                        return errors.New("Failed to add scimitar to inv")
                    }
                case 13:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 14:
                    var trident items.Trident
                    err = p.Inv.Add(trident)
                    if err != nil {
                        return errors.New("Failed to add trident to inv")
                    }
                case 15:
                    var warpick items.Warpick
                    err = p.Inv.Add(warpick)
                    if err != nil {
                        return errors.New("Failed to add warpick to inv")
                    }
                case 16:
                    var warhammer items.Warhammer
                    err = p.Inv.Add(warhammer)
                    if err != nil {
                        return errors.New("Failed to add warhammer to inv")
                    }
                case 17:
                    var whip items.Whip
                    err = p.Inv.Add(whip)
                    if err != nil {
                        return errors.New("Failed to add whip to inv")
                    }
                case 18:
                    var blowgun items.Blowgun
                    err = p.Inv.Add(blowgun)
                    if err != nil {
                        return errors.New("Failed to add blowgun to inv")
                    }
                case 19:
                    var handcrossbow items.Handcrossbow
                    err = p.Inv.Add(handcrossbow)
                    if err != nil {
                        return errors.New("Failed to add handcrossbow to inv")
                    }
                case 20:
                    var heavycrossbow items.Heavycrossbow
                    err = p.Inv.Add(heavycrossbow)
                    if err != nil {
                        return errors.New("Failed to add heavycrossbow to inv")
                    }
                case 21:
                    var longbow items.Longbow
                    err = p.Inv.Add(longbow)
                    if err != nil {
                        return errors.New("Failed to add longbow to inv")
                    }
                case 22:
                    var net items.Net
                    err = p.Inv.Add(net)
                    if err != nil {
                        return errors.New("Failed to add net to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 4)")
                }
                switch option4 {
                case 0:
                    var shield items.Shield
                    err = p.Inv.Add(shield)
                    if err != nil {
                        return errors.New("Failed to add shield to inv")
                    }
                    ac += 2
                case 1:
                    var battleaxe items.Battleaxe
                    err = p.Inv.Add(battleaxe)
                    if err != nil {
                        return errors.New("Failed to add battleaxe to inv")
                    }
                case 2:
                    var flail items.Flail
                    err = p.Inv.Add(flail)
                    if err != nil {
                        return errors.New("Failed to add flail to inv")
                    }
                case 3:
                    var glaive items.Glaive
                    err = p.Inv.Add(glaive)
                    if err != nil {
                        return errors.New("Failed to add glaive to inv")
                    }
                case 4:
                    var greataxe items.Greataxe
                    err = p.Inv.Add(greataxe)
                    if err != nil {
                        return errors.New("Failed to add greataxe to inv")
                    }
                case 5:
                    var greatsword items.Greatsword
                    err = p.Inv.Add(greatsword)
                    if err != nil {
                        return errors.New("Failed to add greatsword to inv")
                    }
                case 6:
                    var halberd items.Halberd
                    err = p.Inv.Add(halberd)
                    if err != nil {
                        return errors.New("Failed to add halberd to inv")
                    }
                case 7:
                    var lance items.Lance
                    err = p.Inv.Add(lance)
                    if err != nil {
                        return errors.New("Failed to add lance to inv")
                    }
                case 8:
                    var longsword items.Longsword
                    err = p.Inv.Add(longsword)
                    if err != nil {
                        return errors.New("Failed to add longsword to inv")
                    }
                case 9:
                    var maul items.Maul
                    err = p.Inv.Add(maul)
                    if err != nil {
                        return errors.New("Failed to add maul to inv")
                    }
                case 10:
                    var morningstar items.Morningstar
                    err = p.Inv.Add(morningstar)
                    if err != nil {
                        return errors.New("Failed to add morningstar to inv")
                    }
                case 11:
                    var pike items.Pike
                    err = p.Inv.Add(pike)
                    if err != nil {
                        return errors.New("Failed to add pike to inv")
                    }
                case 12:
                    var rapier items.Rapier
                    err = p.Inv.Add(rapier)
                    if err != nil {
                        return errors.New("Failed to add rapier to inv")
                    }
                case 13:
                    var scimitar items.Scimitar
                    err = p.Inv.Add(scimitar)
                    if err != nil {
                        return errors.New("Failed to add scimitar to inv")
                    }
                case 14:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 15:
                    var trident items.Trident
                    err = p.Inv.Add(trident)
                    if err != nil {
                        return errors.New("Failed to add trident to inv")
                    }
                case 16:
                    var warpick items.Warpick
                    err = p.Inv.Add(warpick)
                    if err != nil {
                        return errors.New("Failed to add warpick to inv")
                    }
                case 17:
                    var warhammer items.Warhammer
                    err = p.Inv.Add(warhammer)
                    if err != nil {
                        return errors.New("Failed to add warhammer to inv")
                    }
                case 18:
                    var whip items.Whip
                    err = p.Inv.Add(whip)
                    if err != nil {
                        return errors.New("Failed to add whip to inv")
                    }
                case 19:
                    var blowgun items.Blowgun
                    err = p.Inv.Add(blowgun)
                    if err != nil {
                        return errors.New("Failed to add blowgun to inv")
                    }
                case 20:
                    var handcrossbow items.Handcrossbow
                    err = p.Inv.Add(handcrossbow)
                    if err != nil {
                        return errors.New("Failed to add handcrossbow to inv")
                    }
                case 21:
                    var heavycrossbow items.Heavycrossbow
                    err = p.Inv.Add(heavycrossbow)
                    if err != nil {
                        return errors.New("Failed to add heavycrossbow to inv")
                    }
                case 22:
                    var longbow items.Longbow
                    err = p.Inv.Add(longbow)
                    if err != nil {
                        return errors.New("Failed to add longbow to inv")
                    }
                case 23:
                    var net items.Net
                    err = p.Inv.Add(net)
                    if err != nil {
                        return errors.New("Failed to add net to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 4)")
                }
                switch option5 {
                case 0:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 1:
                    var handaxe0 items.Handaxe
                    var handaxe1 items.Handaxe
                    err = p.Inv.Add(handaxe0)
                    if err != nil {
                        return errors.New("Failed to add handaxe0 to inv")
                    }
                    err = p.Inv.Add(handaxe1)
                    if err != nil {
                        return errors.New("Failed to add handaxe1 to inv")
                    }
                default:
                    return errors.New("Invalid value for option5 (case 4)")
                }
                switch option6 {
                case 0:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                case 1:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option6 (case 4)")
                }
            case 5:
                ac = 10 + ((dex - 10) / 2) + ((wis - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "history")
                case 3:
                    proficiencies = append(proficiencies, "insight")
                case 4:
                    proficiencies = append(proficiencies, "religion")
                case 5:
                    proficiencies = append(proficiencies, "stealth")
                default:
                    return errors.New("Invalid value for option0 (case 5)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "history")
                case 3:
                    proficiencies = append(proficiencies, "insight")
                case 4:
                    proficiencies = append(proficiencies, "religion")
                case 5:
                    proficiencies = append(proficiencies, "stealth")
                default:
                    return errors.New("Invalid value for option1 (case 5)")
                }
                switch option2 {
                case 0:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                case 11:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 12:
                    var dart items.Dart
                    err = p.Inv.Add(dart)
                    if err != nil {
                        return errors.New("Failed to add dart to inv")
                    }
                case 13:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 14:
                    var sling items.Sling
                    err = p.Inv.Add(sling)
                    if err != nil {
                        return errors.New("Failed to add sling to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 5)")
                }
                switch option3 {
                case 0:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                case 1:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 5)")
                }
            case 6:
                ac = 16
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "athletics")
                case 1:
                    proficiencies = append(proficiencies, "insight")
                case 2:
                    proficiencies = append(proficiencies, "intimidation")
                case 3:
                    proficiencies = append(proficiencies, "medicine")
                case 4:
                    proficiencies = append(proficiencies, "persuasion")
                case 5:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option0 (case 6)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "athletics")
                case 1:
                    proficiencies = append(proficiencies, "insight")
                case 2:
                    proficiencies = append(proficiencies, "intimidation")
                case 3:
                    proficiencies = append(proficiencies, "medicine")
                case 4:
                    proficiencies = append(proficiencies, "persuasion")
                case 5:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option1 (case 6)")
                }
                switch option2 {
                case 0:
                    var battleaxe items.Battleaxe
                    err = p.Inv.Add(battleaxe)
                    if err != nil {
                        return errors.New("Failed to add battleaxe to inv")
                    }
                case 1:
                    var flail items.Flail
                    err = p.Inv.Add(flail)
                    if err != nil {
                        return errors.New("Failed to add flail to inv")
                    }
                case 2:
                    var glaive items.Glaive
                    err = p.Inv.Add(glaive)
                    if err != nil {
                        return errors.New("Failed to add glaive to inv")
                    }
                case 3:
                    var greataxe items.Greataxe
                    err = p.Inv.Add(greataxe)
                    if err != nil {
                        return errors.New("Failed to add greataxe to inv")
                    }
                case 4:
                    var greatsword items.Greatsword
                    err = p.Inv.Add(greatsword)
                    if err != nil {
                        return errors.New("Failed to add greatsword to inv")
                    }
                case 5:
                    var halberd items.Halberd
                    err = p.Inv.Add(halberd)
                    if err != nil {
                        return errors.New("Failed to add halberd to inv")
                    }
                case 6:
                    var lance items.Lance
                    err = p.Inv.Add(lance)
                    if err != nil {
                        return errors.New("Failed to add lance to inv")
                    }
                case 7:
                    var longsword items.Longsword
                    err = p.Inv.Add(longsword)
                    if err != nil {
                        return errors.New("Failed to add longsword to inv")
                    }
                case 8:
                    var maul items.Maul
                    err = p.Inv.Add(maul)
                    if err != nil {
                        return errors.New("Failed to add maul to inv")
                    }
                case 9:
                    var morningstar items.Morningstar
                    err = p.Inv.Add(morningstar)
                    if err != nil {
                        return errors.New("Failed to add morningstar to inv")
                    }
                case 10:
                    var pike items.Pike
                    err = p.Inv.Add(pike)
                    if err != nil {
                        return errors.New("Failed to add pike to inv")
                    }
                case 11:
                    var rapier items.Rapier
                    err = p.Inv.Add(rapier)
                    if err != nil {
                        return errors.New("Failed to add rapier to inv")
                    }
                case 12:
                    var scimitar items.Scimitar
                    err = p.Inv.Add(scimitar)
                    if err != nil {
                        return errors.New("Failed to add scimitar to inv")
                    }
                case 13:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 14:
                    var trident items.Trident
                    err = p.Inv.Add(trident)
                    if err != nil {
                        return errors.New("Failed to add trident to inv")
                    }
                case 15:
                    var warpick items.Warpick
                    err = p.Inv.Add(warpick)
                    if err != nil {
                        return errors.New("Failed to add warpick to inv")
                    }
                case 16:
                    var warhammer items.Warhammer
                    err = p.Inv.Add(warhammer)
                    if err != nil {
                        return errors.New("Failed to add warhammer to inv")
                    }
                case 17:
                    var whip items.Whip
                    err = p.Inv.Add(whip)
                    if err != nil {
                        return errors.New("Failed to add whip to inv")
                    }
                case 18:
                    var blowgun items.Blowgun
                    err = p.Inv.Add(blowgun)
                    if err != nil {
                        return errors.New("Failed to add blowgun to inv")
                    }
                case 19:
                    var handcrossbow items.Handcrossbow
                    err = p.Inv.Add(handcrossbow)
                    if err != nil {
                        return errors.New("Failed to add handcrossbow to inv")
                    }
                case 20:
                    var heavycrossbow items.Heavycrossbow
                    err = p.Inv.Add(heavycrossbow)
                    if err != nil {
                        return errors.New("Failed to add heavycrossbow to inv")
                    }
                case 21:
                    var longbow items.Longbow
                    err = p.Inv.Add(longbow)
                    if err != nil {
                        return errors.New("Failed to add longbow to inv")
                    }
                case 22:
                    var net items.Net
                    err = p.Inv.Add(net)
                    if err != nil {
                        return errors.New("Failed to add net to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 6)")
                }
                switch option3 {
                case 0:
                    var shield items.Shield
                    err = p.Inv.Add(shield)
                    if err != nil {
                        return errors.New("Failed to add shield to inv")
                    }
                    ac += 2
                case 1:
                    var battleaxe items.Battleaxe
                    err = p.Inv.Add(battleaxe)
                    if err != nil {
                        return errors.New("Failed to add battleaxe to inv")
                    }
                case 2:
                    var flail items.Flail
                    err = p.Inv.Add(flail)
                    if err != nil {
                        return errors.New("Failed to add flail to inv")
                    }
                case 3:
                    var glaive items.Glaive
                    err = p.Inv.Add(glaive)
                    if err != nil {
                        return errors.New("Failed to add glaive to inv")
                    }
                case 4:
                    var greataxe items.Greataxe
                    err = p.Inv.Add(greataxe)
                    if err != nil {
                        return errors.New("Failed to add greataxe to inv")
                    }
                case 5:
                    var greatsword items.Greatsword
                    err = p.Inv.Add(greatsword)
                    if err != nil {
                        return errors.New("Failed to add greatsword to inv")
                    }
                case 6:
                    var halberd items.Halberd
                    err = p.Inv.Add(halberd)
                    if err != nil {
                        return errors.New("Failed to add halberd to inv")
                    }
                case 7:
                    var lance items.Lance
                    err = p.Inv.Add(lance)
                    if err != nil {
                        return errors.New("Failed to add lance to inv")
                    }
                case 8:
                    var longsword items.Longsword
                    err = p.Inv.Add(longsword)
                    if err != nil {
                        return errors.New("Failed to add longsword to inv")
                    }
                case 9:
                    var maul items.Maul
                    err = p.Inv.Add(maul)
                    if err != nil {
                        return errors.New("Failed to add maul to inv")
                    }
                case 10:
                    var morningstar items.Morningstar
                    err = p.Inv.Add(morningstar)
                    if err != nil {
                        return errors.New("Failed to add morningstar to inv")
                    }
                case 11:
                    var pike items.Pike
                    err = p.Inv.Add(pike)
                    if err != nil {
                        return errors.New("Failed to add pike to inv")
                    }
                case 12:
                    var rapier items.Rapier
                    err = p.Inv.Add(rapier)
                    if err != nil {
                        return errors.New("Failed to add rapier to inv")
                    }
                case 13:
                    var scimitar items.Scimitar
                    err = p.Inv.Add(scimitar)
                    if err != nil {
                        return errors.New("Failed to add scimitar to inv")
                    }
                case 14:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 15:
                    var trident items.Trident
                    err = p.Inv.Add(trident)
                    if err != nil {
                        return errors.New("Failed to add trident to inv")
                    }
                case 16:
                    var warpick items.Warpick
                    err = p.Inv.Add(warpick)
                    if err != nil {
                        return errors.New("Failed to add warpick to inv")
                    }
                case 17:
                    var warhammer items.Warhammer
                    err = p.Inv.Add(warhammer)
                    if err != nil {
                        return errors.New("Failed to add warhammer to inv")
                    }
                case 18:
                    var whip items.Whip
                    err = p.Inv.Add(whip)
                    if err != nil {
                        return errors.New("Failed to add whip to inv")
                    }
                case 19:
                    var blowgun items.Blowgun
                    err = p.Inv.Add(blowgun)
                    if err != nil {
                        return errors.New("Failed to add blowgun to inv")
                    }
                case 20:
                    var handcrossbow items.Handcrossbow
                    err = p.Inv.Add(handcrossbow)
                    if err != nil {
                        return errors.New("Failed to add handcrossbow to inv")
                    }
                case 21:
                    var heavycrossbow items.Heavycrossbow
                    err = p.Inv.Add(heavycrossbow)
                    if err != nil {
                        return errors.New("Failed to add heavycrossbow to inv")
                    }
                case 22:
                    var longbow items.Longbow
                    err = p.Inv.Add(longbow)
                    if err != nil {
                        return errors.New("Failed to add longbow to inv")
                    }
                case 23:
                    var net items.Net
                    err = p.Inv.Add(net)
                    if err != nil {
                        return errors.New("Failed to add net to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 6)")
                }
                switch option4 {
                case 0:
                    var javelin0 items.Javelin
                    var javelin1 items.Javelin
                    var javelin2 items.Javelin
                    var javelin3 items.Javelin
                    var javelin4 items.Javelin
                    err = p.Inv.Add(javelin0)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                    err = p.Inv.Add(javelin1)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                    err = p.Inv.Add(javelin2)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                    err = p.Inv.Add(javelin3)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                    err = p.Inv.Add(javelin4)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 6)")
                }
                switch option5 {
                case 0:
                    var candles = items.Candles{Quantity: 10}
                    var tinderbox items.Tinderbox
                    err = p.Inv.Add(candles)
                    if err != nil {
                        return errors.New("Failed to add candles to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                case 1:
                    var tinderbox items.Tinderbox
                    var torches = items.Torches{Quantity: 10}
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option5 (case 6)")
                }
            case 7:
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "animal handling")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "investigation")
                case 4:
                    proficiencies = append(proficiencies, "nature")
                case 5:
                    proficiencies = append(proficiencies, "perception")
                case 6:
                    proficiencies = append(proficiencies, "stealth")
                case 7:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option0 (case 7)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "animal handling")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "investigation")
                case 4:
                    proficiencies = append(proficiencies, "nature")
                case 5:
                    proficiencies = append(proficiencies, "perception")
                case 6:
                    proficiencies = append(proficiencies, "stealth")
                case 7:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for optino1 (case 7)")
                }
                switch option2 {
                case 0:
                    proficiencies = append(proficiencies, "animal handling")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "investigation")
                case 4:
                    proficiencies = append(proficiencies, "nature")
                case 5:
                    proficiencies = append(proficiencies, "perception")
                case 6:
                    proficiencies = append(proficiencies, "stealth")
                case 7:
                    proficiencies = append(proficiencies, "survival")
                default:
                    return errors.New("Invalid value for option2 (case 7)")
                }
                switch option3 {
                case 0:
                    var scalemail items.Scalemail
                    err = p.Inv.Add(scalemail)
                    if err != nil {
                        return errors.New("Failed to add scalemail to inv")
                    }
                    if ((dex - 10) / 2) > 2 {
                        ac = 16
                    } else {
                        ac = 14 + ((dex - 10) / 2)
                    }
                case 1:
                    var leatherarmor items.Leatherarmor
                    err = p.Inv.Add(leatherarmor)
                    if err != nil {
                        return errors.New("Failed to add leatherarmor to inv")
                    }
                    ac = 11 + ((dex - 10) / 2)
                default:
                    return errors.New("Invalid value for option3 (case 7)")
                }
                switch option4 {
                case 0:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 7)")
                }
                switch option5 {
                case 0:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                default:
                    return errors.New("Invalid value for option5 (case 7)")
                }
                switch option6 {
                case 0:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                case 1:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option6 (case 7)")
                }
            case 8:
                ac = 11 + ((dex - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "deception")
                case 3:
                    proficiencies = append(proficiencies, "insight")
                case 4:
                    proficiencies = append(proficiencies, "intimidation")
                case 5:
                    proficiencies = append(proficiencies, "investigation")
                case 6:
                    proficiencies = append(proficiencies, "perception")
                case 7:
                    proficiencies = append(proficiencies, "performance")
                case 8:
                    proficiencies = append(proficiencies, "persuasion")
                case 9:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 10:
                    proficiencies = append(proficiencies, "stealth")
                default:
                    return errors.New("Invalid value for option0 (case 8)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "deception")
                case 3:
                    proficiencies = append(proficiencies, "insight")
                case 4:
                    proficiencies = append(proficiencies, "intimidation")
                case 5:
                    proficiencies = append(proficiencies, "investigation")
                case 6:
                    proficiencies = append(proficiencies, "perception")
                case 7:
                    proficiencies = append(proficiencies, "performance")
                case 8:
                    proficiencies = append(proficiencies, "persuasion")
                case 9:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 10:
                    proficiencies = append(proficiencies, "stealth")
                default:
                    return errors.New("Invalid value for option1 (case 8)")
                }
                switch option2 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "deception")
                case 3:
                    proficiencies = append(proficiencies, "insight")
                case 4:
                    proficiencies = append(proficiencies, "intimidation")
                case 5:
                    proficiencies = append(proficiencies, "investigation")
                case 6:
                    proficiencies = append(proficiencies, "perception")
                case 7:
                    proficiencies = append(proficiencies, "performance")
                case 8:
                    proficiencies = append(proficiencies, "persuasion")
                case 9:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 10:
                    proficiencies = append(proficiencies, "stealth")
                default:
                    return errors.New("Invalid value for option2 (case 8)")
                }
                switch option3 {
                case 0:
                    proficiencies = append(proficiencies, "acrobatics")
                case 1:
                    proficiencies = append(proficiencies, "athletics")
                case 2:
                    proficiencies = append(proficiencies, "deception")
                case 3:
                    proficiencies = append(proficiencies, "insight")
                case 4:
                    proficiencies = append(proficiencies, "intimidation")
                case 5:
                    proficiencies = append(proficiencies, "investigation")
                case 6:
                    proficiencies = append(proficiencies, "perception")
                case 7:
                    proficiencies = append(proficiencies, "performance")
                case 8:
                    proficiencies = append(proficiencies, "persuasion")
                case 9:
                    proficiencies = append(proficiencies, "sleight of hand")
                case 10:
                    proficiencies = append(proficiencies, "stealth")
                default:
                    return errors.New("Invalid value for option3 (case 8)")
                }
                switch option4 {
                case 0:
                    var rapier items.Rapier
                    err = p.Inv.Add(rapier)
                    if err != nil {
                        return errors.New("Failed to add rapier to inv")
                    }
                case 1:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 8)")
                }
                switch option5 {
                case 0:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 1:
                    var shortsword items.Shortsword
                    err = p.Inv.Add(shortsword)
                    if err != nil {
                        return errors.New("Failed to add shortsword to inv")
                    }
                default:
                    return errors.New("Invalid value for option5 (case 8)")
                }
                switch option6 {
                case 0:
                    var candles = items.Candles{Quantity: 5}
                    var oilflasks = items.Oilflask{Quantity: 2}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(candles)
                    if err != nil {
                        return errors.New("Failed to add candles to inv")
                    }
                    err = p.Inv.Add(oilflasks)
                    if err != nil {
                        return errors.New("Failed to add oilflasks to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                case 1:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                case 2:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option6 (case 8)")
                }
            case 9:
                ac = 10 + ((dex - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "deception")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "intimidation")
                case 4:
                    proficiencies = append(proficiencies, "persuasion")
                case 5:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option0 (case 9)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "deception")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "intimidation")
                case 4:
                    proficiencies = append(proficiencies, "persuasion")
                case 5:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option1 (case 9)")
                }
                switch option2 {
                case 0:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                case 11:
                    var dart items.Dart
                    err = p.Inv.Add(dart)
                    if err != nil {
                        return errors.New("Failed to add dart to inv")
                    }
                case 12:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 13:
                    var sling items.Sling
                    err = p.Inv.Add(sling)
                    if err != nil {
                        return errors.New("Failed to add sling to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 9)")
                }
                switch option3 {
                case 0:
                    var componentpouch items.Componentpouch
                    err = p.Inv.Add(componentpouch)
                    if err != nil {
                        return errors.New("Failed to add componentpouch to inv")
                    }
                case 1:
                    var arcanefocus items.Arcanefocus
                    err = p.Inv.Add(arcanefocus)
                    if err != nil {
                        return errors.New("Failed to add arcanefocus to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 9)")
                }
                switch option4 {
                case 0:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                case 1:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 9)")
                }
            case 10:
                ac = 11 + ((dex - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "deception")
                case 2:
                    proficiencies = append(proficiencies, "history")
                case 3:
                    proficiencies = append(proficiencies, "intimidation")
                case 4:
                    proficiencies = append(proficiencies, "investigation")
                case 5:
                    proficiencies = append(proficiencies, "nature")
                case 6:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option0 (case 10)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "deception")
                case 2:
                    proficiencies = append(proficiencies, "history")
                case 3:
                    proficiencies = append(proficiencies, "intimidation")
                case 4:
                    proficiencies = append(proficiencies, "investigation")
                case 5:
                    proficiencies = append(proficiencies, "nature")
                case 6:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option1 (case 10)")
                }
                switch option2 {
                case 0:
                    var lightcrossbow items.Lightcrossbow
                    err = p.Inv.Add(lightcrossbow)
                    if err != nil {
                        return errors.New("Failed to add lightcrossbow to inv")
                    }
                case 1:
                    var club items.Club
                    err = p.Inv.Add(club)
                    if err != nil {
                        return errors.New("Failed to add club to inv")
                    }
                case 2:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                case 3:
                    var greatclub items.Greatclub
                    err = p.Inv.Add(greatclub)
                    if err != nil {
                        return errors.New("Failed to add greatclub to inv")
                    }
                case 4:
                    var handaxe items.Handaxe
                    err = p.Inv.Add(handaxe)
                    if err != nil {
                        return errors.New("Failed to add handaxe to inv")
                    }
                case 5:
                    var javelin items.Javelin
                    err = p.Inv.Add(javelin)
                    if err != nil {
                        return errors.New("Failed to add javelin to inv")
                    }
                case 6:
                    var lighthammer items.Lighthammer
                    err = p.Inv.Add(lighthammer)
                    if err != nil {
                        return errors.New("Failed to add lighthammer to inv")
                    }
                case 7:
                    var mace items.Mace
                    err = p.Inv.Add(mace)
                    if err != nil {
                        return errors.New("Failed to add mace to inv")
                    }
                case 8:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 9:
                    var sickle items.Sickle
                    err = p.Inv.Add(sickle)
                    if err != nil {
                        return errors.New("Failed to add sickle to inv")
                    }
                case 10:
                    var spear items.Spear
                    err = p.Inv.Add(spear)
                    if err != nil {
                        return errors.New("Failed to add spear to inv")
                    }
                case 11:
                    var dart items.Dart
                    err = p.Inv.Add(dart)
                    if err != nil {
                        return errors.New("Failed to add dart to inv")
                    }
                case 12:
                    var shortbow items.Shortbow
                    err = p.Inv.Add(shortbow)
                    if err != nil {
                        return errors.New("Failed to add shortbow to inv")
                    }
                case 13:
                    var sling items.Sling
                    err = p.Inv.Add(sling)
                    if err != nil {
                        return errors.New("Failed to add sling to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 10)")
                }
                switch option3 {
                case 0:
                    var componentpouch items.Componentpouch
                    err = p.Inv.Add(componentpouch)
                    if err != nil {
                        return errors.New("Failed to add componentpouch to inv")
                    }
                case 1:
                    var arcanefocus items.Arcanefocus
                    err = p.Inv.Add(arcanefocus)
                    if err != nil {
                        return errors.New("Failed to add arcanefocus to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 10)")
                }
                switch option4 {
                case 0:
                    var inkbottle items.Inkbottle
                    var inkpen items.Inkpen
                    var papers = items.Paper{Quantity: 10}
                    err = p.Inv.Add(inkbottle)
                    if err != nil {
                        return errors.New("Failed to add inkbottle to inv")
                    }
                    err = p.Inv.Add(inkpen)
                    if err != nil {
                        return errors.New("Failed to add inkpen to inv")
                    }
                    err = p.Inv.Add(papers)
                    if err != nil {
                        return errors.New("Failed to add papers to inv")
                    }
                case 1:
                    var candles = items.Candles{Quantity: 5}
                    var oilflasks = items.Oilflask{Quantity: 2}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(candles)
                    if err != nil {
                        return errors.New("Failed to add candles to inv")
                    }
                    err = p.Inv.Add(oilflasks)
                    if err != nil {
                        return errors.New("Failed to add oilflasks to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 10)")
                }
            case 11:
                ac = 10 + ((dex - 10) / 2)
                switch option0 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "history")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "investigation")
                case 4:
                    proficiencies = append(proficiencies, "medicine")
                case 5:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option0 (case 11)")
                }
                switch option1 {
                case 0:
                    proficiencies = append(proficiencies, "arcana")
                case 1:
                    proficiencies = append(proficiencies, "history")
                case 2:
                    proficiencies = append(proficiencies, "insight")
                case 3:
                    proficiencies = append(proficiencies, "investigation")
                case 4:
                    proficiencies = append(proficiencies, "medicine")
                case 5:
                    proficiencies = append(proficiencies, "religion")
                default:
                    return errors.New("Invalid value for option1 (case 11)")
                }
                switch option2 {
                case 0:
                    var quarterstaff items.Quarterstaff
                    err = p.Inv.Add(quarterstaff)
                    if err != nil {
                        return errors.New("Failed to add quarterstaff to inv")
                    }
                case 1:
                    var dagger items.Dagger
                    err = p.Inv.Add(dagger)
                    if err != nil {
                        return errors.New("Failed to add dagger to inv")
                    }
                default:
                    return errors.New("Invalid value for option2 (case 11)")
                }
                switch option3 {
                case 0:
                    var componentpouch items.Componentpouch
                    err = p.Inv.Add(componentpouch)
                    if err != nil {
                        return errors.New("Failed to add componentpouch to inv")
                    }
                case 1:
                    var arcanefocus items.Arcanefocus
                    err = p.Inv.Add(arcanefocus)
                    if err != nil {
                        return errors.New("Failed to add arcanefocus to inv")
                    }
                default:
                    return errors.New("Invalid value for option3 (case 11)")
                }
                switch option4 {
                case 0:
                    var inkbottle items.Inkbottle
                    var inkpen items.Inkpen
                    var papers = items.Paper{Quantity: 10}
                    err = p.Inv.Add(inkbottle)
                    if err != nil {
                        return errors.New("Failed to add inkbottle to inv")
                    }
                    err = p.Inv.Add(inkpen)
                    if err != nil {
                        return errors.New("Failed to add inkpen to inv")
                    }
                    err = p.Inv.Add(papers)
                    if err != nil {
                        return errors.New("Failed to add papers to inv")
                    }
                case 1:
                    var torches = items.Torches{Quantity: 10}
                    var tinderbox items.Tinderbox
                    var rope = items.Rope{Length: 50}
                    err = p.Inv.Add(torches)
                    if err != nil {
                        return errors.New("Failed to add torches to inv")
                    }
                    err = p.Inv.Add(tinderbox)
                    if err != nil {
                        return errors.New("Failed to add tinderbox to inv")
                    }
                    err = p.Inv.Add(rope)
                    if err != nil {
                        return errors.New("Failed to add rope to inv")
                    }
                default:
                    return errors.New("Invalid value for option4 (case 11)")
                }
            default:
                return errors.New("Invalid value for classsel")
            }
            p.Stats = &player.Stats{
                AC: ac,
                Str: str,
                StrMod: (str - 10) / 2,
                Dex: dex,
                DexMod: (dex - 10) / 2,
                Con: con,
                ConMod: (con - 10) / 2,
                Intel: intel,
                IntelMod: (intel - 10) / 2,
                Wis: wis,
                WisMod: (wis - 10) / 2,
                Cha: cha,
                ChaMod: (cha - 10) / 2,
                ProfBonus: pb,
                MaxHP: hp,
                HP: hp,
                TempHP: 0,
                HitDice: hd,
                DeathSaveSucc: 0,
                DeathSaveFail: 0,
                Speed: speed,
                Languages: languages,
                Size: size,
                Darkvision: darkvision,
                Proficiencies: proficiencies,
                Resistances: resistances,
                Lucky: lucky,
                Nimbleness: nimbleness,
                Brave: brave,
                Ancestry: ancestry,
            }
            p.Race = racemap[racesel]
            p.Class = classmap[classsel]
            p.Level = 1
            p.XP = 0
            p.Equipment = &player.Equipment{}
            choices = false
            creation = false
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
                var invstr string = p.Inv.Save()
                var statsstr string = p.Stats.Save()
                var equipmentstr string = p.Equipment.Save()
                _, err = db.Exec(saveStmt, name, l.GetName(), l.Pos[0], l.Pos[1], csdonestr, invstr, statsstr, p.Race, p.Class, p.Level, p.XP, equipmentstr)
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
                var csdonestr string
                var invstr string
                for rows.Next() {
                    err = rows.Scan(&savename, &levelname, &x, &y, &csdonestr, &invstr)
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
                    p.Inv.Add(items.LoadItem(itemprops[0], itemprops[1], itemprops[2]))
                }
                l = levels.LoadLvl(levelname, 0, x, y)
                p.Pos = [2]int{-l.Pos[0], -l.Pos[1]}
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
            if inpututil.IsKeyJustPressed(ebiten.KeyM) {
                overworld = !overworld
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
                                npcname = npc.GetName()
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
                                npcname = npc.GetName()
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
                                npcname = npc.GetName()
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
                                npcname = npc.GetName()
                                dialogstrs = npc.Dialog()
                                dialogopen = true
                            }
                        }
                    }
                }
            }
            if !dialogopen && !lvlchange && !start {
                for _, npc := range l.NPCs {
                    if npc.GetSpeed() > 0 && (npcCount + npc.GetOffset()) % npc.GetSpeed() == 0 {
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
                    } else if !npc.Stopped && (npcCount + npc.GetOffset() - 4) % npc.GetSpeed() == 0 {
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
                                if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
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
                                if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
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
                                if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
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
                                if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
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
        animcm := ebiten.ColorM{}
        animcm.Scale(1.0, 1.0, 1.0, float64(y) / 65.0)
        if y <= 65 {
            screen.DrawImage(
                startImage, &ebiten.DrawImageOptions{
                    ColorM: animcm})
            for b := 0; b < 9; b++ {
                for c := 0; c < 12; c++ {
                    raingm := ebiten.GeoM{}
                    raingm.Reset()
                    raingm.Translate(float64(c * 64), float64(b * 64))
                    screen.DrawImage(
                        rainImage.SubImage(
                            image.Rect(
                                (y % 16) * 64, 0, ((y % 16) + 1) * 64, 64)).(*ebiten.Image),
                                &ebiten.DrawImageOptions{GeoM: raingm, ColorM: animcm})
                }
            }
        } else {
            screen.DrawImage(startImage, nil)
            for b := 0; b < 9; b++ {
                for c := 0; c < 12; c++ {
                    raingm := ebiten.GeoM{}
                    raingm.Reset()
                    raingm.Translate(float64(c * 64), float64(b * 64))
                    screen.DrawImage(
                        rainImage.SubImage(
                            image.Rect(
                                (y % 16) * 64, 0, ((y % 16) + 1) * 64, 64)).(*ebiten.Image),
                                &ebiten.DrawImageOptions{GeoM: raingm, ColorM: animcm})
                }
            }
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
    } else if creation {
        // character creation
        text.Draw(screen, fmt.Sprintf("Name:       %s", name), fo, 64, 64, color.White)
        text.Draw(screen, fmt.Sprintf("Race:       %s", racemap[racesel]), fo, 64, 128, color.White)
        text.Draw(screen, fmt.Sprintf("Class:      %s", classmap[classsel]), fo, 64, 192, color.White)
        switch creationsel {
        case 0:
            text.Draw(screen, ">", fo, 32, 128, color.White)
        case 1:
            text.Draw(screen, ">", fo, 32, 192, color.White)
        default:
            log.Fatal("Out of bounds (Draw)")
        }
    } else if racechoices {
        text.Draw(screen, fmt.Sprintf("Race: %s", racemap[racesel]), fo, 64, 32, color.White)
        if dupwarning {
            text.Draw(screen, "No duplicates allowed", fo, 256, 512, color.RGBA{0xff, 0x0, 0x0, 0xff})
        }
        switch racesel {
        case 3:
            text.Draw(screen, "Language:", fo, 64, 64, color.White)
            text.Draw(screen, ">", fo, 432, 64, color.White)
            switch raceopt0 {
            case 0:
                text.Draw(screen, "Dwarvish", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Elvish", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Giant", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Gnomish", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Goblin", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Halfling", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Orc", fo, 448, 64, color.White)
            case 7:
                text.Draw(screen, "Abyssal", fo, 448, 64, color.White)
            case 8:
                text.Draw(screen, "Celestial", fo, 448, 64, color.White)
            case 9:
                text.Draw(screen, "Draconic", fo, 448, 64, color.White)
            case 10:
                text.Draw(screen, "Deep Speech", fo, 448, 64, color.White)
            case 11:
                text.Draw(screen, "Infernal", fo, 448, 64, color.White)
            case 12:
                text.Draw(screen, "Primordial", fo, 448, 64, color.White)
            case 13:
                text.Draw(screen, "Sylvan", fo, 448, 64, color.White)
            case 14:
                text.Draw(screen, "Undercommon", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (2779)")
            }
        case 4:
            text.Draw(screen, "Draconic Ancestry:", fo, 64, 64, color.White)
            text.Draw(screen, ">", fo, 432, 64, color.White)
            switch raceopt0 {
            case 0:
                text.Draw(screen, "Black", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Blue", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Brass", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Bronze", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Copper", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Gold", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Green", fo, 448, 64, color.White)
            case 7:
                text.Draw(screen, "Red", fo, 448, 64, color.White)
            case 8:
                text.Draw(screen, "Silver", fo, 448, 64, color.White)
            case 9:
                text.Draw(screen, "White", fo, 448, 64, color.White)
            default:
                log.Fatal("Ount of bounds (2805)")
            }
        case 6:
            text.Draw(screen, "+1 to Attributes:", fo, 64, 64, color.White)
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 160, color.White)
            text.Draw(screen, "Language:", fo, 64, 256, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 256, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch raceopt0 {
            case 0:
                text.Draw(screen, "Strength", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Dexterity", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Constitution", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Intelligence", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Wisdom", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Charisma", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (2836)")
            }
            switch raceopt1 {
            case 0:
                text.Draw(screen, "Strength", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Dexterity", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Constitution", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Intelligence", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Wisdom", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Charisma", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (2852)")
            }
            switch raceopt2 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "History", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 448, 160, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 448, 160, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 448, 160, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 448, 160, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 448, 160, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 448, 160, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 448, 160, color.White)
            case 17:
                text.Draw(screen, "Survival", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (2892)")
            }
            switch raceopt3 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 192, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 448, 192, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 448, 192, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 448, 192, color.White)
            case 5:
                text.Draw(screen, "History", fo, 448, 192, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 448, 192, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 448, 192, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 448, 192, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 448, 192, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 448, 192, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 448, 192, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 448, 192, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 448, 192, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 448, 192, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 448, 192, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 448, 192, color.White)
            case 17:
                text.Draw(screen, "Survival", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (2932)")
            }
            switch raceopt4 {
            case 0:
                text.Draw(screen, "Dwarvish", fo, 448, 256, color.White)
            case 1:
                text.Draw(screen, "Undercommon", fo, 448, 256, color.White)
            case 2:
                text.Draw(screen, "Giant", fo, 448, 256, color.White)
            case 3:
                text.Draw(screen, "Gnomish", fo, 448, 256, color.White)
            case 4:
                text.Draw(screen, "Goblin", fo, 448, 256, color.White)
            case 5:
                text.Draw(screen, "Halfling", fo, 448, 256, color.White)
            case 6:
                text.Draw(screen, "Orc", fo, 448, 256, color.White)
            case 7:
                text.Draw(screen, "Abyssal", fo, 448, 256, color.White)
            case 8:
                text.Draw(screen, "Celestial", fo, 448, 256, color.White)
            case 9:
                text.Draw(screen, "Draconic", fo, 448, 256, color.White)
            case 10:
                text.Draw(screen, "Deep Speech", fo, 448, 256, color.White)
            case 11:
                text.Draw(screen, "Infernal", fo, 448, 256, color.White)
            case 12:
                text.Draw(screen, "Primordial", fo, 448, 256, color.White)
            case 13:
                text.Draw(screen, "Sylvan", fo, 448, 256, color.White)
            default:
                log.Fatal("Out of bounds ()")
            }
        default:
            fmt.Println("Skipping race choices")
        }
    } else if choices {
        text.Draw(screen, fmt.Sprintf("Class: %s", classmap[classsel]), fo, 64, 32, color.White)
        if dupwarning {
            text.Draw(screen, "No duplicates allowed", fo, 256, 512, color.RGBA{0xff, 0x0, 0x0, 0xff})
        }
        switch classsel {
        case 0:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Nature", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Perception", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Survival", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (2604)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Nature", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Perception", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Survival", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (2620)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Greataxe", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Battleaxe", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Flail", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Glaive", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Greatsword", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "Halberd", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Lance", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Longsword", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Maul", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Morningstar", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Pike", fo, 448, 160, color.White)
            case 11:
                text.Draw(screen, "Rapier", fo, 448, 160, color.White)
            case 12:
                text.Draw(screen, "Scimitar", fo, 448, 160, color.White)
            case 13:
                text.Draw(screen, "Shortsword", fo, 448, 160, color.White)
            case 14:
                text.Draw(screen, "Trident", fo, 448, 160, color.White)
            case 15:
                text.Draw(screen, "War pick", fo, 448, 160, color.White)
            case 16:
                text.Draw(screen, "Warhammer", fo, 448, 160, color.White)
            case 17:
                text.Draw(screen, "Whip", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (2660)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Two Handaxes", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 192, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 192, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 192, color.White)
            case 4:
                text.Draw(screen, "Javelin", fo, 448, 192, color.White)
            case 5:
                text.Draw(screen, "Light Hammer", fo, 448, 192, color.White)
            case 6:
                text.Draw(screen, "Mace", fo, 448, 192, color.White)
            case 7:
                text.Draw(screen, "Quarterstaff", fo, 448, 192, color.White)
            case 8:
                text.Draw(screen, "Sickle", fo, 448, 192, color.White)
            case 9:
                text.Draw(screen, "Spear", fo, 448, 192, color.White)
            case 10:
                text.Draw(screen, "Light Crossbow", fo, 448, 192, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 448, 192, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 448, 192, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (2692)")
            }
        case 1:
            text.Draw(screen, "Instrument Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 192, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 320, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 128, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            case 5:
                text.Draw(screen, ">", fo, 432, 256, color.White)
            case 6:
                text.Draw(screen, ">", fo, 432, 320, color.White)
            case 7:
                text.Draw(screen, ">", fo, 432, 352, color.White)
            case 8:
                text.Draw(screen, ">", fo, 432, 384, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 448, 64, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 448, 64, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 448, 64, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (2720)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 448, 96, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 448, 96, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 448, 96, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 448, 96, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (2744)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 448, 128, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 448, 128, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 448, 128, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 448, 128, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 448, 128, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 448, 128, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 448, 128, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 448, 128, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 448, 128, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 448, 128, color.White)
            default:
                log.Fatal("Out of bounds (2768)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 192, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 448, 192, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 448, 192, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 448, 192, color.White)
            case 5:
                text.Draw(screen, "History", fo, 448, 192, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 448, 192, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 448, 192, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 448, 192, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 448, 192, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 448, 192, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 448, 192, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 448, 192, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 448, 192, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 448, 192, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 448, 192, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 448, 192, color.White)
            case 17:
                text.Draw(screen, "Survival", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (2808)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 224, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 448, 224, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 448, 224, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 448, 224, color.White)
            case 5:
                text.Draw(screen, "History", fo, 448, 224, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 448, 224, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 448, 224, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 448, 224, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 448, 224, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 448, 224, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 448, 224, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 448, 224, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 448, 224, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 448, 224, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 448, 224, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 448, 224, color.White)
            case 17:
                text.Draw(screen, "Survival", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (2848)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 256, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 256, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 448, 256, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 448, 256, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 448, 256, color.White)
            case 5:
                text.Draw(screen, "History", fo, 448, 256, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 448, 256, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 448, 256, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 448, 256, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 448, 256, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 448, 256, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 448, 256, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 448, 256, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 448, 256, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 448, 256, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 448, 256, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 448, 256, color.White)
            case 17:
                text.Draw(screen, "Survival", fo, 448, 256, color.White)
            default:
                log.Fatal("Out of bounds (2888)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Rapier", fo, 448, 320, color.White)
            case 1:
                text.Draw(screen, "Longsword", fo, 448, 320, color.White)
            case 2:
                text.Draw(screen, "Club", fo, 448, 320, color.White)
            case 3:
                text.Draw(screen, "Dagger", fo, 448, 320, color.White)
            case 4:
                text.Draw(screen, "Greatclub", fo, 448, 320, color.White)
            case 5:
                text.Draw(screen, "Handaxe", fo, 448, 320, color.White)
            case 6:
                text.Draw(screen, "Javelin", fo, 448, 320, color.White)
            case 7:
                text.Draw(screen, "Light hammer", fo, 448, 320, color.White)
            case 8:
                text.Draw(screen, "Mace", fo, 448, 320, color.White)
            case 9:
                text.Draw(screen, "Quarterstaff", fo, 448, 320, color.White)
            case 10:
                text.Draw(screen, "Sickle", fo, 448, 320, color.White)
            case 11:
                text.Draw(screen, "Spear", fo, 448, 320, color.White)
            case 12:
                text.Draw(screen, "Light crossbow", fo, 448, 320, color.White)
            case 13:
                text.Draw(screen, "Dart", fo, 448, 320, color.White)
            case 14:
                text.Draw(screen, "Shortbow", fo, 448, 320, color.White)
            case 15:
                text.Draw(screen, "Sling", fo, 448, 320, color.White)
            default:
                log.Fatal("Out of bounds (2924)")
            }
            switch option7 {
            case 0:
                text.Draw(screen, "Diplomat's Pack", fo, 448, 352, color.White)
            case 1:
                text.Draw(screen, "Entertainer's Pack", fo, 448, 352, color.White)
            default:
                log.Fatal("Out of bounds (2932)")
            }
            switch option8 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 448, 384, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 448, 384, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 448, 384, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 448, 384, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 448, 384, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 448, 384, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 448, 384, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 448, 384, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 448, 384, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 448, 384, color.White)
            default:
                log.Fatal("Out of bounds (2956)")
            }
        case 2:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            case 5:
                text.Draw(screen, ">", fo, 432, 256, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "History", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Medicine", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Persuasion", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (2973)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "History", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Medicine", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Persuasion", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (2987)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Mace", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Warhammer", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (2995)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Scale mail", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Leather armor", fo, 448, 192, color.White)
            case 2:
                text.Draw(screen, "Chain mail", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3005)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 224, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 224, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 224, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 224, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 224, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 224, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 224, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 224, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 224, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 224, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 448, 224, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 448, 224, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (3037)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Priest's Pack", fo, 448, 256, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 448, 256, color.White)
            default:
                log.Fatal("Out of bounds (3045)")
            }
        case 3:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 448, 64, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3068)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 448, 96, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 448, 96, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3088)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Wooden shield", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 160, color.White)
            case 11:
                text.Draw(screen, "Light crossbow", fo, 448, 160, color.White)
            case 12:
                text.Draw(screen, "Dart", fo, 448, 160, color.White)
            case 13:
                text.Draw(screen, "Shortbow", fo, 448, 160, color.White)
            case 14:
                text.Draw(screen, "Sling", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (3122)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Scimitar", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 192, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 192, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 192, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 192, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 192, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 192, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 192, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 192, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 192, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3148)")
            }
        case 4:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            case 5:
                text.Draw(screen, ">", fo, 432, 256, color.White)
            case 6:
                text.Draw(screen, ">", fo, 432, 288, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Athletics", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "History", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Intimidation", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 448, 64, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3171)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Athletics", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "History", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Intimidation", fo, 448, 96, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 448, 96, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3191)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Chain mail", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Leather armor + longbow", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (3199)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Battleaxe", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Flail", fo, 448, 192, color.White)
            case 2:
                text.Draw(screen, "Glaive", fo, 448, 192, color.White)
            case 3:
                text.Draw(screen, "Greataxe", fo, 448, 192, color.White)
            case 4:
                text.Draw(screen, "Greatsword", fo, 448, 192, color.White)
            case 5:
                text.Draw(screen, "Halberd", fo, 448, 192, color.White)
            case 6:
                text.Draw(screen, "Lance", fo, 448, 192, color.White)
            case 7:
                text.Draw(screen, "Longsword", fo, 448, 192, color.White)
            case 8:
                text.Draw(screen, "Maul", fo, 448, 192, color.White)
            case 9:
                text.Draw(screen, "Morningstar", fo, 448, 192, color.White)
            case 10:
                text.Draw(screen, "Pike", fo, 448, 192, color.White)
            case 11:
                text.Draw(screen, "Rapier", fo, 448, 192, color.White)
            case 12:
                text.Draw(screen, "Scimitar", fo, 448, 192, color.White)
            case 13:
                text.Draw(screen, "Shortsword", fo, 448, 192, color.White)
            case 14:
                text.Draw(screen, "Trident", fo, 448, 192, color.White)
            case 15:
                text.Draw(screen, "War pick", fo, 448, 192, color.White)
            case 16:
                text.Draw(screen, "Warhammer", fo, 448, 192, color.White)
            case 17:
                text.Draw(screen, "Whip", fo, 448, 192, color.White)
            case 18:
                text.Draw(screen, "Blowgun", fo, 448, 192, color.White)
            case 19:
                text.Draw(screen, "Hand crossbow", fo, 448, 192, color.White)
            case 20:
                text.Draw(screen, "Heavy crossbow", fo, 448, 192, color.White)
            case 21:
                text.Draw(screen, "Longbow", fo, 448, 192, color.White)
            case 22:
                text.Draw(screen, "Net", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3249)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Shield", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Battleaxe", fo, 448, 224, color.White)
            case 2:
                text.Draw(screen, "Flail", fo, 448, 224, color.White)
            case 3:
                text.Draw(screen, "Glaive", fo, 448, 224, color.White)
            case 4:
                text.Draw(screen, "Greataxe", fo, 448, 224, color.White)
            case 5:
                text.Draw(screen, "Greatsword", fo, 448, 224, color.White)
            case 6:
                text.Draw(screen, "Halberd", fo, 448, 224, color.White)
            case 7:
                text.Draw(screen, "Lance", fo, 448, 224, color.White)
            case 8:
                text.Draw(screen, "Longsword", fo, 448, 224, color.White)
            case 9:
                text.Draw(screen, "Maul", fo, 448, 224, color.White)
            case 10:
                text.Draw(screen, "Morningstar", fo, 448, 224, color.White)
            case 11:
                text.Draw(screen, "Pike", fo, 448, 224, color.White)
            case 12:
                text.Draw(screen, "Rapier", fo, 448, 224, color.White)
            case 13:
                text.Draw(screen, "Scimitar", fo, 448, 224, color.White)
            case 14:
                text.Draw(screen, "Shortsword", fo, 448, 224, color.White)
            case 15:
                text.Draw(screen, "Trident", fo, 448, 224, color.White)
            case 16:
                text.Draw(screen, "War pick", fo, 448, 224, color.White)
            case 17:
                text.Draw(screen, "Warhammer", fo, 448, 224, color.White)
            case 18:
                text.Draw(screen, "Whip", fo, 448, 224, color.White)
            case 19:
                text.Draw(screen, "Blowgun", fo, 448, 224, color.White)
            case 20:
                text.Draw(screen, "Hand crossbow", fo, 448, 224, color.White)
            case 21:
                text.Draw(screen, "Heavy crossbow", fo, 448, 224, color.White)
            case 22:
                text.Draw(screen, "Longbow", fo, 448, 224, color.White)
            case 23:
                text.Draw(screen, "Net", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (3301)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 448, 256, color.White)
            case 1:
                text.Draw(screen, "Two handaxes", fo, 448, 256, color.White)
            default:
                log.Fatal("Out of bounds (3309)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 448, 288, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 448, 288, color.White)
            default:
                log.Fatal("Out of bounds (3317)")
            }
        case 5:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "History", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Stealth", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3336)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "History", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Stealth", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3352)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Shortsword", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 160, color.White)
            case 11:
                text.Draw(screen, "Light crossbow", fo, 448, 160, color.White)
            case 12:
                text.Draw(screen, "Dart", fo, 448, 160, color.White)
            case 13:
                text.Draw(screen, "Shortbow", fo, 448, 160, color.White)
            case 14:
                text.Draw(screen, "Sling", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (3386)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3394)")
            }
        case 6:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            case 5:
                text.Draw(screen, ">", fo, 432, 256, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Athletics", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3413)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Athletics", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3429)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Battleaxe", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Flail", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Glaive", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Greataxe", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Greatsword", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "Halberd", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Lance", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Longsword", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Maul", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Morningstar", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Pike", fo, 448, 160, color.White)
            case 11:
                text.Draw(screen, "Rapier", fo, 448, 160, color.White)
            case 12:
                text.Draw(screen, "Scimitar", fo, 448, 160, color.White)
            case 13:
                text.Draw(screen, "Shortsword", fo, 448, 160, color.White)
            case 14:
                text.Draw(screen, "Trident", fo, 448, 160, color.White)
            case 15:
                text.Draw(screen, "War pick", fo, 448, 160, color.White)
            case 16:
                text.Draw(screen, "Warhammer", fo, 448, 160, color.White)
            case 17:
                text.Draw(screen, "Whip", fo, 448, 160, color.White)
            case 18:
                text.Draw(screen, "Blowgun", fo, 448, 160, color.White)
            case 19:
                text.Draw(screen, "Hand crossbow", fo, 448, 160, color.White)
            case 20:
                text.Draw(screen, "Heavy crossbow", fo, 448, 160, color.White)
            case 21:
                text.Draw(screen, "Longbow", fo, 448, 160, color.White)
            case 22:
                text.Draw(screen, "Net", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (3479)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Shield", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Battleaxe", fo, 448, 192, color.White)
            case 2:
                text.Draw(screen, "Flail", fo, 448, 192, color.White)
            case 3:
                text.Draw(screen, "Glaive", fo, 448, 192, color.White)
            case 4:
                text.Draw(screen, "Greataxe", fo, 448, 192, color.White)
            case 5:
                text.Draw(screen, "Greatsword", fo, 448, 192, color.White)
            case 6:
                text.Draw(screen, "Halberd", fo, 448, 192, color.White)
            case 7:
                text.Draw(screen, "Lance", fo, 448, 192, color.White)
            case 8:
                text.Draw(screen, "Longsword", fo, 448, 192, color.White)
            case 9:
                text.Draw(screen, "Maul", fo, 448, 192, color.White)
            case 10:
                text.Draw(screen, "Morningstar", fo, 448, 192, color.White)
            case 11:
                text.Draw(screen, "Pike", fo, 448, 192, color.White)
            case 12:
                text.Draw(screen, "Rapier", fo, 448, 192, color.White)
            case 13:
                text.Draw(screen, "Scimitar", fo, 448, 192, color.White)
            case 14:
                text.Draw(screen, "Shortsword", fo, 448, 192, color.White)
            case 15:
                text.Draw(screen, "Trident", fo, 448, 192, color.White)
            case 16:
                text.Draw(screen, "War pick", fo, 448, 192, color.White)
            case 17:
                text.Draw(screen, "Warhammer", fo, 448, 192, color.White)
            case 18:
                text.Draw(screen, "Whip", fo, 448, 192, color.White)
            case 19:
                text.Draw(screen, "Blowgun", fo, 448, 192, color.White)
            case 20:
                text.Draw(screen, "Hand crossbow", fo, 448, 192, color.White)
            case 21:
                text.Draw(screen, "Heavy crossbow", fo, 448, 192, color.White)
            case 22:
                text.Draw(screen, "Longbow", fo, 448, 192, color.White)
            case 23:
                text.Draw(screen, "Net", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3531)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Five javelins", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 224, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 224, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 224, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 224, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 224, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 224, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 224, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 224, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 224, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (3557)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Priest's Pack", fo, 448, 256, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 448, 256, color.White)
            default:
                log.Fatal("Out of bounds (3565)")
            }
        case 7:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 192, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 128, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            case 5:
                text.Draw(screen, ">", fo, 432, 256, color.White)
            case 6:
                text.Draw(screen, ">", fo, 432, 288, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Stealth", fo, 448, 64, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3588)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 448, 96, color.White)
            case 6:
                text.Draw(screen, "Stealth", fo, 448, 96, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3608)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 448, 128, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 128, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 128, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 448, 128, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 448, 128, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 448, 128, color.White)
            case 6:
                text.Draw(screen, "Stealth", fo, 448, 128, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 448, 128, color.White)
            default:
                log.Fatal("Out of bounds (3628)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Scale mail", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Leather armor", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3636)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Shortsword", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 224, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 224, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 224, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 224, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 224, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 224, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 224, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 224, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 224, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (3662)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Shortsword", fo, 448, 256, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 256, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 256, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 256, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 256, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 256, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 256, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 256, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 256, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 256, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 256, color.White)
            default:
                log.Fatal("Out of bounds (3688)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 448, 288, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 448, 288, color.White)
            default:
                log.Fatal("Out of bounds (3696)")
            }
        case 8:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 224, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 128, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            case 5:
                text.Draw(screen, ">", fo, 432, 256, color.White)
            case 6:
                text.Draw(screen, ">", fo, 432, 288, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 448, 64, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 448, 64, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 448, 64, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 448, 64, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3725)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 448, 96, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 448, 96, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 448, 96, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 448, 96, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 448, 96, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3751)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 128, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 128, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 448, 128, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 448, 128, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 448, 128, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 448, 128, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 448, 128, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 448, 128, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 448, 128, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 448, 128, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 448, 128, color.White)
            default:
                log.Fatal("Out of bounds (3777)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (3803)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Rapier", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Shortsword", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (3811)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Shortbow", fo, 448, 256, color.White)
            case 1:
                text.Draw(screen, "Shortsword", fo, 448, 256, color.White)
            default:
                log.Fatal("Out of bounds (3819)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Burglar's Pack", fo, 448, 288, color.White)
            case 1:
                text.Draw(screen, "Dungeoneer's Pack", fo, 448, 288, color.White)
            case 2:
                text.Draw(screen, "Explorer's Pack", fo, 448, 288, color.White)
            default:
                log.Fatal("Out of bounds (3829)")
            }
        case 9:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3848)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3864)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 160, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 448, 160, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 448, 160, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (3896)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Component pouch", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Arcane focus", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3904)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (3912)")
            }
        case 10:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "History", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Investigation", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Nature", fo, 448, 64, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (3933)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "History", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Investigation", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Nature", fo, 448, 96, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (3951)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 448, 160, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 448, 160, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 448, 160, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 448, 160, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 448, 160, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 448, 160, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 448, 160, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 448, 160, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 448, 160, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 448, 160, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 448, 160, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 448, 160, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (3983)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Component pouch", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Arcane focus", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (3991)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Scholar's Pack", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Dungeoneer's Pack", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (3999)")
            }
        case 11:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 432, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 432, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 432, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 432, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 432, 224, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 64, color.White)
            case 1:
                text.Draw(screen, "History", fo, 448, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 64, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 448, 64, color.White)
            case 4:
                text.Draw(screen, "Medicine", fo, 448, 64, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 448, 64, color.White)
            default:
                log.Fatal("Out of bounds (4018)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 448, 96, color.White)
            case 1:
                text.Draw(screen, "History", fo, 448, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 448, 96, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 448, 96, color.White)
            case 4:
                text.Draw(screen, "Medicine", fo, 448, 96, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 448, 96, color.White)
            default:
                log.Fatal("Out of bounds (4034)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Quarterstaff", fo, 448, 160, color.White)
            case 1:
                text.Draw(screen, "Dagger", fo, 448, 160, color.White)
            default:
                log.Fatal("Out of bounds (4042)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Component pouch", fo, 448, 192, color.White)
            case 1:
                text.Draw(screen, "Arcane focus", fo, 448, 192, color.White)
            default:
                log.Fatal("Out of bounds (4050)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Scholar's pack", fo, 448, 224, color.White)
            case 1:
                text.Draw(screen, "Explorer's pack", fo, 448, 224, color.White)
            default:
                log.Fatal("Out of bounds (4058)")
            }
        default:
            log.Fatal("Out of bounds (Draw choices)")
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
    } else {
        ebitenutil.DebugPrintAt(screen, "Loading...", w / 2, h / 2)
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
    if overworld {
        blankImage := ebiten.NewImage(w, h)
        blankImage.Fill(color.RGBA{0x00, 0x00, 0x00, 0xb0})
        screen.DrawImage(blankImage, nil)
        iw, _ := overworldImage.Size()
        owgm := ebiten.GeoM{}
        owgm.Translate(float64(iw) - (float64(iw) / 2.0), 0.0)
        screen.DrawImage(
            overworldImage, &ebiten.DrawImageOptions{
                GeoM: owgm})
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

    //pcimage, _, err := image.Decode(bytes.NewReader(pcimages.PC_Diamond_Sword_PNG))
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

    overworldimg, _, err := image.Decode(bytes.NewReader(assets.Overworld_PNG))
    if err != nil {
        log.Fatal(err)
    }
    overworldImage = ebiten.NewImageFromImage(overworldimg)

    icon16img, _, err = image.Decode(bytes.NewReader(assets.Icon_16_PNG))
    if err != nil {
        log.Fatal(err)
    }

    icon32img, _, err = image.Decode(bytes.NewReader(assets.Icon_32_PNG))
    if err != nil {
        log.Fatal(err)
    }

    icon48img, _, err = image.Decode(bytes.NewReader(assets.Icon_48_PNG))
    if err != nil {
        log.Fatal(err)
    }

    racemap[0] = "Dwarf"
    racemap[1] = "Elf"
    racemap[2] = "Halfling"
    racemap[3] = "Human"
    racemap[4] = "Dragonborn"
    racemap[5] = "Gnome"
    racemap[6] = "Half-Elf"
    racemap[7] = "Half-Orc"
    racemap[8] = "Tiefling"

    classmap[0] = "Barbarian"
    classmap[1] = "Bard"
    classmap[2] = "Cleric"
    classmap[3] = "Druid"
    classmap[4] = "Fighter"
    classmap[5] = "Monk"
    classmap[6] = "Paladin"
    classmap[7] = "Ranger"
    classmap[8] = "Rogue"
    classmap[9] = "Sorceror"
    classmap[10] = "Warlock"
    classmap[11] = "Wizard"

    savesTableSchema = []string{"name,TEXT,1,null,1", "level,TEXT,1,\"One\",0", "x,INT,1,null,0", "y,INT,1,null,0", "csdone,TEXT,0,null,0", "inventory,TEXT,0,null,0", "stats,TEXT,0,null,0", "race,TEXT,0,null,0", "class,TEXT,0,null,0", "playerlevel,INT,0,null,0", "xp,INT,0,null,0", "equipment,TEXT,0,null,0"}
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
        if schemaRowsIndex >= len(savesTableSchema) {
            fixSchema = true
        } else if savesTableSchema[schemaRowsIndex] != schemaRowsName + "," + schemaRowsType + "," + strconv.Itoa(schemaRowsNotNull) + "," + fmt.Sprint(schemaRowsDefault) + "," + strconv.Itoa(schemaRowsPk) {
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
        var colNames []string
        for copyCols.Next() {
            var colName string
            var trash1 string
            var trash2 int
            var trash3 int
            var trash4 string
            var trash5 string
            err = copyCols.Scan(&trash1, &colName, &trash2, &trash3, &trash4, &trash5)
            colNames = append(colNames, colName)
        }
        for _, colName := range colNames {
            for _, colSchema := range savesTableSchema {
                colSchemaArr := strings.Split(colSchema, ",")
                if colSchemaArr[0] == colName {
                    colsStr += colName + ", "
                }
            }
        }
        colsStr = colsStr[:len(colsStr) - 2]
        copyRows, err := db.Query("select " + colsStr + " from copied")
        if err != nil {
            log.Fatal(err)
        }
        defer copyRows.Close()
        var colsArr = strings.Split(colsStr, ",")
        var numCols = len(colsArr)
        var insertStmts = make([]string, 0)
        var copyRowsPtrs = make([]interface{}, numCols)
        var copyRowsArr = make([]interface{}, numCols)
        for i, _ := range copyRowsPtrs {
            copyRowsPtrs[i] = &copyRowsArr[i]
        }
        for copyRows.Next() {
            err = copyRows.Scan(copyRowsPtrs...)
            insertStmt := "insert into saves ("
            for cind, col := range colsArr {
                if cind == numCols - 1 {
                    insertStmt += col + ") values ("
                } else {
                    insertStmt += col + ", "
                }
            }
            for whatever, whateverPtr := range copyRowsArr {
                switch reflect.TypeOf(whateverPtr).String() {
                case "string":
                    if whatever == len(copyRowsArr) - 1 {
                        insertStmt += "\"" + fmt.Sprint(whateverPtr) + "\");"
                    } else {
                        insertStmt += "\"" + fmt.Sprint(whateverPtr) + "\", "
                    }
                case "int64":
                    if whatever == len(copyRowsArr) - 1 {
                        insertStmt += fmt.Sprint(whateverPtr) + ");"
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
}

func main() {
    ebiten.SetWindowSize(768, 576)
    ebiten.SetWindowTitle("CHANGEME")
    ebiten.SetWindowIcon([]image.Image{icon16img, icon32img, icon48img})

    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
