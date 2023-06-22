package npcs

import (
    "github.com/jsnider-mtu/quailgame/player"
)

type Wizard struct {
    name string
    speed int
    offset int
    direction string
    stopped bool
    msgs [][]string
    msgcount int
    PC *player.Player
}

func (w *Wizard) GetPC() *player.Player {
    return w.PC
}

func (w *Wizard) GetName() string {
    return w.name
}

func (w *Wizard) GetSpeed() int {
    return w.speed
}

func (w *Wizard) GetOffset() int {
    return w.offset
}

func (w *Wizard) GetDirection() string {
    return w.direction
}

func (w *Wizard) GetStopped() bool {
    return w.stopped
}

func (w *Wizard) Dialog() []string {
    if w.msgcount == len(w.msgs) {
        w.msgcount = 0
    }
    w.msgcount++
    return w.msgs[w.msgcount - 1]
}

func (w *Wizard) Direction(d string) {
    w.direction = d
    return
}

func (w *Wizard) Stopped(s bool) {
    w.stopped = s
    return
}
