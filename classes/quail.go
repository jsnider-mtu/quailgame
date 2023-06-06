package classes

import (
    "fmt"
    "log"
    "strconv"
)

type Quail struct {
    stats *Stats
    proficiencies []string
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
    strval := q.stats.strmod
    dexval := q.stats.dexmod
    conval := q.stats.conmod
    intelval := q.stats.intelmod
    wisval := q.stats.wismod
    chaval := q.stats.chamod
    for _, v := range q.proficiencies {
        switch v {
        case "str":
            strval += q.stats.pb
        case "dex":
            dexval += q.stats.pb
        case "con":
            conval += q.stats.pb
        case "intel":
            intelval += q.stats.pb
        case "wis":
            wisval += q.stats.pb
        case "cha":
            chaval += q.stats.pb
        }
    }
    switch ability {
    case "str":
        return strval
    case "dex":
        return dexval
    case "con":
        return conval
    case "intel":
        return intelval
    case "wis":
        return wisval
    case "cha":
        return chaval
    default:
        log.Fatal(fmt.Sprintf("%s is not a valid ability (saving throw)", ability))
    }
    return 0
}

func (q *Quail) SkillCheck(skill string) int {
    acr := q.stats.dexmod
    ani := q.stats.wismod
    arc := q.stats.intelmod
    ath := q.stats.strmod
    dec := q.stats.chamod
    his := q.stats.intelmod
    ins := q.stats.wismod
    inti := q.stats.chamod
    inv := q.stats.intelmod
    med := q.stats.wismod
    nat := q.stats.intelmod
    perc := q.stats.wismod
    perf := q.stats.chamod
    pers := q.stats.chamod
    rel := q.stats.intelmod
    sle := q.stats.dexmod
    ste := q.stats.dexmod
    sur := q.stats.wismod
    for _, v := range q.proficiencies {
        switch v {
        case "acrobatics":
            acr += q.stats.pb
        case "animalhandling":
            ani += q.stats.pb
        case "arcana":
            arc += q.stats.pb
        case "athletics":
            ath += q.stats.pb
        case "deception":
            dec += q.stats.pb
        case "history":
            his += q.stats.pb
        case "insight":
            ins += q.stats.pb
        case "intimidation":
            inti += q.stats.pb
        case "investigation":
            inv += q.stats.pb
        case "medicine":
            med += q.stats.pb
        case "nature":
            nat += q.stats.pb
        case "perception":
            perc += q.stats.pb
        case "performance":
            perf += q.stats.pb
        case "persuasion":
            pers += q.stats.pb
        case "religion":
            rel += q.stats.pb
        case "sleightofhand":
            sle += q.stats.pb
        case "stealth":
            ste += q.stats.pb
        case "survival":
            sur += q.stats.pb
        }
    }
    switch skill {
    case "acrobatics":
        return acr
    case "animalhandling":
        return ani
    case "arcana":
        return arc
    case "athletics":
        return ath
    case "deception":
        return dec
    case "history":
        return his
    case "insight":
        return ins
    case "intimidation":
        return inti
    case "investigation":
        return inv
    case "medicine":
        return med
    case "nature":
        return nat
    case "perception":
        return perc
    case "performance":
        return perf
    case "persuasion":
        return pers
    case "religion":
        return rel
    case "sleightofhand":
        return sle
    case "stealth":
        return ste
    case "survival":
        return sur
    default:
        log.Fatal(fmt.Sprintf("%s is not a valid skill", skill))
    }
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

func (q *Quail) EarnXP(earnedxp int) {
    q.xp += earnedxp
    return
}

func (q *Quail) AddProf(trait string) bool {
    for _, val := range q.proficiencies {
        if trait == val {
            return false
        }
    }
    switch trait {
    case "str":
        q.proficiencies = append(q.proficiencies, "str")
    case "dex":
        q.proficiencies = append(q.proficiencies, "dex")
    case "con":
        q.proficiencies = append(q.proficiencies, "con")
    case "intel":
        q.proficiencies = append(q.proficiencies, "intel")
    case "wis":
        q.proficiencies = append(q.proficiencies, "wis")
    case "cha":
        q.proficiencies = append(q.proficiencies, "cha")
    case "acrobatics":
        q.proficiencies = append(q.proficiencies, "acrobatics")
    case "animalhandling":
        q.proficiencies = append(q.proficiencies, "animalhandling")
    case "arcana":
        q.proficiencies = append(q.proficiencies, "arcana")
    case "athletics":
        q.proficiencies = append(q.proficiencies, "athletics")
    case "deception":
        q.proficiencies = append(q.proficiencies, "deception")
    case "history":
        q.proficiencies = append(q.proficiencies, "history")
    case "insight":
        q.proficiencies = append(q.proficiencies, "insight")
    case "intimidation":
        q.proficiencies = append(q.proficiencies, "intimidation")
    case "investigation":
        q.proficiencies = append(q.proficiencies, "investigation")
    case "medicine":
        q.proficiencies = append(q.proficiencies, "medicine")
    case "nature":
        q.proficiencies = append(q.proficiencies, "nature")
    case "perception":
        q.proficiencies = append(q.proficiencies, "perception")
    case "performance":
        q.proficiencies = append(q.proficiencies, "performance")
    case "persuasion":
        q.proficiencies = append(q.proficiencies, "persuasion")
    case "religion":
        q.proficiencies = append(q.proficiencies, "religion")
    case "sleightofhand":
        q.proficiencies = append(q.proficiencies, "sleightofhand")
    case "stealth":
        q.proficiencies = append(q.proficiencies, "stealth")
    case "survival":
        q.proficiencies = append(q.proficiencies, "survival")
    default:
        log.Fatal(fmt.Sprintf("%s is not a skill or saving throw", trait))
    }
    return true
}

func (q *Quail) LevelUp() {
    if q.level < 20 {
        q.level += 1
    } else {
        log.Print(fmt.Sprintf("Level is %d, cannot level more", q.level))
    }
    return
}
