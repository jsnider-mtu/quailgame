package utils

import (
    "github.com/jsnider-mtu/projectx/levels"
    "github.com/jsnider-mtu/projectx/player"
)

func TryUpdatePos(pc bool, p *player.Player, l *levels.Level, vert bool, dist int, mc *player.Player) bool {
    if vert {
        // up
        if dist < 0 {
            if p.Pos[1] + dist > l.Pos[1] && p.Pos[1] + dist < l.Max[1] {
                for _, a := range l.Boxes {
                    if p.Pos[0] + 48 > a[0] && p.Pos[1] + dist >= a[1] && p.Pos[0] < a[2] && p.Pos[1] + dist < a[3] - 24 {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if p.Pos[0] + 24 > b.PC.Pos[0] && p.Pos[1] + dist > b.PC.Pos[1] - 24 && p.Pos[0] < b.PC.Pos[0] + 24 && p.Pos[1] - 24 + dist < b.PC.Pos[1] {
                        return false
                    }
                }
                if !pc {
                    if p.Pos[0] + 24 > mc.Pos[0] && p.Pos[1] + dist > mc.Pos[1] - 24 && p.Pos[0] < mc.Pos[0] + 24 && p.Pos[1] - 24 + dist < mc.Pos[1] {
                        return false
                    }
                }
                p.Pos[1] += dist
                if pc {
                    l.Pos[1] -= dist
                }
                return true
            }
            return false
        } else {
            // down
            if p.Pos[1] + dist > l.Pos[1] && p.Pos[1] + dist < l.Max[1] {
                for _, a := range l.Boxes {
                    if p.Pos[0] + 48 > a[0] && p.Pos[1] + 48 + dist > a[1] && p.Pos[0] < a[2] && p.Pos[1] + 48 + dist < a[3] {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if p.Pos[0] + 24 > b.PC.Pos[0] && p.Pos[1] + dist > b.PC.Pos[1] - 24 && p.Pos[0] < b.PC.Pos[0] + 24 && p.Pos[1] + 24 + dist < b.PC.Pos[1] + 48 {
                        return false
                    }
                }
                if !pc {
                    if p.Pos[0] + 24 > mc.Pos[0] && p.Pos[1] + dist > mc.Pos[1] - 24 && p.Pos[0] < mc.Pos[0] + 24 && p.Pos[1] + 24 + dist < mc.Pos[1] + 48 {
                        return false
                    }
                }
                p.Pos[1] += dist
                if pc {
                    l.Pos[1] -= dist
                }
                return true
            }
            return false
        }
    } else {
        // left
        if dist < 0 {
            if p.Pos[0] + dist > l.Pos[0] && p.Pos[0] + dist < l.Max[0] {
                for _, a := range l.Boxes {
                    if p.Pos[0] + dist > a[0] && p.Pos[1] + 48 > a[1] && p.Pos[0] + dist < a[2] && p.Pos[1] < a[3] - 24 {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if p.Pos[0] + dist > b.PC.Pos[0] - 24 && p.Pos[1] > b.PC.Pos[1] - 24 && p.Pos[0] + dist < b.PC.Pos[0] + 24 && p.Pos[1] < b.PC.Pos[1] + 24 {
                        return false
                    }
                }
                if !pc {
                    if p.Pos[0] + dist > mc.Pos[0] - 24 && p.Pos[1] > mc.Pos[1] - 24 && p.Pos[0] + dist < mc.Pos[0] + 24 && p.Pos[1] < mc.Pos[1] + 24 {
                        return false
                    }
                }
                p.Pos[0] += dist
                if pc {
                    l.Pos[0] -= dist
                }
                return true
            }
            return false
        } else {
            // right
            if p.Pos[0] + dist > l.Pos[0] && p.Pos[0] + dist < l.Max[0] {
                for _, a := range l.Boxes {
                    if p.Pos[0] + 48 + dist > a[0] && p.Pos[1] + 48 > a[1] && p.Pos[0] + 48 + dist < a[2] && p.Pos[1] < a[3] - 24 {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if p.Pos[0] + dist > b.PC.Pos[0] - 24 && p.Pos[1] > b.PC.Pos[1] - 24 && p.Pos[0] + dist < b.PC.Pos[0] + 24 && p.Pos[1] < b.PC.Pos[1] + 24 {
                        return false
                    }
                }
                if !pc {
                    if p.Pos[0] + dist > mc.Pos[0] - 24 && p.Pos[1] > mc.Pos[1] - 24 && p.Pos[0] + dist < mc.Pos[0] + 24 && p.Pos[1] < mc.Pos[1] + 24 {
                        return false
                    }
                }
                p.Pos[0] += dist
                if pc {
                    l.Pos[0] -= dist
                }
                return true
            }
            return false
        }
    }
}
