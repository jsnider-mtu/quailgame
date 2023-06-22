package npcs

import (
    "github.com/jsnider-mtu/quailgame/player"
)

type JaneDoe struct {
    name string
    speed int
    offset int
    direction string
    stopped bool
    msgs [][]string
    msgcount int
    PC *player.Player
}

func (j *JaneDoe) GetPC() *player.Player {
    return j.PC
}

func (j *JaneDoe) GetName() string {
    return j.name
}

func (j *JaneDoe) GetSpeed() int {
    return j.speed
}

func (j *JaneDoe) GetOffset() int {
    return j.offset
}

func (j *JaneDoe) GetDirection() string {
    return j.direction
}

func (j *JaneDoe) GetStopped() bool {
    return j.stopped
}

func (j *JaneDoe) Dialog() []string {
    if j.msgcount == len(j.msgs) {
        j.msgcount = 0
    }
    j.msgcount++
    return j.msgs[j.msgcount - 1]
}

func (j *JaneDoe) Direction(d string) {
    j.direction = d
    return
}

func (j *JaneDoe) Stopped(s bool) {
    j.stopped = s
    return
}
