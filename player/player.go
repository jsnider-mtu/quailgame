package player

import (
//    "log"

    "github.com/hajimehoshi/ebiten/v2"
//    "github.com/jsnider-mtu/quailgame/levels"
)

//var (
//    P = &Player{}
//)

type Player struct {
    Pos [2]int
    Image *ebiten.Image
}

//func (p *Player) StartLoc(l *levels.Level, pos [2]int) {
//    // log.Fatal if pos is out of bounds or in a box
//    if pos[0] >= 0 && pos[1] >= 0 && pos[0] <= l.Max[0] - 48 && pos[1] <= l.Max[1] - 48 {
//        // check if in a box
//        for _, a := range l.Boxes {
//            if pos[0] >= a[0] && pos[1] >= a[1] && pos[0] < a[2] && pos[1] < a[3] - 24 {
//                log.Fatal("StartLoc is in a box")
//            }
//        }
//    } else {
//        log.Fatal("StartLoc is out of bounds")
//    }
//    p.Pos = pos
//}
