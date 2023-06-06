package classes

import (
    "fmt"
    "log"
    "strconv"

//    "github.com/jsnider-mtu/quailgame/player"
)

type Quail struct {
    stats *Stats
    level int
    xp int
    illuminated [3]int
}

func (q *Quail) Create(abilities [6]int) bool {
    q.stats = &Stats{}
    for i, x := range abilities {
        if x > 18 || x < 3 {
            log.Fatal(fmt.Sprintf("Ability score %d at index %d is invalid", x, i))
        }
    }
    q.stats.str = abilities[0]
    q.stats.con = abilities[1]
    q.stats.dex = abilities[2]
    q.stats.intel = abilities[3]
    q.stats.wis = abilities[4]
    q.stats.cha = abilities[5]
    if q.stats.str < 10 && q.stats.str % 2 == 1 {
        q.stats.strmod = ((q.stats.str - 10) / 2) - 1
    } else {
        q.stats.strmod = (q.stats.str - 10) / 2
    }
    if q.stats.dex < 10 && q.stats.dex % 2 == 1 {
        q.stats.dexmod = ((q.stats.dex - 10) / 2) - 1
    } else {
        q.stats.dexmod = (q.stats.dex - 10) / 2
    }
    if q.stats.con < 10 && q.stats.con % 2 == 1 {
        q.stats.conmod = ((q.stats.con - 10) / 2) - 1
    } else {
        q.stats.conmod = (q.stats.con - 10) / 2
    }
    if q.stats.intel < 10 && q.stats.intel % 2 == 1 {
        q.stats.intelmod = ((q.stats.intel - 10) / 2) - 1
    } else {
        q.stats.intelmod = (q.stats.intel - 10) / 2
    }
    if q.stats.wis < 10 && q.stats.wis % 2 == 1 {
        q.stats.wismod = ((q.stats.wis - 10) / 2) - 1
    } else {
        q.stats.wismod = (q.stats.wis - 10) / 2
    }
    if q.stats.cha < 10 && q.stats.cha % 2 == 1 {
        q.stats.chamod = ((q.stats.cha - 10) / 2) - 1
    } else {
        q.stats.chamod = (q.stats.cha - 10) / 2
    }
    q.stats.pb = 2
    q.level = 1
    q.xp = 0
    return true
}

func (q *Quail) GetName() string {
    return "Quail"
}

func (q *Quail) GetStr() [2]int {
    return [2]int{q.stats.str, q.stats.strmod}
}

func (q *Quail) GetDex() [2]int {
    return [2]int{q.stats.dex, q.stats.dexmod}
}

func (q *Quail) GetCon() [2]int {
    return [2]int{q.stats.con, q.stats.conmod}
}

func (q *Quail) GetIntel() [2]int {
    return [2]int{q.stats.intel, q.stats.intelmod}
}

func (q *Quail) GetWis() [2]int {
    return [2]int{q.stats.wis, q.stats.wismod}
}

func (q *Quail) GetCha() [2]int {
    return [2]int{q.stats.cha, q.stats.chamod}
}

func (q *Quail) GetPB() int {
    return q.stats.pb
}

func (q *Quail) GetLevel() int {
    return q.level
}

func (q *Quail) GetXP() int {
    return q.xp
}

func (q *Quail) SavingThrow(ability string) int {
    return 0
}

func (q *Quail) SkillCheck(skill string) int {
    return 0
}

func (q *Quail) CanSpeak(language string) bool {
    return false
}

func (q *Quail) Illuminate(data [3]int) bool {
    q.illuminated = data
    return true
}

func (q *Quail) Illuminated() (bool, [3]int) {
    if q.illuminated == [3]int{0, 0, 0} {
        return false, q.illuminated
    }
    return true, q.illuminated
}

func (q *Quail) Save() string {
    var finalstr string
    for _, num := range q.stats.Save() {
        finalstr += strconv.Itoa(num)
        finalstr += ";"
    }
    return finalstr + strconv.Itoa(q.level) + ";" + strconv.Itoa(q.xp) + ";"
}
