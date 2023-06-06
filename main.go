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
    "sort"
    "strconv"
    "strings"
    "time"

    "golang.org/x/image/font"

    "github.com/jsnider-mtu/quailgame/assets"
    "github.com/jsnider-mtu/quailgame/classes"
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
    invmenu bool = false
    charsheet0 bool = false
    charsheet1 bool = false
    charsheet2 bool = false
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
    blankImage *ebiten.Image
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
    //fon *truetype.Font
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
    pagesTableSchema []string
    schemaRowsCount int = 0
    colsStr string
    animCount int = 0
    icon16img image.Image
    icon32img image.Image
    icon48img image.Image
    creation bool = false
    creationsel int = 0
    creationpage [4]int
    str int
    dex int
    con int
    intel int
    wis int
    cha int
    strmod int
    dexmod int
    conmod int
    intelmod int
    wismod int
    chamod int
    pb int
    //size int // 0: Small, 1: Medium, 2: Large
    darkvision bool = false
    lucky bool = false
    nimbleness bool = false
    brave bool = false
    ancestry string
    targeted int = -1
    passattempts int = 0
    npchp string
    levelslice []*levels.Level
    lvlloaded bool = false
    invsel int = 0
    invsel2 int = 0
    invselmenu bool = false
    invselitem inventory.Item
    timestart time.Time
    nextturn bool = false
    effectact string
    effectmsg bool = false
    countend int = 0
    overflowcur int = 0
    overflownum int = 0
    pageind int = 0
    pageexists bool = false
    maxw int = 0
    maxh int = 0
    numlines int = 0
    r image.Rectangle
    r2 image.Rectangle
    r3 image.Rectangle
    maxlines int = 0
    readgm ebiten.GeoM
    readimg *ebiten.Image
    readgm2 ebiten.GeoM
    readimg2 *ebiten.Image
    readgm3 ebiten.GeoM
    readimg3 *ebiten.Image
    readimgopt *ebiten.DrawImageOptions
    moreshown bool = false
    hei int = 0
    wid int = 0
    hei2 int = 0
    wid2 int = 0
    hei3 int = 0
    wid3 int = 0
    dagm ebiten.GeoM
    ismgm ebiten.GeoM
    ismimg *ebiten.Image
    ismgm2 ebiten.GeoM
    ismimg2 *ebiten.Image
    actcheck string
    dur time.Duration
    effgm ebiten.GeoM
    effimg *ebiten.Image
    effgm2 ebiten.GeoM
    effimg2 *ebiten.Image
    owgm ebiten.GeoM
    iw int = 0
    pausegm ebiten.GeoM
    pauseimg *ebiten.Image
    pausegm2 ebiten.GeoM
    pauseimg2 *ebiten.Image
    npchporig string
    npchpslice []string
    npchporigval string
    npchpupdate bool = false
    paperind int = -1
    throwtarget [2]int
    targetedBoxVert *ebiten.Image
    targetedBoxHoriz *ebiten.Image
    throwTargetBoxHoriz *ebiten.Image
    throwTargetBoxVert *ebiten.Image
    shortrange int = 0
    longrange int = 0
    gainProfST bool = false
    gainprofstsel int = 0
    gainProfSkill bool = false
    gainprofskillsel int = 0
    stprofexists string
    skillprofexists string
)

var abilities = make([]int, 6)
var lines = make([]string, 0)
var pages = make([]*items.Page, 0)
var savingthrows = make(map[string]int)
var languages = make([]string, 0)
var proficiencies = make([]string, 0)
var resistances = make([]string, 0)
var spellsslice = make([]string, 0)
var c1 = make(chan int)

var classessli = []string{
    "Quail", "Bard", "Cleric",
    "Druid", "Fighter", "Monk",
    "Paladin", "Ranger", "Rogue",
    "Sorceror", "Warlock", "Wizard"}

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
                    _, err = db.Exec("delete from pages where charname = ?", loads[loadsel][0])
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
                        creationsel = 0
                        creation = true
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
                    Input(&sb)
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
                        name = strings.Trim(name, "\n")
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
                        targeted = -1
                        p.Name = name
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
                        cutscene = false
                        creationsel = 0
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
                            levelslice = append(levelslice, l)
                            targeted = -1
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
            creation = false
            start = true
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
            creationsel--
            if creationsel < 0 {
                creationsel = 11
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
            if creationsel > 5 {
                creationsel -= 6
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
            if creationsel < 6 {
                creationsel += 6
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
            creationsel++
            if creationsel > 11 {
                creationsel = 0
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
            p.Inv.Clear()
            curCS = 0
            csDone = make([]int, 0)
            onescore := make([]int, 4)
            for x := 0; x < 6; x++ {
                for a := 0; a < 4; a++ {
                    onescore[a] = rand.Intn(6) + 1
                }
                sort.Slice(onescore, func(i, j int) bool {
                    return onescore[i] > onescore[j]
                })
                score := onescore[0] + onescore[1] + onescore[2]
                abilities[x] = score
            }
            sort.Slice(abilities, func(i, j int) bool {
                return abilities[i] > abilities[j]
            })
            switch creationsel {
            case 0:
                p.Class = &classes.Quail{}
            default:
                return errors.New("Invalid value for creationsel")
            }
            var abilarr [6]int
            copy(abilarr[:], abilities)
            if cc := p.Class.Create(abilarr); !cc {
                panic("Character creation failed")
            }
            p.Equipment = &player.Equipment{}
            creationsel = 0
            creation = false
        }
    } else if gainProfST {
        if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
            if gainprofstsel > 0 {
                gainprofstsel--
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
            if gainprofstsel < 5 {
                gainprofstsel++
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
            stprofexists = ""
            switch gainprofstsel {
            case 0:
                if success := p.Class.AddProf("str"); !success {
                    stprofexists = "strength"
                } else {
                    gainProfST = false
                }
            case 1:
                if success := p.Class.AddProf("dex"); !success {
                    stprofexists = "dexterity"
                } else {
                    gainProfST = false
                }
            case 2:
                if success := p.Class.AddProf("con"); !success {
                    stprofexists = "constitution"
                } else {
                    gainProfST = false
                }
            case 3:
                if success := p.Class.AddProf("intel"); !success {
                    stprofexists = "intelligence"
                } else {
                    gainProfST = false
                }
            case 4:
                if success := p.Class.AddProf("wis"); !success {
                    stprofexists = "wisdom"
                } else {
                    gainProfST = false
                }
            case 5:
                if success := p.Class.AddProf("cha"); !success {
                    stprofexists = "charisma"
                } else {
                    gainProfST = false
                }
            default:
                log.Fatal(fmt.Sprintf("%d is not a valid value for gainprofstsel", gainprofstsel))
            }
        }
    } else if gainProfSkill {
        if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
            if gainprofskillsel > 0 {
                gainprofskillsel--
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
            if gainprofskillsel < 17 {
                gainprofskillsel++
            }
        }
        if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
            skillprofexists = ""
            switch gainprofskillsel {
            case 0:
                if success := p.Class.AddProf("acrobatics"); !success {
                    skillprofexists = "acrobatics"
                } else {
                    gainProfSkill = false
                }
            case 1:
                if success := p.Class.AddProf("animalhandling"); !success {
                    skillprofexists = "animalhandling"
                } else {
                    gainProfSkill = false
                }
            case 2:
                if success := p.Class.AddProf("arcana"); !success {
                    skillprofexists = "arcana"
                } else {
                    gainProfSkill = false
                }
            case 3:
                if success := p.Class.AddProf("athletics"); !success {
                    skillprofexists = "athletics"
                } else {
                    gainProfSkill = false
                }
            case 4:
                if success := p.Class.AddProf("deception"); !success {
                    skillprofexists = "deception"
                } else {
                    gainProfSkill = false
                }
            case 5:
                if success := p.Class.AddProf("history"); !success {
                    skillprofexists = "history"
                } else {
                    gainProfSkill = false
                }
            case 6:
                if success := p.Class.AddProf("insight"); !success {
                    skillprofexists = "insight"
                } else {
                    gainProfSkill = false
                }
            case 7:
                if success := p.Class.AddProf("intimidation"); !success {
                    skillprofexists = "intimidation"
                } else {
                    gainProfSkill = false
                }
            case 8:
                if success := p.Class.AddProf("investigation"); !success {
                    skillprofexists = "investigation"
                } else {
                    gainProfSkill = false
                }
            case 9:
                if success := p.Class.AddProf("medicine"); !success {
                    skillprofexists = "medicine"
                } else {
                    gainProfSkill = false
                }
            case 10:
                if success := p.Class.AddProf("nature"); !success {
                    skillprofexists = "nature"
                } else {
                    gainProfSkill = false
                }
            case 11:
                if success := p.Class.AddProf("perception"); !success {
                    skillprofexists = "perception"
                } else {
                    gainProfSkill = false
                }
            case 12:
                if success := p.Class.AddProf("performance"); !success {
                    skillprofexists = "performance"
                } else {
                    gainProfSkill = false
                }
            case 13:
                if success := p.Class.AddProf("persuasion"); !success {
                    skillprofexists = "persuasion"
                } else {
                    gainProfSkill = false
                }
            case 14:
                if success := p.Class.AddProf("religion"); !success {
                    skillprofexists = "religion"
                } else {
                    gainProfSkill = false
                }
            case 15:
                if success := p.Class.AddProf("sleightofhand"); !success {
                    skillprofexists = "sleight of hand"
                } else {
                    gainProfSkill = false
                }
            case 16:
                if success := p.Class.AddProf("stealth"); !success {
                    skillprofexists = "stealth"
                } else {
                    gainProfSkill = false
                }
            case 17:
                if success := p.Class.AddProf("survival"); !success {
                    skillprofexists = "survival"
                } else {
                    gainProfSkill = false
                }
            default:
                log.Fatal(fmt.Sprintf("%d is not a valid value for gainprofskillsel", gainprofskillsel))
            }
        }
    } else {
        if inpututil.IsKeyJustPressed(ebiten.KeyY) {
            if leveled := p.Class.EarnXP(300); leveled {
                levelUp(p)
            }
        }
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
                    overworld = false
                    invmenu = false
                    charsheet0 = false
                    charsheet1 = false
                    charsheet2 = false
                    pause = false
                case 2:
                    start = true
                    loads = [][2]string{}
                    loadsfound = false
                    findloads = true
                    overworld = false
                    invmenu = false
                    charsheet0 = false
                    charsheet1 = false
                    charsheet2 = false
                    pause = false
                case 3:
                    os.Exit(0)
                }
            }
        } else {
            if timestart.IsZero() {
                log.Println("Starting the clock")
                timestart = time.Now()
            }
            t := time.Now()
            dur := t.Sub(timestart)
            if (dur / 1000000000) % 6 == 0 {
                if nextturn {
                    log.Println("Next turn")
                    nextturn = false
                }
            } else {
                nextturn = true
            }
            if effectact == "write" {
                Input(&sb)
                if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
                    p.WriteMsg = sb.String()
                    c1 <- 0
                    sb.Reset()
                    effectmsg = false
                    effectact = ""
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyC) {
                    if inpututil.KeyPressDuration(ebiten.KeyControlLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyControlRight) > 0 {
                        c1 <- 1
                        sb.Reset()
                        effectmsg = false
                        effectact = ""
                    }
                }
                return nil
            }
            if effectact == "read" {
                if paperind < 0 {
                    log.Println("paperind < 0 (Update)")
                    return nil
                }
                pages = p.Inv.GetItems()[paperind].(*items.Paper).GetPages()
                for _, pa := range pages {
                    pageexists = false
                    for _, pm := range p.PageMsgs {
                        if pa.GetName() == pm[3].(string) {
                            pageexists = true
                            break
                        }
                    }
                    if pageexists {
                        continue
                    }
                    lines = strings.Split(pa.Read(), "\n")
                    numlines = len(lines)
                    maxlines = (552 - 48) / 28
                    if !pageexists {
                        p.PageMsgs = append(p.PageMsgs, []interface{}{lines, numlines, maxlines, pa.GetName()})
                    }
                }
                if len(p.PageMsgs) > 0 {
                    overflownum = p.PageMsgs[pageind][1].(int) / p.PageMsgs[pageind][2].(int)
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                    if overflowcur < overflownum {
                        overflowcur++
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                    if overflowcur > 0 {
                        overflowcur--
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
                    pageind--
                    if pageind < 0 {
                        pageind = 0
                    }
                    overflowcur = 0
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
                    pageind++
                    if pageind > len(p.PageMsgs) - 1 {
                        pageind = len(p.PageMsgs) - 1
                    }
                    overflowcur = 0
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
                    effectmsg = false
                    effectact = ""
                    overflowcur = 0
                    overflownum = 0
                    pageind = 0
                }
            }
            if effectact == "throw" {
                if throwtarget == [2]int{} {
                    throwtarget = p.Pos
                }
                if shortrange == 0 && longrange == 0 {
                    shortrange = <-c1
                    longrange = <-c1
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.KeyPressDuration(ebiten.KeyLeft) % 4 == 3 || inpututil.KeyPressDuration(ebiten.KeyA) % 4 == 3 {
                    if throwtarget[0] - 24 > 0 && throwtarget[0] - 24 < l.GetMax()[0] {
                        if ok, _, _ := l.LineOfSight(p, [2]int{throwtarget[0] - 24, throwtarget[1]}); ok {
                            if l.Distance(p, [2]int{throwtarget[0] - 24, throwtarget[1]}) <= float64(longrange) {
                                throwtarget[0] -= 24
                            }
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.KeyPressDuration(ebiten.KeyRight) % 4 == 3 || inpututil.KeyPressDuration(ebiten.KeyD) % 4 == 3 {
                    if throwtarget[0] + 24 > 0 && throwtarget[0] + 24 < l.GetMax()[0] {
                        if ok, _, _ := l.LineOfSight(p, [2]int{throwtarget[0] + 24, throwtarget[1]}); ok {
                            if l.Distance(p, [2]int{throwtarget[0] + 24, throwtarget[1]}) <= float64(longrange) {
                                throwtarget[0] += 24
                            }
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.KeyPressDuration(ebiten.KeyUp) % 4 == 3 || inpututil.KeyPressDuration(ebiten.KeyW) % 4 == 3 {
                    if throwtarget[1] - 24 > 0 && throwtarget[1] - 24 < l.GetMax()[1] {
                        if ok, _, _ := l.LineOfSight(p, [2]int{throwtarget[0], throwtarget[1] - 24}); ok {
                            if l.Distance(p, [2]int{throwtarget[0], throwtarget[1] - 24}) <= float64(longrange) {
                                throwtarget[1] -= 24
                            }
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.KeyPressDuration(ebiten.KeyDown) % 4 == 3 || inpututil.KeyPressDuration(ebiten.KeyS) % 4 == 3 {
                    if throwtarget[1] + 24 > 0 && throwtarget[1] + 24 < l.GetMax()[1] {
                        if ok, _, _ := l.LineOfSight(p, [2]int{throwtarget[0], throwtarget[1] + 24}); ok {
                            if l.Distance(p, [2]int{throwtarget[0], throwtarget[1] + 24}) <= float64(longrange) {
                                throwtarget[1] += 24
                            }
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                    ready := <-c1
                    switch ready {
                    case -1:
                        log.Println("NPC was attacked")
                    case 0: // Oil Flask
                        l.OilSpot(throwtarget)
                        log.Println("Oil Flask thrown")
                    case 1: // Dagger
                        log.Println("Dagger thrown")
                    case 2: // Handaxe
                        log.Println("Handaxe thrown")
                    case 3: // Light Hammer
                        log.Println("Light Hammer thrown")
                    default:
                        log.Println(fmt.Sprintf("%d is invalid for throw", ready))
                    }
                    c1 <- ready
                    effectmsg = false
                    effectact = ""
                    shortrange = 0
                    longrange = 0
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
                    c1 <- 0
                    throwtarget = [2]int{}
                    effectmsg = false
                    effectact = ""
                    shortrange = 0
                    longrange = 0
                }
            }
            if effectact == "playmusic" {
                if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyQ) {
                    c1 <- 0
                    effectmsg = false
                    effectact = ""
                }
            }
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
                var statsstr string = p.Class.Save()
                //fmt.Println(statsstr)
                var equipmentstr string = p.Equipment.Save()
                _, err = db.Exec(saveStmt, name, l.GetName(), l.Pos[0], l.Pos[1], csdonestr, invstr, statsstr, equipmentstr)
                if err != nil {
                    log.Fatal(fmt.Sprintf("%q: %s\n", err, saveStmt))
                }
                if strings.Contains(invstr, "Paper") {
                    afterpaper := strings.Split(strings.Split(invstr, "Paper,")[1], ";")[0]
                    pagenames := strings.Split(afterpaper, ",")
                    pagenames = pagenames[:len(pagenames) - 1]
                    if len(pagenames) > 0 {
                        pagesSaveStmt := "insert or replace into pages (name, msg, charname) values (?, ?, ?);"
                        for itemind, item := range p.Inv.GetItems() {
                            if strings.HasPrefix(item.PrettyPrint(), "Paper") {
                                pages := p.Inv.GetItems()[itemind].(*items.Paper).GetPages()
                                for _, page := range pages {
                                    log.Println("Saving " + page.GetName())
                                    _, err = db.Exec(pagesSaveStmt, page.GetName(), page.Read(), name)
                                    if err != nil {
                                        log.Fatal(fmt.Sprintf("%q: %s\n", err, pagesSaveStmt))
                                    }
                                }
                            }
                        }
                    }
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
                var statsstr string
                var equipmentstr string
                for rows.Next() {
                    err = rows.Scan(&savename, &levelname, &x, &y, &csdonestr, &invstr, &statsstr, &equipmentstr)
                }
                err = rows.Err()
                if err != nil {
                    log.Fatal(err)
                }
                p.Name = savename
                p.Class = &classes.Quail{}
                var loadstatsarr [6]int
                var loadstatslvl int
                var loadstatsxp int
                log.Printf("Need to use loadstatsxp: %d", loadstatsxp)
                loadstatssli := strings.Split(statsstr, ";")
                for statind, statval := range loadstatssli {
                    if statval == "" {
                        break
                    }
                    statint, err := strconv.Atoi(statval)
                    if err != nil {
                        panic(err)
                    }
                    switch statind {
                    case 6:
                        loadstatslvl = statint
                    case 7:
                        loadstatsxp = statint
                    default:
                        loadstatsarr[statind] = statint
                    }
                }
                p.Class.Create(loadstatsarr)
                if loadstatslvl > 1 {
                    // Level up incrementally
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
                p.Inv = &inventory.Inv{}
                invstrarr := strings.Split(invstr, ";")
                for _, item := range invstrarr {
                    if item == "" {
                        break
                    }
                    itemprops := strings.Split(item, ",")
                    switch len(itemprops) {
                    case 1:
                        p.Inv.Add(items.LoadItem(itemprops[0], nil))
                    case 2:
                        p.Inv.Add(items.LoadItem(itemprops[0], itemprops[1]))
                    default:
                        if itemprops[0] == "Paper" {
                            p.Inv.Add(items.LoadItem(itemprops[0], itemprops[len(itemprops) - 1]))
                        } else if itemprops[0] == "Candles" {
                            p.Inv.Add(items.LoadItem(itemprops[0], itemprops[2]))
                            turns, err := strconv.Atoi(itemprops[1])
                            if err != nil {
                                return err
                            }
                            p.Inv.GetItems()[len(p.Inv.GetItems()) - 1].(*items.Candles).Turns = turns
                        } else if itemprops[0] == "Lamp" {
                            p.Inv.Add(items.LoadItem(itemprops[0], itemprops[2]))
                            turns, err := strconv.Atoi(itemprops[1])
                            if err != nil {
                                return err
                            }
                            p.Inv.GetItems()[len(p.Inv.GetItems()) - 1].(*items.Lamp).Turns = turns
                        } else {
                            return errors.New("Too many itemprops (inventory)")
                        }
                    }
                }
                p.Equipment = &player.Equipment{}
                if len(strings.Split(equipmentstr, "|")) > 1 {
                    for _, equipped := range strings.Split(equipmentstr, "|") {
                        itemprops := strings.Split(equipped, ",")
                        itemname := strings.Split(itemprops[0], "=")[1]
                        switch len(itemprops) {
                        case 1:
                            newitem := items.LoadItem(strings.Split(itemname, ";")[0], nil)
                            p.Inv.Add(newitem)
                            p.Equip(newitem)
                        case 2:
                            newitem := items.LoadItem(itemname, strings.Split(itemprops[1], ";")[0])
                            p.Inv.Add(newitem)
                            p.Equip(newitem)
                        default:
                            if itemname == "Candles" {
                                newitem := items.LoadItem(itemname, strings.Split(itemprops[2], ";")[0])
                                p.Inv.Add(newitem)
                                p.Equip(newitem)
                                turns, err := strconv.Atoi(itemprops[1])
                                if err != nil {
                                    return err
                                }
                                p.Equipment.LeftHand.(*items.Candles).Turns = turns
                            } else if itemname == "Lamp" {
                                newitem := items.LoadItem(itemname, strings.Split(itemprops[2], ";")[0])
                                p.Inv.Add(newitem)
                                p.Equip(newitem)
                                turns, err := strconv.Atoi(itemprops[1])
                                if err != nil {
                                    return err
                                }
                                p.Equipment.LeftHand.(*items.Lamp).Turns = turns
                            } else {
                                return errors.New("Too many itemprops (equipment)")
                            }
                        }
                    }
                }
                pageRows, err := db.Query("select * from pages where charname = ?", name)
                if err != nil {
                    log.Fatal(err)
                }
                defer pageRows.Close()
                var pagename string
                var pagemsg string
                var charname string
                for pageRows.Next() {
                    err = pageRows.Scan(&pagename, &pagemsg, &charname)
                    if pagemsg != "" {
                        for itemind, item := range p.Inv.GetItems() {
                            if strings.HasPrefix(item.PrettyPrint(), "Paper") {
                                pages := p.Inv.GetItems()[itemind].(*items.Paper).GetPages()
                                newpage := true
                                for _, p := range pages {
                                    if p.GetName() == pagename {
                                        newpage = false
                                        break
                                    }
                                }
                                if newpage {
                                    p.Inv.GetItems()[itemind].(*items.Paper).LoadPage(pagename, pagemsg)
                                }
                            }
                        }
                    }
                }
                l = levels.LoadLvl(levelname, 0, x, y)
                targeted = -1
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
            if invselmenu {
                if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                    if invsel2 > 0 {
                        if a := p.Inv.GetItems()[invsel].Action(); a == "" {
                            invsel2 -= 2
                        } else {
                            invsel2--
                        }
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                    if invsel2 < 2 {
                        if a := p.Inv.GetItems()[invsel].Action(); a == "" {
                            invsel2 += 2
                        } else {
                            invsel2++
                        }
                    }
                }
                if p.Inv.GetItems()[invsel].Slot() == "" {
                    if a := p.Inv.GetItems()[invsel].Action(); a == "" {
                        if invsel2 < 2 {
                            invsel2 = 2
                        }
                    } else {
                        if invsel2 < 1 {
                            invsel2 = 1
                        }
                    }
                } else {
                    if invsel2 < 0 {
                        invsel2 = 0
                    }
                }
                if invsel2 > 2 {
                    invsel2 = 2
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                    switch invsel2 {
                    case 0:
                        p.Equip(p.Inv.GetItems()[invsel])
                        invselmenu = false
                    case 1:
                        action, data := p.Inv.GetItems()[invsel].Use()
                        if action == "write" || action == "read" {
                            for in, i := range p.Inv.GetItems() {
                                if strings.HasPrefix(i.PrettyPrint(), "Paper") {
                                    paperind = in
                                    data = append(data, in)
                                }
                            }
                        } else if action == "illuminate" {
                            if data[0] == 5 && data[1] == 5 {
                                for _, i := range p.Inv.GetItems() {
                                    if strings.HasPrefix(i.PrettyPrint(), "Candles") {
                                        p.Equip(i)
                                        break
                                    }
                                }
                            } else if data[0] == 15 && data[1] == 30 {
                                for _, i := range p.Inv.GetItems() {
                                    if strings.HasPrefix(i.PrettyPrint(), "Lamp") {
                                        p.Equip(i)
                                        break
                                    }
                                }
                            }
                        } else if action == "throw" || action == "playmusic" {
                            data = append(data, invsel)
                        }
                        if action != "" {
                            effectact = action
                            effectmsg = true
                            go p.Effects(action, data, c1)
                        }
                        invselmenu = false
                        invmenu = false
                    case 2:
                        p.Inv.Drop(p.Inv.GetItems()[invsel])
                        invselmenu = false
                    default:
                        log.Fatal(fmt.Sprintf("Invalid value %d for invsel in invselmenu", invsel))
                    }
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyI) {
                    invselmenu = false
                }
            } else if invmenu {
                if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
                    invsel--
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
                    invsel++
                }
                if invsel < 0 {
                    invsel = 0
                } else if invsel > len(p.Inv.GetItems()) - 1 {
                    invsel = len(p.Inv.GetItems()) - 1
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
                    //invselitem = p.Inv.GetItems()[invsel]
                    invsel2 = 0
                    invselmenu = true
                }
                if inpututil.IsKeyJustPressed(ebiten.KeyI) {
                    invmenu = false
                }
            } else if inpututil.IsKeyJustPressed(ebiten.KeyI) && !overworld {
                charsheet0 = false
                charsheet1 = false
                charsheet2 = false
                invmenu = !invmenu
                invselmenu = false
            }
            if inpututil.IsKeyJustPressed(ebiten.KeyC) && !overworld {
                invmenu = false
                invselmenu = false
                if charsheet0 || charsheet1 || charsheet2 {
                    charsheet0 = false
                    charsheet1 = false
                    charsheet2 = false
                } else {
                    charsheet0 = true
                }
            }
            if charsheet0 && inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
                charsheet0 = false
                charsheet1 = true
                charsheet2 = false
                return nil
            }
            if charsheet1 && inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
                charsheet1 = false
                charsheet0 = true
                charsheet2 = false
                return nil
            }
            if charsheet1 && inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
                charsheet1 = false
                charsheet0 = false
                charsheet2 = true
                return nil
            }
            if charsheet2 && inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
                charsheet1 = true
                charsheet0 = false
                charsheet2 = false
                return nil
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
            } else if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
                targeted++
                if len(l.NPCs) == 0 {
                    targeted = -1
                } else if targeted == len(l.NPCs) {
                    targeted = 0
                }
            }
            if !dialogopen && !lvlchange && !start {
                for _, npc := range l.NPCs {
                    if npc.GetSpeed() > 0 && (npcCount + npc.GetOffset()) % npc.GetSpeed() == 0 {
                        npc.Stopped = false
                        switch rand.Intn(4) {
                        case 0:
                            npc.Direction = "down"
                            l.TryUpdatePos(false, npc.PC, true, 24, 0, p)
                        case 1:
                            npc.Direction = "up"
                            l.TryUpdatePos(false, npc.PC, true, -24, 0, p)
                        case 2:
                            npc.Direction = "right"
                            l.TryUpdatePos(false, npc.PC, false, 24, 0, p)
                        case 3:
                            npc.Direction = "left"
                            l.TryUpdatePos(false, npc.PC, false, -24, 0, p)
                        }
                    } else if !npc.Stopped && (npcCount + npc.GetOffset() - 4) % npc.GetSpeed() == 0 {
                        npc.Stopped = true
                    }
                }
                if effectact != "throw" {
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
                            ok, blocker := l.TryUpdatePos(true, p, true, -24, passattempts, p)
                            if ok {
                                for _, a := range l.Doors {
                                    if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
                                        newlvl = a.NewLvl
                                        lvlchange = true
                                    }
                                }
                            } else {
                                if blocker == "npc" {
                                    passattempts++
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
                            ok, blocker := l.TryUpdatePos(true, p, false, -24, passattempts, p)
                            if ok {
                                for _, a := range l.Doors {
                                    if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
                                        newlvl = a.NewLvl
                                        lvlchange = true
                                    }
                                }
                            } else {
                                if blocker == "npc" {
                                    passattempts++
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
                            ok, blocker := l.TryUpdatePos(true, p, false, 24, passattempts, p)
                            if ok {
                                for _, a := range l.Doors {
                                    if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
                                        newlvl = a.NewLvl
                                        lvlchange = true
                                    }
                                }
                            } else {
                                if blocker == "npc" {
                                    passattempts++
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
                            ok, blocker := l.TryUpdatePos(true, p, true, 24, passattempts, p)
                            if ok {
                                for _, a := range l.Doors {
                                    if p.Pos[0] == a.GetCoords()[0] && p.Pos[1] == a.GetCoords()[1] {
                                        newlvl = a.NewLvl
                                        lvlchange = true
                                    }
                                }
                            } else {
                                if blocker == "npc" {
                                    passattempts++
                                }
                            }
                        }
                        count++
                    case 4:
                        stopped = true
                        count = 0
                        passattempts = 0
                    }
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
            r = text.BoundString(fo, fmt.Sprint("> aaaaaaaaaaaaaaaaaaaaaaaa -- Level: aaaaaaaaaaaa"))
            hei = r.Max.Y - r.Min.Y
            wid = r.Max.X - r.Min.X
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
            r = text.BoundString(fo, warning + selection)
            hei = r.Max.Y - r.Min.Y
            wid = r.Max.X - r.Min.X
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
            r = text.BoundString(fo, "aaaaaaaaaaaaaaaaaaaaaaaa")
            hei = r.Max.Y - r.Min.Y
            wid = r.Max.X - r.Min.X
            inputgm := ebiten.GeoM{}
            inputgm.Translate(float64((w / 2) - (wid / 2) - 8), float64((h / 2) - (hei / 2) - 16))
            inputimg := ebiten.NewImage(wid + 8, hei + 16)
            inputimg.Fill(color.Black)
            screen.DrawImage(
                inputimg, &ebiten.DrawImageOptions{
                    GeoM: inputgm})
            text.Draw(screen, sb.String(), fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + (3 * hei / 8), color.White)
            r2 = text.BoundString(fo, "Name")
            hei2 = r2.Max.Y - r2.Min.Y
            wid2 = r2.Max.X - r2.Min.X
            text.Draw(screen, "Name", fo, (w / 2) - (wid2 / 2), (h / 2) - (hei2 * 2), color.White)
        } else {
            r = text.BoundString(fo, "> New Game <")
            hei = r.Max.Y - r.Min.Y
            wid = r.Max.X - r.Min.X
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
                save = true
            }
        }
    } else if creation {
        text.Draw(screen, fmt.Sprintf("Name:       %s", name), fo, 32, 64, color.White)
        text.Draw(screen, "Class:", fo, 32, 128, color.White)
        if creationsel < 6 {
            text.Draw(screen, ">", fo, 112, (64 * (creationsel + 1)) + 64, color.White)
        } else {
            text.Draw(screen, ">", fo, 256, (64 * ((creationsel % 6) + 1)) + 64, color.White)
        }
        for classind, classval := range classessli {
            if classind < 6 {
                text.Draw(screen, classval, fo, 128, (64 * (classind + 1)) + 64, color.White)
            } else {
                text.Draw(screen, classval, fo, 272, (64 * ((classind % 6) + 1)) + 64, color.White)
            }
        }
    } else if l != nil {
        lgm := ebiten.GeoM{}
        lgm.Translate(float64((w / 2) + l.Pos[0]), float64((h / 2) + l.Pos[1]))
        screen.DrawImage(l.Image, &ebiten.DrawImageOptions{GeoM: lgm})
        for npcind, npc := range l.NPCs {
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
            if npcind == targeted {
                tbvgm := ebiten.GeoM{}
                tbvgm.Translate(float64((w / 2) + l.Pos[0] + npc.PC.Pos[0]), float64((h / 2) + l.Pos[1] + npc.PC.Pos[1]))
                screen.DrawImage(targetedBoxVert, &ebiten.DrawImageOptions{GeoM: tbvgm})
                tbvgm.Translate(float64(46), float64(0))
                screen.DrawImage(targetedBoxVert, &ebiten.DrawImageOptions{GeoM: tbvgm})
                tbhgm := ebiten.GeoM{}
                tbhgm.Translate(float64((w / 2) + l.Pos[0] + npc.PC.Pos[0]), float64((h / 2) + l.Pos[1] + npc.PC.Pos[1]))
                screen.DrawImage(targetedBoxHoriz, &ebiten.DrawImageOptions{GeoM: tbhgm})
                tbhgm.Translate(float64(0), float64(46))
                screen.DrawImage(targetedBoxHoriz, &ebiten.DrawImageOptions{GeoM: tbhgm})
                lineofsight, losvert, slope := l.LineOfSight(p, npc.PC.Pos)
                if lineofsight {
                    if losvert {
                        dist := p.Pos[1] - npc.PC.Pos[1]
                        if dist < 0 {
                            dist = -dist
                        }
                        losline := ebiten.NewImage(2, dist)
                        losline.Fill(color.RGBA{0x0, 0xff, 0x0, 0xff})
                        loslinegm := ebiten.GeoM{}
                        if p.Pos[1] > npc.PC.Pos[1] {
                            loslinegm.Translate(float64((w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24), float64((h / 2) + l.Pos[1] + npc.PC.Pos[1] + 24))
                        } else {
                            loslinegm.Translate(float64((w / 2) + l.Pos[0] + p.Pos[0] + 24), float64((h / 2) + l.Pos[1] + p.Pos[1] + 24))
                        }
                        screen.DrawImage(losline, &ebiten.DrawImageOptions{GeoM: loslinegm})
                    } else {
                        if p.Pos[0] > npc.PC.Pos[0] {
                            for linex := (w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24; linex <= (w / 2) + l.Pos[0] + p.Pos[0] + 24; linex++ {
                                liney := int((float64(linex - ((w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24)) * slope) + float64((h / 2) + l.Pos[1] + npc.PC.Pos[1] + 24))
                                screen.Set(linex, liney, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                if slope > 2.0 {
                                    for step := int(slope); step > 0; step-- {
                                        screen.Set(linex, liney + step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                        screen.Set(linex + 1, liney + step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                    }
                                } else if slope < -2.0 {
                                    for step := int(slope); step < 0; step++ {
                                        screen.Set(linex, liney - step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                        screen.Set(linex + 1, liney - step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                    }
                                } else {
                                    screen.Set(linex, liney + 1, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                    screen.Set(linex + 1, liney + 1, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                }
                            }
                        } else {
                            for linex := (w / 2) + l.Pos[0] + p.Pos[0] + 24; linex <= (w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24; linex++ {
                                liney := int((float64(linex - ((w / 2) + l.Pos[0] + p.Pos[0] + 24)) * slope) + float64((h / 2) + l.Pos[1] + p.Pos[1] + 24))
                                screen.Set(linex, liney, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                if slope > 2.0 {
                                    for step := int(slope); step > 0; step-- {
                                        screen.Set(linex, liney + step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                        screen.Set(linex + 1, liney + step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                    }
                                } else if slope < -2.0 {
                                    for step := int(slope); step < 0; step++ {
                                        screen.Set(linex, liney - step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                        screen.Set(linex + 1, liney - step, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                    }
                                } else {
                                    screen.Set(linex, liney + 1, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                    screen.Set(linex + 1, liney + 1, color.RGBA{0x0, 0xff, 0x0, 0xff})
                                }
                            }
                        }
                    }
                } else {
                    if losvert {
                        dist := p.Pos[1] - npc.PC.Pos[1]
                        if dist < 0 {
                            dist = -dist
                        }
                        losline := ebiten.NewImage(2, dist)
                        losline.Fill(color.RGBA{0xff, 0x0, 0x0, 0xff})
                        loslinegm := ebiten.GeoM{}
                        if p.Pos[1] > npc.PC.Pos[1] {
                            loslinegm.Translate(float64((w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24), float64((h / 2) + l.Pos[1] + npc.PC.Pos[1] + 24))
                        } else {
                            loslinegm.Translate(float64((w / 2) + l.Pos[0] + p.Pos[0] + 24), float64((h / 2) + l.Pos[1] + p.Pos[1] + 24))
                        }
                        screen.DrawImage(losline, &ebiten.DrawImageOptions{GeoM: loslinegm})
                    } else {
                        if p.Pos[0] > npc.PC.Pos[0] {
                            for linex := (w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24; linex <= (w / 2) + l.Pos[0] + p.Pos[0] + 24; linex++ {
                                liney := int((float64(linex - ((w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24)) * slope) + float64((h / 2) + l.Pos[1] + npc.PC.Pos[1] + 24))
                                screen.Set(linex, liney, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                if slope > 2.0 {
                                    for step := int(slope); step > 0; step-- {
                                        screen.Set(linex, liney + step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                        screen.Set(linex + 1, liney + step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                    }
                                } else if slope < -2.0 {
                                    for step := int(slope); step < 0; step++ {
                                        screen.Set(linex, liney - step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                        screen.Set(linex + 1, liney - step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                    }
                                } else {
                                    screen.Set(linex, liney + 1, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                    screen.Set(linex + 1, liney + 1, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                }
                            }
                        } else {
                            for linex := (w / 2) + l.Pos[0] + p.Pos[0] + 24; linex <= (w / 2) + l.Pos[0] + npc.PC.Pos[0] + 24; linex++ {
                                liney := int((float64(linex - ((w / 2) + l.Pos[0] + p.Pos[0] + 24)) * slope) + float64((h / 2) + l.Pos[1] + p.Pos[1] + 24))
                                screen.Set(linex, liney, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                if slope > 2.0 {
                                    for step := int(slope); step > 0; step-- {
                                        screen.Set(linex, liney + step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                        screen.Set(linex + 1, liney + step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                    }
                                } else if slope < -2.0 {
                                    for step := int(slope); step < 0; step++ {
                                        screen.Set(linex, liney - step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                        screen.Set(linex + 1, liney - step, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                    }
                                } else {
                                    screen.Set(linex, liney + 1, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                    screen.Set(linex + 1, liney + 1, color.RGBA{0xff, 0x0, 0x0, 0xff})
                                }
                            }
                        }
                    }
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
        r = text.BoundString(fo, dialogstrs[0])
        hei = r.Max.Y - r.Min.Y
        if s < len(dialogstrs) {
            text.Draw(screen, npcname, fo, 140, 500, color.RGBA{200, 36, 121, 255})
            text.Draw(screen, dialogstrs[s], fo, 140, 516 + hei, color.Black)
            if s + 1 < len(dialogstrs) {
                text.Draw(screen, dialogstrs[s + 1], fo, 140, 524 + (hei * 2), color.Black)
                if s + 2 < len(dialogstrs) {
                    dagm = ebiten.GeoM{}
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
    if invmenu {
        screen.DrawImage(blankImage, nil)
        invitems := p.Inv.GetItems()
        for iind, ival := range invitems {
            if iind == invsel {
                text.Draw(screen, ">", fo, 32, 64 + (32 * iind), color.White)
            }
            text.Draw(screen, ival.PrettyPrint(), fo, 64, 64 + (32 * iind), color.White)
        }
    }
    if invselmenu {
        r = text.BoundString(fo, "> Use (illuminate)")
        hei = r.Max.Y - r.Min.Y
        wid = r.Max.X - r.Min.X
        ismgm = ebiten.GeoM{}
        ismgm.Translate(float64((w / 2) - (wid / 2) - 8), float64((h / 2) - (3 * hei / 2) - 16))
        ismimg = ebiten.NewImage(wid + 28, (hei * 3) + 48)
        ismimg.Fill(color.Black)
        screen.DrawImage(
            ismimg, &ebiten.DrawImageOptions{
                GeoM: ismgm})
        ismgm2 = ebiten.GeoM{}
        ismgm2.Translate(float64((w / 2) - (wid / 2) - 4), float64((h / 2) - (3 * hei / 2) - 12))
        ismimg2 = ebiten.NewImage(wid + 20, (hei * 3) + 40)
        ismimg2.Fill(color.White)
        screen.DrawImage(
            ismimg2, &ebiten.DrawImageOptions{
                GeoM: ismgm2})
        switch invsel2 {
        case 0:
            text.Draw(screen, "> Equip", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            if actcheck = p.Inv.GetItems()[invsel].Action(); actcheck == "" {
                text.Draw(screen, "  Use", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Gray16{0x8000})
            } else {
                text.Draw(screen, fmt.Sprintf("  Use (%s)", actcheck), fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            }
            text.Draw(screen, "  Drop", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
        case 1:
            if p.Inv.GetItems()[invsel].Slot() == "" {
                text.Draw(screen, "  Equip", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Gray16{0x8000})
            } else {
                text.Draw(screen, "  Equip", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            }
            if actcheck = p.Inv.GetItems()[invsel].Action(); actcheck == "" {
                text.Draw(screen, "> Use", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            } else {
                text.Draw(screen, fmt.Sprintf("> Use (%s)", actcheck), fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            }
            text.Draw(screen, "  Drop", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
        case 2:
            if p.Inv.GetItems()[invsel].Slot() == "" {
                text.Draw(screen, "  Equip", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Gray16{0x8000})
            } else {
                text.Draw(screen, "  Equip", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) + 16, color.Black)
            }
            if actcheck = p.Inv.GetItems()[invsel].Action(); actcheck == "" {
                text.Draw(screen, "  Use", fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Gray16{0x8000})
            } else {
                text.Draw(screen, fmt.Sprintf("  Use (%s)", actcheck), fo, (w / 2) - (wid / 2), (h / 2) - (hei / 2) + 24, color.Black)
            }
            text.Draw(screen, "> Drop", fo, (w / 2) - (wid / 2), (h / 2) + (hei / 2) + 32, color.Black)
        }
    }
    if effectmsg {
        if countend == 0 {
            countend = (npcCount + 300) % 6000
        }
        switch effectact {
        case "illuminate":
            r = text.BoundString(fo, "Your path is illuminated:")
            hei = r.Max.Y - r.Min.Y
            wid = r.Max.X - r.Min.X
            illumed, illumstats := p.Class.Illuminated()
            if illumed {
                dur, err = time.ParseDuration(strconv.Itoa(illumstats[2] * 6) + "s")
                if err != nil {
                    log.Fatal(err)
                }
                r2 = text.BoundString(fo, fmt.Sprintf("%d feet bright light, then %d feet dim light", illumstats[0], illumstats[1]))
                hei2 = r2.Max.Y - r.Min.Y
                wid2 = r2.Max.X - r.Min.X
                r3 = text.BoundString(fo, fmt.Sprintf("The effect will last for the next %d turns (%v)", illumstats[2], dur))
                hei3 = r3.Max.Y - r.Min.Y
                wid3 = r3.Max.X - r.Min.X
                effgm = ebiten.GeoM{}
                effgm.Translate(float64((w / 2) - (wid3 / 2) - 8), float64((h / 2) - (3 * hei3) - 20))
                effimg = ebiten.NewImage(wid3 + 28, (hei3 * 3) + 48)
                effimg.Fill(color.Black)
                screen.DrawImage(
                    effimg, &ebiten.DrawImageOptions{
                        GeoM: effgm})
                effgm2 = ebiten.GeoM{}
                effgm2.Translate(float64((w / 2) - (wid3 / 2) - 4), float64((h / 2) - (3 * hei3) - 16))
                effimg2 = ebiten.NewImage(wid3 + 20, (hei3 * 3) + 40)
                effimg2.Fill(color.White)
                screen.DrawImage(
                    effimg2, &ebiten.DrawImageOptions{
                        GeoM: effgm2})
                text.Draw(screen, "Your path is illuminated:", fo, (w / 2) - (wid / 2), (h / 2) - (3 * hei / 2) - 16, color.RGBA{159, 11, 19, 255})
                text.Draw(screen, fmt.Sprintf("%d feet bright light, then %d feet dim light", illumstats[0], illumstats[1]), fo, (w / 2) - (wid2 / 2), (h / 2) - (hei2 / 2) - 8, color.RGBA{159, 11, 19, 255})
                text.Draw(screen, fmt.Sprintf("The effect will last for the next %d turns (%v)", illumstats[2], dur), fo, (w / 2) - (wid3 / 2), (h / 2) + (hei3 / 2), color.RGBA{159, 11, 19, 255})
            }
        case "disguise":
            log.Println("Need to implement disguise menu")
        case "write":
//            p.WriteMsg = `This is a test file written to work out the kinks with reading written pages in the game.
//I will be trying to get these kinks worked out over the course of the next few days.
//
//Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam ligula dolor, condimentum rhoncus eros ac,
//tincidunt faucibus felis. Mauris efficitur sagittis ipsum, malesuada feugiat nisl tincidunt id.
//Cras vitae purus facilisis, venenatis erat ac, sollicitudin tortor. Morbi euismod consequat eros in tincidunt.
//Nunc rhoncus odio vel lectus hendrerit eleifend. Sed vestibulum neque non mattis varius. Nullam eget nibh elementum,
//luctus odio sed, ullamcorper felis. Fusce feugiat pellentesque ligula eu placerat.
//Nulla suscipit lacus eget tellus condimentum, ac iaculis neque condimentum. Praesent vitae lectus neque.
//
//Nulla arcu leo, interdum nec aliquet vel, dictum ac elit. Phasellus tempus massa in eleifend venenatis.
//Mauris accumsan leo eget egestas ornare. Maecenas varius iaculis nibh, ac volutpat mauris tempus vel.
//Morbi et bibendum nisl, vel dignissim neque. Suspendisse nec metus faucibus, tincidunt neque quis, pulvinar justo.
//Integer tortor ante, euismod faucibus dictum a, iaculis vel ex. Duis pellentesque in mauris in tempus.
//Fusce commodo iaculis vehicula. Aenean ornare ante a magna euismod accumsan. In at justo ac quam bibendum commodo. Sed eu mollis nisi.
//
//Nunc venenatis efficitur lacus, sed malesuada lectus. Pellentesque blandit enim urna, non laoreet mauris accumsan quis.
//Quisque a metus tellus. Pellentesque condimentum velit et bibendum molestie. Pellentesque vehicula cursus erat,
//vel convallis sem mattis quis. Quisque at consectetur sem. Nulla magna leo, vulputate et vulputate vitae, malesuada in eros.
//Curabitur vel iaculis mi.
//
//Quisque dictum nisl vel ligula condimentum, sit amet ultricies massa dictum. Curabitur nec lacus ac odio dapibus fringilla.
//Vivamus non aliquet quam. Nunc condimentum ipsum in nisl hendrerit mattis aliquam a orci. Etiam eleifend sagittis enim a mollis.
//Nullam volutpat ac risus in fermentum.`
            readimg.Fill(color.White)
            readimg2.Fill(color.Black)
            screen.DrawImage(
                readimg, &ebiten.DrawImageOptions{
                    GeoM: readgm})
            screen.DrawImage(
                readimg2, &ebiten.DrawImageOptions{
                    GeoM: readgm2})
            var y int
            result := ""
            lines := strings.Split(sb.String(), "\n")
            for ind, line := range lines {
                if len(line) > 55 {
                    for x := 54; x < len(line); x = y + 56 {
                        for y = x; line[y] != ' '; y-- {
                            if y == 0 {
                                y = 54
                                break
                            }
                            continue
                        }
                        line = line[:y + 1] + "\n" + line[y + 1:]
                    }
                }
                if ind + 1 < len(lines) {
                    result += line + "\n"
                } else {
                    result += line
                }
            }
            if len(result) > 0 {
                sb.Reset()
                _, err = sb.WriteString(result)
            }
            resslice := strings.Split(result, "\n")
            if len(resslice) > 18 {
                offset := len(resslice) - 18
                for z := offset; z < len(resslice); z++ {
                    text.Draw(screen, resslice[z], fo, (768 / 2) - (724 / 2) + 28, (576 / 2) - (552 / 2) + 48 + (28 * ((z - offset) % 19)), color.White)
                }
            } else {
                for z := 0; z < len(resslice); z++ {
                    text.Draw(screen, resslice[z], fo, (768 / 2) - (724 / 2) + 28, (576 / 2) - (552 / 2) + 48 + (28 * z), color.White)
                }
            }
        case "read":
            if len(p.PageMsgs) == len(pages) {
                readimg.Fill(color.Black)
                readimg2.Fill(color.White)
                if len(pages) > 0 {
                    screen.DrawImage(
                        readimg, &ebiten.DrawImageOptions{
                            GeoM: readgm})
                    screen.DrawImage(
                        readimg2, &ebiten.DrawImageOptions{
                            GeoM: readgm2})
                    moreshown = false
                    for y := ((overflowcur - 1) * p.PageMsgs[pageind][2].(int)) + p.PageMsgs[pageind][2].(int); y < p.PageMsgs[pageind][1].(int); y++ {
                        if y < (overflowcur * p.PageMsgs[pageind][2].(int) - 1) + p.PageMsgs[pageind][2].(int) {
                            text.Draw(screen, p.PageMsgs[pageind][0].([]string)[y], fo, (768 / 2) - (724 / 2) + 28, (576 / 2) - (552 / 2) + 48 + (28 * (y % p.PageMsgs[pageind][2].(int))), color.Black)
                        } else {
                            if !moreshown {
                                text.Draw(screen, "V", fo, (768 / 2) + (724 / 2) - 48, (576 / 2) + (552 / 2) - 48, color.Black)
                                moreshown = true
                            }
                        }
                    }
                } else {
                    r = text.BoundString(fo, fmt.Sprint("You do not have any written pages"))
                    wid = r.Max.X - r.Min.X
                    screen.DrawImage(
                        readimg3, &ebiten.DrawImageOptions{
                            GeoM: readgm3})
                    text.Draw(screen, "You do not have any written pages", fo, (768 / 2) - (wid / 2), (576 / 2), color.White)
                    if npcCount >= countend {
                        countend = 0
                        effectmsg = false
                        effectact = ""
                    }
                }
            } else {
                log.Println(fmt.Sprintf("len(p.PageMsgs) == %d but len(pages) == %d", len(p.PageMsgs), len(pages)))
                effectact = ""
                effectmsg = false
                overflowcur = 0
                overflownum = 0
                pageind = 0
                return
            }
        case "throw":
            // DRAW throwtarget box
            if throwtarget == [2]int{} {
                return
            } else {
                ttbvgm := ebiten.GeoM{}
                ttbvgm.Translate(float64((w / 2) + l.Pos[0] + throwtarget[0] + 12), float64((h / 2) + l.Pos[1] + throwtarget[1] + 12))
                screen.DrawImage(throwTargetBoxVert, &ebiten.DrawImageOptions{GeoM: ttbvgm})
                ttbvgm.Translate(float64(22), float64(0))
                screen.DrawImage(throwTargetBoxVert, &ebiten.DrawImageOptions{GeoM: ttbvgm})
                ttbhgm := ebiten.GeoM{}
                ttbhgm.Translate(float64((w / 2) + l.Pos[0] + throwtarget[0] + 12), float64((h / 2) + l.Pos[1] + throwtarget[1] + 12))
                screen.DrawImage(throwTargetBoxHoriz, &ebiten.DrawImageOptions{GeoM: ttbhgm})
                ttbhgm.Translate(float64(0), float64(22))
                screen.DrawImage(throwTargetBoxHoriz, &ebiten.DrawImageOptions{GeoM: ttbhgm})
            }
        case "playmusic":
            return
        default:
            log.Fatal(effectact + " is not defined")
        }
        if effectact == "illuminate" && npcCount >= countend {
            countend = 0
            effectmsg = false
            effectact = ""
        }
    }
    var strstats [2]int
    var dexstats [2]int
    var constats [2]int
    var intelstats [2]int
    var wisstats [2]int
    var chastats [2]int
    if p != nil {
        if p.Class != nil {
            strstats = p.Class.GetStr()
            dexstats = p.Class.GetDex()
            constats = p.Class.GetCon()
            intelstats = p.Class.GetIntel()
            wisstats = p.Class.GetWis()
            chastats = p.Class.GetCha()
        }
    }
    if charsheet0 {
        screen.DrawImage(blankImage, nil)
        text.Draw(screen, fmt.Sprintf("Name: %s", p.Name), fo, 32, 32, color.White)
        text.Draw(screen, fmt.Sprintf("Class: %s", p.Class.GetName()), fo, 32, 64, color.White)
        text.Draw(screen, fmt.Sprintf("Level: %d", p.Class.GetLevel()), fo, 256, 64, color.White)
        text.Draw(screen, fmt.Sprintf("Str: %d (%+d)", strstats[0], strstats[1]), fo, 32, 96, color.White)
        text.Draw(screen, fmt.Sprintf("Dex: %d (%+d)", dexstats[0], dexstats[1]), fo, 32, 128, color.White)
        text.Draw(screen, fmt.Sprintf("Con: %d (%+d)", constats[0], constats[1]), fo, 32, 160, color.White)
        text.Draw(screen, fmt.Sprintf("Int: %d (%+d)", intelstats[0], intelstats[1]), fo, 32, 192, color.White)
        text.Draw(screen, fmt.Sprintf("Wis: %d (%+d)", wisstats[0], wisstats[1]), fo, 32, 224, color.White)
        text.Draw(screen, fmt.Sprintf("Cha: %d (%+d)", chastats[0], chastats[1]), fo, 32, 256, color.White)
        text.Draw(screen, fmt.Sprintf("Proficiency Bonus: %+d", p.Class.GetPB()), fo, 32, 288, color.White)
        text.Draw(screen, ">", fo, 704, 512, color.White)
    } else if charsheet1 {
        screen.DrawImage(blankImage, nil)
        text.Draw(screen, "<", fo, 64, 512, color.White)
        text.Draw(screen, fmt.Sprintf("Name: %s", p.Name), fo, 32, 32, color.White)
        text.Draw(screen, "Saving Throws:", fo, 32, 64, color.White)
        text.Draw(screen, fmt.Sprintf("Str: %+d", p.Class.SavingThrow("str")), fo, 32, 96, color.White)
        text.Draw(screen, fmt.Sprintf("Dex: %+d", p.Class.SavingThrow("dex")), fo, 32, 128, color.White)
        text.Draw(screen, fmt.Sprintf("Con: %+d", p.Class.SavingThrow("con")), fo, 32, 160, color.White)
        text.Draw(screen, fmt.Sprintf("Int: %+d", p.Class.SavingThrow("intel")), fo, 32, 192, color.White)
        text.Draw(screen, fmt.Sprintf("Wis: %+d", p.Class.SavingThrow("wis")), fo, 32, 224, color.White)
        text.Draw(screen, fmt.Sprintf("Cha: %+d", p.Class.SavingThrow("cha")), fo, 32, 256, color.White)
        text.Draw(screen, "Skills:", fo, 256, 64, color.White)
        text.Draw(screen, fmt.Sprintf("Acrobatics:      %+d", p.Class.SkillCheck("acrobatics")), fo, 256, 96, color.White)
        text.Draw(screen, fmt.Sprintf("Animal Handling: %+d", p.Class.SkillCheck("animalhandling")), fo, 256, 128, color.White)
        text.Draw(screen, fmt.Sprintf("Arcana:          %+d", p.Class.SkillCheck("arcana")), fo, 256, 160, color.White)
        text.Draw(screen, fmt.Sprintf("Athletics:       %+d", p.Class.SkillCheck("athletics")), fo, 256, 192, color.White)
        text.Draw(screen, fmt.Sprintf("Deception:       %+d", p.Class.SkillCheck("deception")), fo, 256, 224, color.White)
        text.Draw(screen, fmt.Sprintf("History:         %+d", p.Class.SkillCheck("history")), fo, 256, 256, color.White)
        text.Draw(screen, fmt.Sprintf("Insight:         %+d", p.Class.SkillCheck("insight")), fo, 256, 288, color.White)
        text.Draw(screen, fmt.Sprintf("Intimidation:    %+d", p.Class.SkillCheck("intimidation")), fo, 256, 320, color.White)
        text.Draw(screen, fmt.Sprintf("Investigation:   %+d", p.Class.SkillCheck("investigation")), fo, 256, 352, color.White)
        text.Draw(screen, fmt.Sprintf("Medicine:        %+d", p.Class.SkillCheck("medicine")), fo, 512, 96, color.White)
        text.Draw(screen, fmt.Sprintf("Nature:          %+d", p.Class.SkillCheck("nature")), fo, 512, 128, color.White)
        text.Draw(screen, fmt.Sprintf("Perception:      %+d", p.Class.SkillCheck("perception")), fo, 512, 160, color.White)
        text.Draw(screen, fmt.Sprintf("Performance:     %+d", p.Class.SkillCheck("performance")), fo, 512, 192, color.White)
        text.Draw(screen, fmt.Sprintf("Persuasion:      %+d", p.Class.SkillCheck("persuasion")), fo, 512, 224, color.White)
        text.Draw(screen, fmt.Sprintf("Religion:        %+d", p.Class.SkillCheck("religion")), fo, 512, 256, color.White)
        text.Draw(screen, fmt.Sprintf("Sleight of Hand: %+d", p.Class.SkillCheck("sleightofhand")), fo, 512, 288, color.White)
        text.Draw(screen, fmt.Sprintf("Stealth:         %+d", p.Class.SkillCheck("stealth")), fo, 512, 320, color.White)
        text.Draw(screen, fmt.Sprintf("Survival:        %+d", p.Class.SkillCheck("survival")), fo, 512, 352, color.White)
        text.Draw(screen, ">", fo, 704, 512, color.White)
    } else if charsheet2 {
        screen.DrawImage(blankImage, nil)
        text.Draw(screen, "<", fo, 64, 512, color.White)
        text.Draw(screen, fmt.Sprintf("Name: %s", p.Name), fo, 32, 32, color.White)
        text.Draw(screen, "Equipment:", fo, 32, 64, color.White)
        if p.Equipment.Head != nil {
            text.Draw(screen, fmt.Sprintf("Head: %s", p.Equipment.Head.PrettyPrint()), fo, 64, 96, color.White)
        } else {
            text.Draw(screen, "Head:", fo, 64, 96, color.Gray16{0x8000})
        }
        if p.Equipment.Torso != nil {
            text.Draw(screen, fmt.Sprintf("Torso: %s", p.Equipment.Torso.PrettyPrint()), fo, 64, 128, color.White)
        } else {
            text.Draw(screen, "Torso:", fo, 64, 128, color.Gray16{0x8000})
        }
        if p.Equipment.Legs != nil {
            text.Draw(screen, fmt.Sprintf("Legs: %s", p.Equipment.Legs.PrettyPrint()), fo, 64, 160, color.White)
        } else {
            text.Draw(screen, "Legs:", fo, 64, 160, color.Gray16{0x8000})
        }
        if p.Equipment.Feet != nil {
            text.Draw(screen, fmt.Sprintf("Feet: %s", p.Equipment.Feet.PrettyPrint()), fo, 64, 192, color.White)
        } else {
            text.Draw(screen, "Feet:", fo, 64, 192, color.Gray16{0x8000})
        }
        if p.Equipment.LeftHand != nil {
            text.Draw(screen, fmt.Sprintf("Left Hand: %s", p.Equipment.LeftHand.PrettyPrint()), fo, 384, 96, color.White)
        } else {
            text.Draw(screen, "Left Hand:", fo, 384, 96, color.Gray16{0x8000})
        }
        if p.Equipment.RightHand != nil {
            text.Draw(screen, fmt.Sprintf("Right Hand: %s", p.Equipment.RightHand.PrettyPrint()), fo, 384, 128, color.White)
        } else {
            text.Draw(screen, "Right Hand:", fo, 384, 128, color.Gray16{0x8000})
        }
        if p.Equipment.BothHands != nil {
            text.Draw(screen, fmt.Sprintf("Both Hands: %s", p.Equipment.BothHands.PrettyPrint()), fo, 384, 160, color.White)
        } else {
            text.Draw(screen, "Both Hands:", fo, 384, 160, color.Gray16{0x8000})
        }
    }
    if overworld {
        screen.DrawImage(blankImage, nil)
        owgm = ebiten.GeoM{}
        owgm.Translate(float64(iw) - (float64(iw) / 2.0), 0.0)
        screen.DrawImage(
            overworldImage, &ebiten.DrawImageOptions{
                GeoM: owgm})
    }
    if pause {
        r = text.BoundString(fo, "> Save game")
        hei = r.Max.Y - r.Min.Y
        wid = r.Max.X - r.Min.X
        pausegm = ebiten.GeoM{}
        pausegm.Translate(float64((w / 2) - (wid / 2) - 8), float64((h / 2) - (3 * hei / 2) - 16))
        pauseimg = ebiten.NewImage(wid + 28, (hei * 5) + 64)
        pauseimg.Fill(color.Black)
        screen.DrawImage(
            pauseimg, &ebiten.DrawImageOptions{
                GeoM: pausegm})
        pausegm2 = ebiten.GeoM{}
        pausegm2.Translate(float64((w / 2) - (wid / 2) - 4), float64((h / 2) - (3 * hei / 2) - 12))
        pauseimg2 = ebiten.NewImage(wid + 20, (hei * 5) + 56)
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
        if save {
            return
        }
        fadegm := ebiten.GeoM{}
        fadegm.Translate(float64((w / 2) + l.Pos[0]), float64((h / 2) + l.Pos[1]))
        op := &ebiten.DrawImageOptions{GeoM: fadegm}
        if npcCount % 5 == 0 {
            f++
        }
        if f == 0 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.0)
            screen.DrawImage(fadeImage, op)
        } else if f == 1 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.1)
            screen.DrawImage(fadeImage, op)
        } else if f == 2 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.2)
            screen.DrawImage(fadeImage, op)
        } else if f == 3 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.3)
            screen.DrawImage(fadeImage, op)
        } else if f == 4 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.4)
            screen.DrawImage(fadeImage, op)
        } else if f == 5 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.5)
            screen.DrawImage(fadeImage, op)
        } else if f == 6 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.6)
            screen.DrawImage(fadeImage, op)
        } else if f == 7 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.7)
            screen.DrawImage(fadeImage, op)
        } else if f == 8 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.8)
            screen.DrawImage(fadeImage, op)
        } else if f == 9 {
            op.ColorM.Scale(1.0, 1.0, 1.0, 0.9)
            screen.DrawImage(fadeImage, op)
        } else if f == 10 {
            screen.DrawImage(fadeImage, nil)
            f = 0
            lvlchange = false
            for _, lvl := range levelslice {
                if lvl.GetName() == newlvl[0] {
                    l = lvl
                    lvlloaded = true
                }
            }
            if !lvlloaded {
                l = levels.LoadLvl(newlvl...)
                levelslice = append(levelslice, l)
            }
            targeted = -1
            p.Pos[0] = -l.Pos[0]
            p.Pos[1] = -l.Pos[1]
            if l.Cutscene > 0 {
                curCS = l.Cutscene
                cutscene = true
            }
            lvlloaded = false
        }
    }
    if gainProfST {
        screen.DrawImage(blankImage, nil)
        text.Draw(screen, fmt.Sprintf("You're now level %d!\nGain proficiency in a saving throw:", p.Class.GetLevel()), fo, 32, 64, color.White)
        switch gainprofstsel {
        case 0:
            text.Draw(screen, ">", fo, 32, 96, color.White)
        case 1:
            text.Draw(screen, ">", fo, 32, 128, color.White)
        case 2:
            text.Draw(screen, ">", fo, 32, 160, color.White)
        case 3:
            text.Draw(screen, ">", fo, 32, 192, color.White)
        case 4:
            text.Draw(screen, ">", fo, 32, 224, color.White)
        case 5:
            text.Draw(screen, ">", fo, 32, 256, color.White)
        default:
            log.Fatal(fmt.Sprintf("%d is not a valid gainprofstsel", gainprofstsel))
        }
        strmod := p.Class.GetStr()[1]
        if p.Class.SavingThrow("str") > strmod {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("str")), fo, 40, 96, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("str")), fo, 40, 96, color.White)
        }
        dexmod := p.Class.GetDex()[1]
        if p.Class.SavingThrow("dex") > dexmod {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("dex")), fo, 40, 128, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("dex")), fo, 40, 128, color.White)
        }
        conmod := p.Class.GetCon()[1]
        if p.Class.SavingThrow("con") > conmod {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("con")), fo, 40, 160, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("con")), fo, 40, 160, color.White)
        }
        intelmod := p.Class.GetIntel()[1]
        if p.Class.SavingThrow("intel") > intelmod {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("intel")), fo, 40, 192, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("intel")), fo, 40, 192, color.White)
        }
        wismod := p.Class.GetWis()[1]
        if p.Class.SavingThrow("wis") > wismod {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("wis")), fo, 40, 224, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("wis")), fo, 40, 224, color.White)
        }
        chamod := p.Class.GetCha()[1]
        if p.Class.SavingThrow("cha") > chamod {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("cha")), fo, 40, 256, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Strength: %+d", p.Class.SavingThrow("cha")), fo, 40, 256, color.White)
        }
        if stprofexists != "" {
            text.Draw(screen, fmt.Sprintf("You are already proficient in %s saving throws", stprofexists), fo, 64, 16, color.RGBA{159, 11, 19, 255})
        }
    }
    if gainProfSkill {
        screen.DrawImage(blankImage, nil)
        text.Draw(screen, fmt.Sprintf("You're now level %d!\nGain proficiency in a skill:", p.Class.GetLevel()), fo, 32, 64, color.White)
        switch gainprofskillsel {
        case 0:
            text.Draw(screen, ">", fo, 32, 96, color.White)
        case 1:
            text.Draw(screen, ">", fo, 32, 128, color.White)
        case 2:
            text.Draw(screen, ">", fo, 32, 160, color.White)
        case 3:
            text.Draw(screen, ">", fo, 32, 192, color.White)
        case 4:
            text.Draw(screen, ">", fo, 32, 224, color.White)
        case 5:
            text.Draw(screen, ">", fo, 32, 256, color.White)
        case 6:
            text.Draw(screen, ">", fo, 32, 288, color.White)
        case 7:
            text.Draw(screen, ">", fo, 32, 320, color.White)
        case 8:
            text.Draw(screen, ">", fo, 32, 352, color.White)
        case 9:
            text.Draw(screen, ">", fo, 256, 96, color.White)
        case 10:
            text.Draw(screen, ">", fo, 256, 128, color.White)
        case 11:
            text.Draw(screen, ">", fo, 256, 160, color.White)
        case 12:
            text.Draw(screen, ">", fo, 256, 192, color.White)
        case 13:
            text.Draw(screen, ">", fo, 256, 224, color.White)
        case 14:
            text.Draw(screen, ">", fo, 256, 256, color.White)
        case 15:
            text.Draw(screen, ">", fo, 256, 288, color.White)
        case 16:
            text.Draw(screen, ">", fo, 256, 320, color.White)
        case 17:
            text.Draw(screen, ">", fo, 256, 352, color.White)
        default:
            log.Fatal(fmt.Sprintf("%d is not a valid gainprofskillsel", gainprofstsel))
        }
        strmod := p.Class.GetStr()[1]
        dexmod := p.Class.GetDex()[1]
        //conmod := p.Class.GetCon()[1]
        intelmod := p.Class.GetIntel()[1]
        wismod := p.Class.GetWis()[1]
        chamod := p.Class.GetCha()[1]
        if p.Class.SkillCheck("acrobatics") > dexmod {
            text.Draw(screen, fmt.Sprintf("Acrobatics: %+d", p.Class.SkillCheck("acrobatics")), fo, 40, 96, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Acrobatics: %+d", p.Class.SkillCheck("acrobatics")), fo, 40, 96, color.White)
        }
        if p.Class.SkillCheck("animalhandling") > wismod {
            text.Draw(screen, fmt.Sprintf("Animal Handling: %+d", p.Class.SkillCheck("animalhandling")), fo, 40, 128, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Animal Handling: %+d", p.Class.SkillCheck("animalhandling")), fo, 40, 128, color.White)
        }
        if p.Class.SkillCheck("arcana") > intelmod {
            text.Draw(screen, fmt.Sprintf("Arcana: %+d", p.Class.SkillCheck("arcana")), fo, 40, 160, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Arcana: %+d", p.Class.SkillCheck("arcana")), fo, 40, 160, color.White)
        }
        if p.Class.SkillCheck("athletics") > strmod {
            text.Draw(screen, fmt.Sprintf("Athletics: %+d", p.Class.SkillCheck("athletics")), fo, 40, 192, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Athletics: %+d", p.Class.SkillCheck("athletics")), fo, 40, 192, color.White)
        }
        if p.Class.SkillCheck("deception") > chamod {
            text.Draw(screen, fmt.Sprintf("Deception: %+d", p.Class.SkillCheck("deception")), fo, 40, 224, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Deception: %+d", p.Class.SkillCheck("deception")), fo, 40, 224, color.White)
        }
        if p.Class.SkillCheck("history") > intelmod {
            text.Draw(screen, fmt.Sprintf("History: %+d", p.Class.SkillCheck("history")), fo, 40, 256, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("History: %+d", p.Class.SkillCheck("history")), fo, 40, 256, color.White)
        }
        if p.Class.SkillCheck("insight") > wismod {
            text.Draw(screen, fmt.Sprintf("Insight: %+d", p.Class.SkillCheck("insight")), fo, 40, 288, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Insight: %+d", p.Class.SkillCheck("insight")), fo, 40, 288, color.White)
        }
        if p.Class.SkillCheck("intimidation") > chamod {
            text.Draw(screen, fmt.Sprintf("Intimidation: %+d", p.Class.SkillCheck("intimidation")), fo, 40, 320, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Intimidation: %+d", p.Class.SkillCheck("intimidation")), fo, 40, 320, color.White)
        }
        if p.Class.SkillCheck("investigation") > intelmod {
            text.Draw(screen, fmt.Sprintf("Investigation: %+d", p.Class.SkillCheck("investigation")), fo, 40, 352, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Investigation: %+d", p.Class.SkillCheck("investigation")), fo, 40, 352, color.White)
        }
        if p.Class.SkillCheck("medicine") > wismod {
            text.Draw(screen, fmt.Sprintf("Medicine: %+d", p.Class.SkillCheck("medicine")), fo, 264, 96, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Medicine: %+d", p.Class.SkillCheck("medicine")), fo, 264, 96, color.White)
        }
        if p.Class.SkillCheck("nature") > intelmod {
            text.Draw(screen, fmt.Sprintf("Nature: %+d", p.Class.SkillCheck("nature")), fo, 264, 128, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Nature: %+d", p.Class.SkillCheck("nature")), fo, 264, 128, color.White)
        }
        if p.Class.SkillCheck("perception") > wismod {
            text.Draw(screen, fmt.Sprintf("Perception: %+d", p.Class.SkillCheck("perception")), fo, 264, 160, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Perception: %+d", p.Class.SkillCheck("perception")), fo, 264, 160, color.White)
        }
        if p.Class.SkillCheck("performance") > chamod {
            text.Draw(screen, fmt.Sprintf("Performance: %+d", p.Class.SkillCheck("performance")), fo, 264, 192, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Performance: %+d", p.Class.SkillCheck("performance")), fo, 264, 192, color.White)
        }
        if p.Class.SkillCheck("persuasion") > chamod {
            text.Draw(screen, fmt.Sprintf("Persuasion: %+d", p.Class.SkillCheck("persuasion")), fo, 264, 224, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Persuasion: %+d", p.Class.SkillCheck("persuasion")), fo, 264, 224, color.White)
        }
        if p.Class.SkillCheck("religion") > intelmod {
            text.Draw(screen, fmt.Sprintf("Religion: %+d", p.Class.SkillCheck("religion")), fo, 264, 256, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Religion: %+d", p.Class.SkillCheck("religion")), fo, 264, 256, color.White)
        }
        if p.Class.SkillCheck("sleightofhand") > dexmod {
            text.Draw(screen, fmt.Sprintf("Sleight of Hand: %+d", p.Class.SkillCheck("sleightofhand")), fo, 264, 288, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Sleight of Hand: %+d", p.Class.SkillCheck("sleightofhand")), fo, 264, 288, color.White)
        }
        if p.Class.SkillCheck("stealth") > dexmod {
            text.Draw(screen, fmt.Sprintf("Stealth: %+d", p.Class.SkillCheck("stealth")), fo, 264, 320, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Stealth: %+d", p.Class.SkillCheck("stealth")), fo, 264, 320, color.White)
        }
        if p.Class.SkillCheck("survival") > wismod {
            text.Draw(screen, fmt.Sprintf("Survival: %+d", p.Class.SkillCheck("survival")), fo, 264, 352, color.Gray16{0x8000})
        } else {
            text.Draw(screen, fmt.Sprintf("Survival: %+d", p.Class.SkillCheck("survival")), fo, 264, 352, color.White)
        }
        if skillprofexists != "" {
            text.Draw(screen, fmt.Sprintf("You are already proficient in the skill %s", skillprofexists), fo, 64, 16, color.RGBA{159, 11, 19, 255})
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

func Input(sb *strings.Builder) {
    switch {
    case inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.KeyPressDuration(ebiten.KeyA) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('A')
        } else {
            err = sb.WriteByte('a')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyB) || inpututil.KeyPressDuration(ebiten.KeyB) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('B')
        } else {
            err = sb.WriteByte('b')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyC) || inpututil.KeyPressDuration(ebiten.KeyC) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('C')
        } else {
            err = sb.WriteByte('c')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.KeyPressDuration(ebiten.KeyD) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('D')
        } else {
            err = sb.WriteByte('d')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyE) || inpututil.KeyPressDuration(ebiten.KeyE) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('E')
        } else {
            err = sb.WriteByte('e')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.KeyPressDuration(ebiten.KeyF) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('F')
        } else {
            err = sb.WriteByte('f')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyG) || inpututil.KeyPressDuration(ebiten.KeyG) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('G')
        } else {
            err = sb.WriteByte('g')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyH) || inpututil.KeyPressDuration(ebiten.KeyH) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('H')
        } else {
            err = sb.WriteByte('h')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyI) || inpututil.KeyPressDuration(ebiten.KeyI) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('I')
        } else {
            err = sb.WriteByte('i')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyJ) || inpututil.KeyPressDuration(ebiten.KeyJ) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('J')
        } else {
            err = sb.WriteByte('j')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyK) || inpututil.KeyPressDuration(ebiten.KeyK) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('K')
        } else {
            err = sb.WriteByte('k')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyL) || inpututil.KeyPressDuration(ebiten.KeyL) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('L')
        } else {
            err = sb.WriteByte('l')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyM) || inpututil.KeyPressDuration(ebiten.KeyM) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('M')
        } else {
            err = sb.WriteByte('m')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyN) || inpututil.KeyPressDuration(ebiten.KeyN) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('N')
        } else {
            err = sb.WriteByte('n')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyO) || inpututil.KeyPressDuration(ebiten.KeyO) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('O')
        } else {
            err = sb.WriteByte('o')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.KeyPressDuration(ebiten.KeyP) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('P')
        } else {
            err = sb.WriteByte('p')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.KeyPressDuration(ebiten.KeyQ) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('Q')
        } else {
            err = sb.WriteByte('q')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyR) || inpututil.KeyPressDuration(ebiten.KeyR) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('R')
        } else {
            err = sb.WriteByte('r')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.KeyPressDuration(ebiten.KeyS) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('S')
        } else {
            err = sb.WriteByte('s')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyT) || inpututil.KeyPressDuration(ebiten.KeyT) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('T')
        } else {
            err = sb.WriteByte('t')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyU) || inpututil.KeyPressDuration(ebiten.KeyU) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('U')
        } else {
            err = sb.WriteByte('u')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyV) || inpututil.KeyPressDuration(ebiten.KeyV) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('V')
        } else {
            err = sb.WriteByte('v')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.KeyPressDuration(ebiten.KeyW) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('W')
        } else {
            err = sb.WriteByte('w')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.KeyPressDuration(ebiten.KeyX) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('X')
        } else {
            err = sb.WriteByte('x')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyY) || inpututil.KeyPressDuration(ebiten.KeyY) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('Y')
        } else {
            err = sb.WriteByte('y')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyZ) || inpututil.KeyPressDuration(ebiten.KeyZ) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('Z')
        } else {
            err = sb.WriteByte('z')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit0) || inpututil.KeyPressDuration(ebiten.KeyDigit0) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte(')')
        } else {
            err = sb.WriteByte('0')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit1) || inpututil.KeyPressDuration(ebiten.KeyDigit1) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('!')
        } else {
            err = sb.WriteByte('1')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit2) || inpututil.KeyPressDuration(ebiten.KeyDigit2) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('@')
        } else {
            err = sb.WriteByte('2')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit3) || inpututil.KeyPressDuration(ebiten.KeyDigit3) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('#')
        } else {
            err = sb.WriteByte('3')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit4) || inpututil.KeyPressDuration(ebiten.KeyDigit4) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('$')
        } else {
            err = sb.WriteByte('4')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit5) || inpututil.KeyPressDuration(ebiten.KeyDigit5) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('%')
        } else {
            err = sb.WriteByte('5')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit6) || inpututil.KeyPressDuration(ebiten.KeyDigit6) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('^')
        } else {
            err = sb.WriteByte('6')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit7) || inpututil.KeyPressDuration(ebiten.KeyDigit7) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('&')
        } else {
            err = sb.WriteByte('7')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit8) || inpututil.KeyPressDuration(ebiten.KeyDigit8) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('*')
        } else {
            err = sb.WriteByte('8')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit9) || inpututil.KeyPressDuration(ebiten.KeyDigit9) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('(')
        } else {
            err = sb.WriteByte('9')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyBackspace) || inpututil.KeyPressDuration(ebiten.KeyBackspace) > 20:
        str := sb.String()
        if len(str) > 0 {
            str = str[:len(str) - 1]
            sb.Reset()
            _, err = sb.WriteString(str)
        }
    case inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.KeyPressDuration(ebiten.KeySpace) > 20:
        err = sb.WriteByte(' ')
    case inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.KeyPressDuration(ebiten.KeyEnter) > 20:
        err = sb.WriteByte('\n')
    case inpututil.IsKeyJustPressed(ebiten.KeyBackquote) || inpututil.KeyPressDuration(ebiten.KeyBackquote) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('~')
        } else {
            err = sb.WriteByte('`')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyMinus) || inpututil.KeyPressDuration(ebiten.KeyMinus) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('_')
        } else {
            err = sb.WriteByte('-')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyEqual) || inpututil.KeyPressDuration(ebiten.KeyEqual) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('+')
        } else {
            err = sb.WriteByte('=')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyBracketLeft) || inpututil.KeyPressDuration(ebiten.KeyBracketLeft) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('{')
        } else {
            err = sb.WriteByte('[')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyBracketRight) || inpututil.KeyPressDuration(ebiten.KeyBracketRight) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('}')
        } else {
            err = sb.WriteByte(']')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyBackslash) || inpututil.KeyPressDuration(ebiten.KeyBackslash) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('|')
        } else {
            err = sb.WriteByte('\\')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeySemicolon) || inpututil.KeyPressDuration(ebiten.KeySemicolon) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte(':')
        } else {
            err = sb.WriteByte(';')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyQuote) || inpututil.KeyPressDuration(ebiten.KeyQuote) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('"')
        } else {
            err = sb.WriteByte('\'')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyComma) || inpututil.KeyPressDuration(ebiten.KeyComma) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('<')
        } else {
            err = sb.WriteByte(',')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeyPeriod) || inpututil.KeyPressDuration(ebiten.KeyPeriod) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('>')
        } else {
            err = sb.WriteByte('.')
        }
    case inpututil.IsKeyJustPressed(ebiten.KeySlash) || inpututil.KeyPressDuration(ebiten.KeySlash) > 20:
        if inpututil.KeyPressDuration(ebiten.KeyShiftLeft) > 0 || inpututil.KeyPressDuration(ebiten.KeyShiftRight) > 0 {
            err = sb.WriteByte('?')
        } else {
            err = sb.WriteByte('/')
        }
    }
    return
}

func levelUp(p *player.Player) {
    switch p.Class.GetLevel() {
    case 1:
        gainProfSkill = true
    case 2:
        gainProfST = true
    case 3:
        gainProfSkill = true
    case 4:
        gainProfST = true
    case 5:
        gainProfSkill = true
    case 6:
        gainProfST = true
    case 7:
        gainProfSkill = true
    case 8:
        gainProfST = true
    case 9:
        gainProfSkill = true
    case 10:
        gainProfST = true
    case 11:
        gainProfSkill = true
    case 12:
        gainProfST = true
    case 13:
        gainProfSkill = true
    case 14:
        gainProfST = true
    case 15:
        gainProfSkill = true
    case 16:
        gainProfST = true
    case 17:
        gainProfSkill = true
    case 18:
        gainProfST = true
    case 19:
        gainProfSkill = true
    default:
        log.Print("Already max level")
    }
    if p.Class.GetLevel() < 20 {
        p.Class.LevelUp()
    }
    return
}

func init() {
    fo = utils.Fo()

    readgm.Translate(float64((768 / 2) - (724 / 2)), float64((576 / 2) - (552 / 2)))
    readimg = ebiten.NewImage(724, 552)
    readimg.Fill(color.Black)

    readgm2.Translate(float64((768 / 2) - (724 / 2) + 20), float64((576 / 2) - (552 / 2) + 20))
    readimg2 = ebiten.NewImage(724 - 40, 552 - 40)
    readimg2.Fill(color.White)

    readgm3r := text.BoundString(fo, "You do not have any written pages")
    readgm3h := readgm3r.Max.Y - readgm3r.Min.Y
    readgm3w := readgm3r.Max.X - readgm3r.Min.X
    readgm3.Translate(float64((768 / 2) - (readgm3w / 2) - 8), float64((576 / 2) - (readgm3h / 2) - 24))
    readimg3 = ebiten.NewImage(readgm3w + 16, (2 * readgm3h) + 16)
    readimg3.Fill(color.Black)

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

    fadeImage = ebiten.NewImage(768, 576)
    fadeImage.Fill(color.Black)

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
    iw, _ = overworldImage.Size()

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

    blankImage = ebiten.NewImage(768, 576)
    blankImage.Fill(color.RGBA{0x00, 0x00, 0x00, 0xb0})

    targetedBoxVert = ebiten.NewImage(2, 24)
    targetedBoxHoriz = ebiten.NewImage(24, 2)
    targetedBoxVert.Fill(color.RGBA{0xff, 0x0, 0x0, 0xff})
    targetedBoxHoriz.Fill(color.RGBA{0xff, 0x0, 0x0, 0xff})

    throwTargetBoxHoriz = ebiten.NewImage(24, 2)
    throwTargetBoxVert = ebiten.NewImage(2, 24)
    throwTargetBoxHoriz.Fill(color.RGBA{0xff, 0x0, 0x0, 0xff})
    throwTargetBoxVert.Fill(color.RGBA{0xff, 0x0, 0x0, 0xff})

    savesTableSchema = []string{"name,TEXT,1,null,1", "level,TEXT,1,\"One\",0", "x,INT,1,null,0", "y,INT,1,null,0", "csdone,TEXT,0,null,0", "inventory,TEXT,0,null,0", "stats,TEXT,0,null,0", "equipment,TEXT,0,null,0"}
    pagesTableSchema = []string{"name,TEXT,1,null,0", "msg,TEXT,1,null,0", "charname,TEXT,1,null,0"}
    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    db, err := sql.Open("sqlite3", homeDir + "/quailsaves.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    var savesCreateStmt string = "create table if not exists saves ("
    for cind, col := range savesTableSchema {
        colArr := strings.Split(col, ",")
        savesCreateStmt += colArr[0] + " " + colArr[1]
        if colArr[2] == "1" {
            savesCreateStmt += " not null"
        }
        if colArr[3] != "null" {
            savesCreateStmt += " default " + colArr[3]
        }
        if colArr[4] == "1" {
            savesCreateStmt += " primary key"
        }
        if cind == len(savesTableSchema) - 1 {
            savesCreateStmt += ");"
        } else {
            savesCreateStmt += ", "
        }
    }
    _, err = db.Exec(savesCreateStmt)
    if err != nil {
        log.Fatal(fmt.Sprintf("%q: %s\n", err, savesCreateStmt))
    }
    var pagesCreateStmt string = "create table if not exists pages ("
    for cind, col := range pagesTableSchema {
        colArr := strings.Split(col, ",")
        pagesCreateStmt += colArr[0] + " " + colArr[1]
        if colArr[2] == "1" {
            pagesCreateStmt += " not null"
        }
        if colArr[3] != "null" {
            pagesCreateStmt += " default " + colArr[3]
        }
        if colArr[4] == "1" {
            pagesCreateStmt += " primary key"
        }
        if cind == len(pagesTableSchema) - 1 {
            pagesCreateStmt += ");"
        } else {
            pagesCreateStmt += ", "
        }
    }
    _, err = db.Exec(pagesCreateStmt)
    if err != nil {
        log.Fatal(fmt.Sprintf("%q: %s\n", err, pagesCreateStmt))
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
