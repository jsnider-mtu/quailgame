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
)

var racemap = make(map[int]string)
var classmap = make(map[int]string)
var backgroundmap = make(map[int]string)
var equipmentmap = make(map[int]string)

type Game struct {}

func (g *Game) Update() error {
    abilities := make([]int, 6)
    var str int
    var dex int
    var con int
    var intel int
    var wis int
    var cha int
    var pb int
    var hp int
    var hd string
    var speed int
    var size int // 0: Small, 1: Medium, 2: Large
    var languages = make([]string, 0)
    var proficiencies = make([]string, 0)
    var resistances = make([]string, 0)
    var darkvision bool = false
    var lucky bool = false
    var nimbleness bool = false
    var savingthrows = make(map[string]int)
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
            case 2:
                backgroundsel--
                if backgroundsel < 0 {
                    backgroundsel = 0
                }
            case 3:
                equipmentsel--
                if equipmentsel < 0 {
                    equipmentsel = 0
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
            case 2:
                backgroundsel++
                if backgroundsel > 12 {
                    backgroundsel = 12
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
        } else if creationsel > 2 {
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
                // brave
            case 3:
                str++
                dex++
                con++
                intel++
                wis++
                cha++
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
            choices = true
        }
    } else if choices {
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
                        if prof == "warhammer" || prof == "martial weapons" {
                            if option2 > 1 {
                                option2 = 1
                            }
                        } else {
                            if option2 > 0 {
                                option2 = 0
                            }
                        }
                    }
                case 3:
                    option3++
                    for _, prof := range proficiencies {
                        if prof == "heavy armor" || prof == "all armor" {
                            if option3 > 2 {
                                option3 = 2
                            }
                        } else {
                            if option3 > 1 {
                                option3 = 1
                            }
                        }
                    }
                case 4:
                    option4++
                    if option4 > 14 {
                        option4 = 14
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
            if option0 == option1 {
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
                case 4:
                    option4--
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 { option5 = 0
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
                    if option4 > 1 {
                        option4 = 1
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
            if option0 == option1 {
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
                    if option4 < 0 {
                        option4 = 0
                    }
                case 5:
                    option5--
                    if option5 < 0 {
                        option5 = 0
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
            } else if creationsel > 3 {
                creationsel = 3
            }
            if option0 == option1 || option1 == option2 || option0 == option2 {
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
                    if option4 < 1 {
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
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
            switch classsel {
            case 0:
            case 1:
            case 2:
            case 3:
            case 4:
            case 5:
            case 6:
            case 7:
            case 8:
            case 9:
            case 10:
            case 11:
            default:
                return errors.New("Invalid value for classsel")
            }
            p.Stats = &player.Stats{
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
            }
            p.Race = racemap[racesel]
            p.Class = classmap[classsel]
            //p.Background = backgroundmap[backgroundsel]
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
                var invstr string
                for itemind, item := range p.Inv.GetItems() {
                    if itemind == len(p.Inv.GetItems()) - 1 {
                        invstr += item.Save()
                    } else {
                        invstr += item.Save() + ";"
                    }
                }
                _, err = db.Exec(saveStmt, name, l.GetName(), l.Pos[0], l.Pos[1], csdonestr, invstr)
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
        text.Draw(screen, fmt.Sprintf("Background: %s", backgroundmap[backgroundsel]), fo, 64, 256, color.White)
        switch creationsel {
        case 0:
            text.Draw(screen, ">", fo, 32, 128, color.White)
        case 1:
            text.Draw(screen, ">", fo, 32, 192, color.White)
        case 2:
            text.Draw(screen, ">", fo, 32, 256, color.White)
        default:
            log.Fatal("Out of bounds (Draw)")
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
                text.Draw(screen, ">", fo, 496, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 496, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 496, 160, color.White)
            case 3:
                text.Draw(screen, ">", fo, 496, 192, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 512, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 512, 64, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 512, 64, color.White)
            case 3:
                text.Draw(screen, "Nature", fo, 512, 64, color.White)
            case 4:
                text.Draw(screen, "Perception", fo, 512, 64, color.White)
            case 5:
                text.Draw(screen, "Survival", fo, 512, 64, color.White)
            default:
                log.Fatal("Out of bounds (2604)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 512, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 512, 96, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 512, 96, color.White)
            case 3:
                text.Draw(screen, "Nature", fo, 512, 96, color.White)
            case 4:
                text.Draw(screen, "Perception", fo, 512, 96, color.White)
            case 5:
                text.Draw(screen, "Survival", fo, 512, 96, color.White)
            default:
                log.Fatal("Out of bounds (2620)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Greataxe", fo, 512, 160, color.White)
            case 1:
                text.Draw(screen, "Battleaxe", fo, 512, 160, color.White)
            case 2:
                text.Draw(screen, "Flail", fo, 512, 160, color.White)
            case 3:
                text.Draw(screen, "Glaive", fo, 512, 160, color.White)
            case 4:
                text.Draw(screen, "Greatsword", fo, 512, 160, color.White)
            case 5:
                text.Draw(screen, "Halberd", fo, 512, 160, color.White)
            case 6:
                text.Draw(screen, "Lance", fo, 512, 160, color.White)
            case 7:
                text.Draw(screen, "Longsword", fo, 512, 160, color.White)
            case 8:
                text.Draw(screen, "Maul", fo, 512, 160, color.White)
            case 9:
                text.Draw(screen, "Morningstar", fo, 512, 160, color.White)
            case 10:
                text.Draw(screen, "Pike", fo, 512, 160, color.White)
            case 11:
                text.Draw(screen, "Rapier", fo, 512, 160, color.White)
            case 12:
                text.Draw(screen, "Scimitar", fo, 512, 160, color.White)
            case 13:
                text.Draw(screen, "Shortsword", fo, 512, 160, color.White)
            case 14:
                text.Draw(screen, "Trident", fo, 512, 160, color.White)
            case 15:
                text.Draw(screen, "War pick", fo, 512, 160, color.White)
            case 16:
                text.Draw(screen, "Warhammer", fo, 512, 160, color.White)
            case 17:
                text.Draw(screen, "Whip", fo, 512, 160, color.White)
            default:
                log.Fatal("Out of bounds (2660)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Two Handaxes", fo, 512, 192, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 512, 192, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 512, 192, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 512, 192, color.White)
            case 4:
                text.Draw(screen, "Javelin", fo, 512, 192, color.White)
            case 5:
                text.Draw(screen, "Light Hammer", fo, 512, 192, color.White)
            case 6:
                text.Draw(screen, "Mace", fo, 512, 192, color.White)
            case 7:
                text.Draw(screen, "Quarterstaff", fo, 512, 192, color.White)
            case 8:
                text.Draw(screen, "Sickle", fo, 512, 192, color.White)
            case 9:
                text.Draw(screen, "Spear", fo, 512, 192, color.White)
            case 10:
                text.Draw(screen, "Light Crossbow", fo, 512, 192, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 512, 192, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 512, 192, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 512, 192, color.White)
            default:
                log.Fatal("Out of bounds (2692)")
            }
        case 1:
            text.Draw(screen, "Instrument Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 192, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 320, color.White)
            switch creationsel {
            case 0:
                text.Draw(screen, ">", fo, 496, 64, color.White)
            case 1:
                text.Draw(screen, ">", fo, 496, 96, color.White)
            case 2:
                text.Draw(screen, ">", fo, 496, 128, color.White)
            case 3:
                text.Draw(screen, ">", fo, 496, 192, color.White)
            case 4:
                text.Draw(screen, ">", fo, 496, 224, color.White)
            case 5:
                text.Draw(screen, ">", fo, 496, 256, color.White)
            case 6:
                text.Draw(screen, ">", fo, 496, 320, color.White)
            case 7:
                text.Draw(screen, ">", fo, 496, 352, color.White)
            case 8:
                text.Draw(screen, ">", fo, 496, 384, color.White)
            default:
                log.Fatal("Out of bounds (2592)")
            }
            switch option0 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 512, 64, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 512, 64, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 512, 64, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 512, 64, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 512, 64, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 512, 64, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 512, 64, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 512, 64, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 512, 64, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 512, 64, color.White)
            default:
                log.Fatal("Out of bounds (2720)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 512, 96, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 512, 96, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 512, 96, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 512, 96, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 512, 96, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 512, 96, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 512, 96, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 512, 96, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 512, 96, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 512, 96, color.White)
            default:
                log.Fatal("Out of bounds (2744)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 512, 128, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 512, 128, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 512, 128, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 512, 128, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 512, 128, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 512, 128, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 512, 128, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 512, 128, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 512, 128, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 512, 128, color.White)
            default:
                log.Fatal("Out of bounds (2768)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 512, 192, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 512, 192, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 512, 192, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 512, 192, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 512, 192, color.White)
            case 5:
                text.Draw(screen, "History", fo, 512, 192, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 512, 192, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 512, 192, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 512, 192, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 512, 192, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 512, 192, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 512, 192, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 512, 192, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 512, 192, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 512, 192, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 512, 192, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 512, 192, color.White)
            case 17:
                text.Draw(screen, "Acrobics", fo, 512, 192, color.White)
            default:
                log.Fatal("Out of bounds (2808)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 512, 224, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 512, 224, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 512, 224, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 512, 224, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 512, 224, color.White)
            case 5:
                text.Draw(screen, "History", fo, 512, 224, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 512, 224, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 512, 224, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 512, 224, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 512, 224, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 512, 224, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 512, 224, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 512, 224, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 512, 224, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 512, 224, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 512, 224, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 512, 224, color.White)
            case 17:
                text.Draw(screen, "Acrobics", fo, 512, 224, color.White)
            default:
                log.Fatal("Out of bounds (2848)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 512, 256, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 512, 256, color.White)
            case 2:
                text.Draw(screen, "Arcana", fo, 512, 256, color.White)
            case 3:
                text.Draw(screen, "Athletics", fo, 512, 256, color.White)
            case 4:
                text.Draw(screen, "Deception", fo, 512, 256, color.White)
            case 5:
                text.Draw(screen, "History", fo, 512, 256, color.White)
            case 6:
                text.Draw(screen, "Insight", fo, 512, 256, color.White)
            case 7:
                text.Draw(screen, "Intimidation", fo, 512, 256, color.White)
            case 8:
                text.Draw(screen, "Investigation", fo, 512, 256, color.White)
            case 9:
                text.Draw(screen, "Medicine", fo, 512, 256, color.White)
            case 10:
                text.Draw(screen, "Nature", fo, 512, 256, color.White)
            case 11:
                text.Draw(screen, "Perception", fo, 512, 256, color.White)
            case 12:
                text.Draw(screen, "Performance", fo, 512, 256, color.White)
            case 13:
                text.Draw(screen, "Persuasion", fo, 512, 256, color.White)
            case 14:
                text.Draw(screen, "Religion", fo, 512, 256, color.White)
            case 15:
                text.Draw(screen, "Sleight of Hand", fo, 512, 256, color.White)
            case 16:
                text.Draw(screen, "Stealth", fo, 512, 256, color.White)
            case 17:
                text.Draw(screen, "Acrobics", fo, 512, 256, color.White)
            default:
                log.Fatal("Out of bounds (2888)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Rapier", fo, 512, 320, color.White)
            case 1:
                text.Draw(screen, "Longsword", fo, 512, 320, color.White)
            case 2:
                text.Draw(screen, "Club", fo, 512, 320, color.White)
            case 3:
                text.Draw(screen, "Dagger", fo, 512, 320, color.White)
            case 4:
                text.Draw(screen, "Greatclub", fo, 512, 320, color.White)
            case 5:
                text.Draw(screen, "Handaxe", fo, 512, 320, color.White)
            case 6:
                text.Draw(screen, "Javelin", fo, 512, 320, color.White)
            case 7:
                text.Draw(screen, "Light hammer", fo, 512, 320, color.White)
            case 8:
                text.Draw(screen, "Mace", fo, 512, 320, color.White)
            case 9:
                text.Draw(screen, "Quarterstaff", fo, 512, 320, color.White)
            case 10:
                text.Draw(screen, "Sickle", fo, 512, 320, color.White)
            case 11:
                text.Draw(screen, "Spear", fo, 512, 320, color.White)
            case 12:
                text.Draw(screen, "Light crossbow", fo, 512, 320, color.White)
            case 13:
                text.Draw(screen, "Dart", fo, 512, 320, color.White)
            case 14:
                text.Draw(screen, "Shortbow", fo, 512, 320, color.White)
            case 15:
                text.Draw(screen, "Sling", fo, 512, 320, color.White)
            default:
                log.Fatal("Out of bounds (2924)")
            }
            switch option7 {
            case 0:
                text.Draw(screen, "Diplomat's Pack", fo, 512, 352, color.White)
            case 1:
                text.Draw(screen, "Entertainer's Pack", fo, 512, 352, color.White)
            default:
                log.Fatal("Out of bounds (2932)")
            }
            switch option8 {
            case 0:
                text.Draw(screen, "Bagpipes", fo, 512, 384, color.White)
            case 1:
                text.Draw(screen, "Drum", fo, 512, 384, color.White)
            case 2:
                text.Draw(screen, "Dulcimer", fo, 512, 384, color.White)
            case 3:
                text.Draw(screen, "Flute", fo, 512, 384, color.White)
            case 4:
                text.Draw(screen, "Lute", fo, 512, 384, color.White)
            case 5:
                text.Draw(screen, "Lyre", fo, 512, 384, color.White)
            case 6:
                text.Draw(screen, "Horn", fo, 512, 384, color.White)
            case 7:
                text.Draw(screen, "Pan flute", fo, 512, 384, color.White)
            case 8:
                text.Draw(screen, "Shawm", fo, 512, 384, color.White)
            case 9:
                text.Draw(screen, "Viol", fo, 512, 384, color.White)
            default:
                log.Fatal("Out of bounds (2956)")
            }
        case 2:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "History", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Medicine", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Persuasion", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (2973)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "History", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Medicine", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Persuasion", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (2987)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Mace", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Warhammer", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (2995)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Scale mail", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Leather armor", fo, 128, 160, color.White)
            case 2:
                text.Draw(screen, "Chain mail", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3005)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 128, 192, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 192, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 192, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 192, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 192, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 192, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 192, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 192, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 192, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 192, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 192, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 128, 192, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 128, 192, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 128, 192, color.White)
            default:
                log.Fatal("Out of bounds (3037)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Priest's Pack", fo, 128, 224, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 128, 224, color.White)
            default:
                log.Fatal("Out of bounds (3045)")
            }
        case 3:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 128, 64, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 128, 64, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3068)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 128, 96, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 128, 96, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3088)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Wooden shield", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 128, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 128, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 128, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 128, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 128, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 128, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 128, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 128, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 128, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 128, color.White)
            case 11:
                text.Draw(screen, "Light crossbow", fo, 128, 128, color.White)
            case 12:
                text.Draw(screen, "Dart", fo, 128, 128, color.White)
            case 13:
                text.Draw(screen, "Shortbow", fo, 128, 128, color.White)
            case 14:
                text.Draw(screen, "Sling", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3122)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Scimitar", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 160, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 160, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 160, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 160, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 160, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 160, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 160, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 160, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 160, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3148)")
            }
        case 4:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Athletics", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "History", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Intimidation", fo, 128, 64, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 128, 64, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3171)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Animal Handling", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Athletics", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "History", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Intimidation", fo, 128, 96, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 128, 96, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3191)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Chain mail", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Leather armor + longbow", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3199)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Battleaxe", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Flail", fo, 128, 160, color.White)
            case 2:
                text.Draw(screen, "Glaive", fo, 128, 160, color.White)
            case 3:
                text.Draw(screen, "Greataxe", fo, 128, 160, color.White)
            case 4:
                text.Draw(screen, "Greatsword", fo, 128, 160, color.White)
            case 5:
                text.Draw(screen, "Halberd", fo, 128, 160, color.White)
            case 6:
                text.Draw(screen, "Lance", fo, 128, 160, color.White)
            case 7:
                text.Draw(screen, "Longsword", fo, 128, 160, color.White)
            case 8:
                text.Draw(screen, "Maul", fo, 128, 160, color.White)
            case 9:
                text.Draw(screen, "Morningstar", fo, 128, 160, color.White)
            case 10:
                text.Draw(screen, "Pike", fo, 128, 160, color.White)
            case 11:
                text.Draw(screen, "Rapier", fo, 128, 160, color.White)
            case 12:
                text.Draw(screen, "Scimitar", fo, 128, 160, color.White)
            case 13:
                text.Draw(screen, "Shortsword", fo, 128, 160, color.White)
            case 14:
                text.Draw(screen, "Trident", fo, 128, 160, color.White)
            case 15:
                text.Draw(screen, "War pick", fo, 128, 160, color.White)
            case 16:
                text.Draw(screen, "Warhammer", fo, 128, 160, color.White)
            case 17:
                text.Draw(screen, "Whip", fo, 128, 160, color.White)
            case 18:
                text.Draw(screen, "Blowgun", fo, 128, 160, color.White)
            case 19:
                text.Draw(screen, "Hand crossbow", fo, 128, 160, color.White)
            case 20:
                text.Draw(screen, "Heavy crossbow", fo, 128, 160, color.White)
            case 21:
                text.Draw(screen, "Longbow", fo, 128, 160, color.White)
            case 22:
                text.Draw(screen, "Net", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3249)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Shield", fo, 128, 192, color.White)
            case 1:
                text.Draw(screen, "Battleaxe", fo, 128, 192, color.White)
            case 2:
                text.Draw(screen, "Flail", fo, 128, 192, color.White)
            case 3:
                text.Draw(screen, "Glaive", fo, 128, 192, color.White)
            case 4:
                text.Draw(screen, "Greataxe", fo, 128, 192, color.White)
            case 5:
                text.Draw(screen, "Greatsword", fo, 128, 192, color.White)
            case 6:
                text.Draw(screen, "Halberd", fo, 128, 192, color.White)
            case 7:
                text.Draw(screen, "Lance", fo, 128, 192, color.White)
            case 8:
                text.Draw(screen, "Longsword", fo, 128, 192, color.White)
            case 9:
                text.Draw(screen, "Maul", fo, 128, 192, color.White)
            case 10:
                text.Draw(screen, "Morningstar", fo, 128, 192, color.White)
            case 11:
                text.Draw(screen, "Pike", fo, 128, 192, color.White)
            case 12:
                text.Draw(screen, "Rapier", fo, 128, 192, color.White)
            case 13:
                text.Draw(screen, "Scimitar", fo, 128, 192, color.White)
            case 14:
                text.Draw(screen, "Shortsword", fo, 128, 192, color.White)
            case 15:
                text.Draw(screen, "Trident", fo, 128, 192, color.White)
            case 16:
                text.Draw(screen, "War pick", fo, 128, 192, color.White)
            case 17:
                text.Draw(screen, "Warhammer", fo, 128, 192, color.White)
            case 18:
                text.Draw(screen, "Whip", fo, 128, 192, color.White)
            case 19:
                text.Draw(screen, "Blowgun", fo, 128, 192, color.White)
            case 20:
                text.Draw(screen, "Hand crossbow", fo, 128, 192, color.White)
            case 21:
                text.Draw(screen, "Heavy crossbow", fo, 128, 192, color.White)
            case 22:
                text.Draw(screen, "Longbow", fo, 128, 192, color.White)
            case 23:
                text.Draw(screen, "Net", fo, 128, 192, color.White)
            default:
                log.Fatal("Out of bounds (3301)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 128, 224, color.White)
            case 1:
                text.Draw(screen, "Two handaxes", fo, 128, 224, color.White)
            default:
                log.Fatal("Out of bounds (3309)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 128, 256, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 128, 256, color.White)
            default:
                log.Fatal("Out of bounds (3317)")
            }
        case 5:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "History", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Stealth", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3336)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "History", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Religion", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Stealth", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3352)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Shortsword", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 128, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 128, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 128, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 128, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 128, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 128, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 128, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 128, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 128, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 128, color.White)
            case 11:
                text.Draw(screen, "Light crossbow", fo, 128, 128, color.White)
            case 12:
                text.Draw(screen, "Dart", fo, 128, 128, color.White)
            case 13:
                text.Draw(screen, "Shortbow", fo, 128, 128, color.White)
            case 14:
                text.Draw(screen, "Sling", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3386)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3394)")
            }
        case 6:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Athletics", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3413)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Athletics", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Intimidation", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Medicine", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3429)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Battleaxe", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Flail", fo, 128, 128, color.White)
            case 2:
                text.Draw(screen, "Glaive", fo, 128, 128, color.White)
            case 3:
                text.Draw(screen, "Greataxe", fo, 128, 128, color.White)
            case 4:
                text.Draw(screen, "Greatsword", fo, 128, 128, color.White)
            case 5:
                text.Draw(screen, "Halberd", fo, 128, 128, color.White)
            case 6:
                text.Draw(screen, "Lance", fo, 128, 128, color.White)
            case 7:
                text.Draw(screen, "Longsword", fo, 128, 128, color.White)
            case 8:
                text.Draw(screen, "Maul", fo, 128, 128, color.White)
            case 9:
                text.Draw(screen, "Morningstar", fo, 128, 128, color.White)
            case 10:
                text.Draw(screen, "Pike", fo, 128, 128, color.White)
            case 11:
                text.Draw(screen, "Rapier", fo, 128, 128, color.White)
            case 12:
                text.Draw(screen, "Scimitar", fo, 128, 128, color.White)
            case 13:
                text.Draw(screen, "Shortsword", fo, 128, 128, color.White)
            case 14:
                text.Draw(screen, "Trident", fo, 128, 128, color.White)
            case 15:
                text.Draw(screen, "War pick", fo, 128, 128, color.White)
            case 16:
                text.Draw(screen, "Warhammer", fo, 128, 128, color.White)
            case 17:
                text.Draw(screen, "Whip", fo, 128, 128, color.White)
            case 18:
                text.Draw(screen, "Blowgun", fo, 128, 128, color.White)
            case 19:
                text.Draw(screen, "Hand crossbow", fo, 128, 128, color.White)
            case 20:
                text.Draw(screen, "Heavy crossbow", fo, 128, 128, color.White)
            case 21:
                text.Draw(screen, "Longbow", fo, 128, 128, color.White)
            case 22:
                text.Draw(screen, "Net", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3479)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Shield", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Battleaxe", fo, 128, 160, color.White)
            case 2:
                text.Draw(screen, "Flail", fo, 128, 160, color.White)
            case 3:
                text.Draw(screen, "Glaive", fo, 128, 160, color.White)
            case 4:
                text.Draw(screen, "Greataxe", fo, 128, 160, color.White)
            case 5:
                text.Draw(screen, "Greatsword", fo, 128, 160, color.White)
            case 6:
                text.Draw(screen, "Halberd", fo, 128, 160, color.White)
            case 7:
                text.Draw(screen, "Lance", fo, 128, 160, color.White)
            case 8:
                text.Draw(screen, "Longsword", fo, 128, 160, color.White)
            case 9:
                text.Draw(screen, "Maul", fo, 128, 160, color.White)
            case 10:
                text.Draw(screen, "Morningstar", fo, 128, 160, color.White)
            case 11:
                text.Draw(screen, "Pike", fo, 128, 160, color.White)
            case 12:
                text.Draw(screen, "Rapier", fo, 128, 160, color.White)
            case 13:
                text.Draw(screen, "Scimitar", fo, 128, 160, color.White)
            case 14:
                text.Draw(screen, "Shortsword", fo, 128, 160, color.White)
            case 15:
                text.Draw(screen, "Trident", fo, 128, 160, color.White)
            case 16:
                text.Draw(screen, "War pick", fo, 128, 160, color.White)
            case 17:
                text.Draw(screen, "Warhammer", fo, 128, 160, color.White)
            case 18:
                text.Draw(screen, "Whip", fo, 128, 160, color.White)
            case 19:
                text.Draw(screen, "Blowgun", fo, 128, 160, color.White)
            case 20:
                text.Draw(screen, "Hand crossbow", fo, 128, 160, color.White)
            case 21:
                text.Draw(screen, "Heavy crossbow", fo, 128, 160, color.White)
            case 22:
                text.Draw(screen, "Longbow", fo, 128, 160, color.White)
            case 23:
                text.Draw(screen, "Net", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3531)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Five javelins", fo, 128, 192, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 192, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 192, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 192, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 192, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 192, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 192, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 192, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 192, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 192, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 192, color.White)
            default:
                log.Fatal("Out of bounds (3557)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Priest's Pack", fo, 128, 224, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 128, 224, color.White)
            default:
                log.Fatal("Out of bounds (3565)")
            }
        case 7:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 160, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 128, 64, color.White)
            case 6:
                text.Draw(screen, "Stealth", fo, 128, 64, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3588)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 128, 96, color.White)
            case 6:
                text.Draw(screen, "Stealth", fo, 128, 96, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3608)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Animal Handling", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 128, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 128, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 128, 128, color.White)
            case 4:
                text.Draw(screen, "Nature", fo, 128, 128, color.White)
            case 5:
                text.Draw(screen, "Perception", fo, 128, 128, color.White)
            case 6:
                text.Draw(screen, "Stealth", fo, 128, 128, color.White)
            case 7:
                text.Draw(screen, "Survival", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3628)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Scale mail", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Leather armor", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3636)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Shortsword", fo, 128, 192, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 192, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 192, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 192, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 192, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 192, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 192, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 192, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 192, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 192, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 192, color.White)
            default:
                log.Fatal("Out of bounds (3662)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Shortsword", fo, 128, 224, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 224, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 224, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 224, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 224, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 224, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 224, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 224, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 224, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 224, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 224, color.White)
            default:
                log.Fatal("Out of bounds (3688)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 128, 256, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 128, 256, color.White)
            default:
                log.Fatal("Out of bounds (3696)")
            }
        case 8:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 192, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 128, 64, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 128, 64, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 128, 64, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 128, 64, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 128, 64, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3725)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 128, 96, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 128, 96, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 128, 96, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 128, 96, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 128, 96, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3751)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 128, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 128, 128, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 128, 128, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 128, 128, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 128, 128, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 128, 128, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 128, 128, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 128, 128, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 128, 128, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3777)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Acrobatics", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Athletics", fo, 128, 160, color.White)
            case 2:
                text.Draw(screen, "Deception", fo, 128, 160, color.White)
            case 3:
                text.Draw(screen, "Insight", fo, 128, 160, color.White)
            case 4:
                text.Draw(screen, "Intimidation", fo, 128, 160, color.White)
            case 5:
                text.Draw(screen, "Investigation", fo, 128, 160, color.White)
            case 6:
                text.Draw(screen, "Perception", fo, 128, 160, color.White)
            case 7:
                text.Draw(screen, "Performance", fo, 128, 160, color.White)
            case 8:
                text.Draw(screen, "Persuasion", fo, 128, 160, color.White)
            case 9:
                text.Draw(screen, "Sleight of Hand", fo, 128, 160, color.White)
            case 10:
                text.Draw(screen, "Stealth", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3803)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Rapier", fo, 128, 192, color.White)
            case 1:
                text.Draw(screen, "Shortsword", fo, 128, 192, color.White)
            default:
                log.Fatal("Out of bounds (3811)")
            }
            switch option5 {
            case 0:
                text.Draw(screen, "Shortbow", fo, 128, 224, color.White)
            case 1:
                text.Draw(screen, "Shortsword", fo, 128, 224, color.White)
            default:
                log.Fatal("Out of bounds (3819)")
            }
            switch option6 {
            case 0:
                text.Draw(screen, "Burglar's Pack", fo, 128, 256, color.White)
            case 1:
                text.Draw(screen, "Dungeoneer's Pack", fo, 128, 256, color.White)
            case 2:
                text.Draw(screen, "Explorer's Pack", fo, 128, 256, color.White)
            default:
                log.Fatal("Out of bounds (3829)")
            }
        case 9:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3848)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Persuasion", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3864)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 128, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 128, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 128, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 128, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 128, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 128, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 128, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 128, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 128, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 128, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 128, 128, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 128, 128, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3896)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Component pouch", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Arcane focus", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3904)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Dungeoneer's Pack", fo, 128, 192, color.White)
            case 1:
                text.Draw(screen, "Explorer's Pack", fo, 128, 192, color.White)
            default:
                log.Fatal("Out of bounds (3912)")
            }
        case 10:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "History", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Investigation", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Nature", fo, 128, 64, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (3933)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "Deception", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "History", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Intimidation", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Investigation", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Nature", fo, 128, 96, color.White)
            case 6:
                text.Draw(screen, "Religion", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (3951)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Light crossbow", fo, 128, 128, color.White)
            case 1:
                text.Draw(screen, "Club", fo, 128, 128, color.White)
            case 2:
                text.Draw(screen, "Dagger", fo, 128, 128, color.White)
            case 3:
                text.Draw(screen, "Greatclub", fo, 128, 128, color.White)
            case 4:
                text.Draw(screen, "Handaxe", fo, 128, 128, color.White)
            case 5:
                text.Draw(screen, "Javelin", fo, 128, 128, color.White)
            case 6:
                text.Draw(screen, "Light hammer", fo, 128, 128, color.White)
            case 7:
                text.Draw(screen, "Mace", fo, 128, 128, color.White)
            case 8:
                text.Draw(screen, "Quarterstaff", fo, 128, 128, color.White)
            case 9:
                text.Draw(screen, "Sickle", fo, 128, 128, color.White)
            case 10:
                text.Draw(screen, "Spear", fo, 128, 128, color.White)
            case 11:
                text.Draw(screen, "Dart", fo, 128, 128, color.White)
            case 12:
                text.Draw(screen, "Shortbow", fo, 128, 128, color.White)
            case 13:
                text.Draw(screen, "Sling", fo, 128, 128, color.White)
            default:
                log.Fatal("Out of bounds (3983)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Component pouch", fo, 128, 160, color.White)
            case 1:
                text.Draw(screen, "Arcane focus", fo, 128, 160, color.White)
            default:
                log.Fatal("Out of bounds (3991)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Scholar's Pack", fo, 128, 192, color.White)
            case 1:
                text.Draw(screen, "Dungeoneer's Pack", fo, 128, 192, color.White)
            default:
                log.Fatal("Out of bounds (3999)")
            }
        case 11:
            text.Draw(screen, "Skill Proficiencies:", fo, 64, 64, color.White)
            text.Draw(screen, "Equipment:", fo, 64, 128, color.White)
            switch option0 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "History", fo, 128, 64, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 64, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 128, 64, color.White)
            case 4:
                text.Draw(screen, "Medicine", fo, 128, 64, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (4018)")
            }
            switch option1 {
            case 0:
                text.Draw(screen, "Arcana", fo, 128, 96, color.White)
            case 1:
                text.Draw(screen, "History", fo, 128, 96, color.White)
            case 2:
                text.Draw(screen, "Insight", fo, 128, 96, color.White)
            case 3:
                text.Draw(screen, "Investigation", fo, 128, 96, color.White)
            case 4:
                text.Draw(screen, "Medicine", fo, 128, 96, color.White)
            case 5:
                text.Draw(screen, "Religion", fo, 128, 96, color.White)
            default:
                log.Fatal("Out of bounds (4034)")
            }
            switch option2 {
            case 0:
                text.Draw(screen, "Quarterstaff", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Dagger", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (4042)")
            }
            switch option3 {
            case 0:
                text.Draw(screen, "Component pouch", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Arcane focus", fo, 128, 64, color.White)
            default:
                log.Fatal("Out of bounds (4050)")
            }
            switch option4 {
            case 0:
                text.Draw(screen, "Scholar's pack", fo, 128, 64, color.White)
            case 1:
                text.Draw(screen, "Explorer's pack", fo, 128, 64, color.White)
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

    backgroundmap[0] = "Acolyte"
    backgroundmap[1] = "Charlatan"
    backgroundmap[2] = "Criminal"
    backgroundmap[3] = "Entertainer"
    backgroundmap[4] = "Folk Hero"
    backgroundmap[5] = "Guild Artisan"
    backgroundmap[6] = "Hermit"
    backgroundmap[7] = "Noble"
    backgroundmap[8] = "Outlander"
    backgroundmap[9] = "Sage"
    backgroundmap[10] = "Sailor"
    backgroundmap[11] = "Soldier"
    backgroundmap[12] = "Urchin"

    equipmentmap[0] = "Pack 1"
    equipmentmap[1] = "Pack 2"

    savesTableSchema = []string{"name,TEXT,1,null,1", "level,TEXT,1,\"One\",0", "x,INT,1,null,0", "y,INT,1,null,0", "csdone,TEXT,0,null,0", "inventory,TEXT,0,null,0"}
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
