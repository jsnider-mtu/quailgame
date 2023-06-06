package classes

import (
    "sort"
)

type Class interface {
    Create([6]int) bool
    GetName() string
    GetLevel() int
    GetXP() int
    GetStr() [2]int
    GetDex() [2]int
    GetCon() [2]int
    GetIntel() [2]int
    GetWis() [2]int
    GetCha() [2]int
    GetPB() int
    SavingThrow(string) int
    SkillCheck(string) int
    CanSpeak(string) bool
    Illuminate([3]int) bool
    Illuminated() (bool, [3]int)
    Save() string
    EarnXP(int)
    AddProf(string) bool
    LevelUp()
}

type Stats struct {
    str int
    strmod int
    dex int
    dexmod int
    con int
    conmod int
    intel int
    intelmod int
    wis int
    wismod int
    cha int
    chamod int
    pb int
}

func (s *Stats) Save() []int {
    abilities := []int{s.str, s.dex, s.con, s.intel, s.wis, s.cha}
    sort.Slice(abilities, func(i, j int) bool {
        return abilities[i] > abilities[j]
    })
    return abilities
}
