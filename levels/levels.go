package levels

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/jsnider-mtu/projectx/main"
)

type Level struct {
    Max [2]int
    Pos [2]int
    Boxes [][4]int
    Doors []*Door
    Image *ebiten.Image
}

type Door struct {
    Coords [4]int
    Direction string // "up", "down", "left", "right"
    Image *ebiten.Image
}

func TryUpdatePos(pc bool, pos [2]int, l *Level, vert bool, dist int) [2]int {
    if vert {
        // up
        if dist < 0 {
            if pos[1] + dist > 0 && pos[1] + dist < l.Max[1] {
                for _, a := range l.Boxes {
                    if pos[0] + 48 > a[0] && pos[1] + dist > a[1] && pos[0] < a[2] && pos[1] + dist < a[3] - 24 {
                        return [2]int{-1, -1}
                    }
                }
                pos[1] += dist
                if pc {
                    l.Pos[1] -= dist
                    main.P.Pos[0] += dist
                }
                return pos
            }
            return [2]int{-1, -1}
        } else {
            // down
            if pos[1] + dist > 0 && pos[1] + dist < l.Max[1] {
                for _, a := range l.Boxes {
                    if pos[0] + 48 > a[0] && pos[1] + 48 + dist > a[1] && pos[0] < a[2] && pos[1] + 48 + dist < a[3] {
                        return [2]int{-1, -1}
                    }
                }
                pos[1] += dist
                if pc {
                    l.Pos[1] -= dist
                    p.Pos[0] += dist
                }
                return pos
            }
            return [2]int{-1, -1}
        }
    } else {
        // left
        if dist < 0 {
            if pos[0] + dist > 0 && pos[0] + dist < l.Max[0] {
                for _, b := range l.Boxes {
                    if pos[0] + dist > b[0] && pos[1] + 48 > b[1] && pos[0] + dist < b[2] && pos[1] < b[3] - 24 {
                        return [2]int{-1, -1}
                    }
                }
                pos[0] += dist
                if pc {
                    l.Pos[0] -= dist
                    p.Pos[0] += dist
                }
                return pos
            }
            return [2]int{-1, -1}
        } else {
            // right
            if pos[0] + dist > 0 && pos[0] + dist < l.Max[0] {
                for _, b := range l.Boxes {
                    if pos[0] + 48 + dist > b[0] && pos[1] + 48 > b[1] && pos[0] + 48 + dist < b[2] && pos[1] < b[3] - 24 {
                        return [2]int{-1, -1}
                    }
                }
                pos[0] += dist
                if pc {
                    l.Pos[0] -= dist
                    p.Pos[0] += dist
                }
                return pos
            }
            return [2]int{-1, -1}
        }
    }
}
