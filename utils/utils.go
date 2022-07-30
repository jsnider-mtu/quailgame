package utils

import (
    "github.com/jsnider-mtu/projectx/levels"
    "github.com/jsnider-mtu/projectx/player"
)

func TryUpdatePos(pc bool, p *player.Player, l *levels.Level, vert bool, dist int, mc *player.Player) bool {
    if vert {
        if p.Pos[1] + dist > l.Pos[1] && p.Pos[1] + dist < l.Max[1] {
            if dist < 0 {
                if !pc {
                    if p.Pos[0] == mc.Pos[0] && p.Pos[1] + dist > mc.Pos[1] && p.Pos[1] + dist < mc.Pos[1] + 48 {
                        return false
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] > a[0] - 48 && p.Pos[0] < a[2] && p.Pos[1] + dist >= a[1] && p.Pos[1] + dist < a[3] - 24 {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist < b.PC.Pos[1] + 48 {
                            return false
                        }
                    } else {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist == b.PC.Pos[1] {
                            return false
                        }
                    }
                }
            } else {
                if !pc {
                    if p.Pos[0] == mc.Pos[0] && p.Pos[1] + dist > mc.Pos[1] - 48 && p.Pos[1] + dist < mc.Pos[1] + 48 {
                        return false
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] > a[0] - 24 && p.Pos[0] < a[2] - 24 && p.Pos[1] + dist >= a[1] - 24 && p.Pos[1] + dist < a[3] {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist > b.PC.Pos[1] - 48 {
                            return false
                        }
                    } else {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] + dist == b.PC.Pos[1] {
                            return false
                        }
                    }
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
        if p.Pos[0] + dist > l.Pos[0] && p.Pos[0] + dist < l.Max[0] {
            if dist < 0 {
                if !pc {
                    if p.Pos[0] + dist == mc.Pos[0] && p.Pos[1] >= mc.Pos[1] - 24 && p.Pos[1] <= mc.Pos[1] + 24 {
                        return false
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] + dist >= a[0] && p.Pos[0] + dist < a[2] && p.Pos[1] >= a[1] - 24 && p.Pos[1] < a[3] - 24 {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] + dist == b.PC.Pos[0] && p.Pos[1] >= b.PC.Pos[1] - 24 && p.Pos[1] <= b.PC.Pos[1] + 24 {
                            return false
                        }
                    } else {
                        if p.Pos[0] + dist >= b.PC.Pos[0] && p.Pos[0] + dist < b.PC.Pos[0] + 24 && p.Pos[1] == b.PC.Pos[1] {
                            return false
                        }
                    }
                }
            } else {
                if !pc {
                    if p.Pos[0] + dist == mc.Pos[0] && p.Pos[1] >= mc.Pos[1] - 24 && p.Pos[1] <= mc.Pos[1] + 24 {
                        return false
                    }
                }
                for _, a := range l.Boxes {
                    if p.Pos[0] + dist >= a[0] - 24 && p.Pos[0] + dist < a[2] && p.Pos[1] >= a[1] - 24 && p.Pos[1] < a[3] - 24 {
                        return false
                    }
                }
                for _, b := range l.NPCs {
                    if !pc {
                        if p.Pos[0] == b.PC.Pos[0] && p.Pos[1] == b.PC.Pos[1] {
                            continue
                        }
                        if p.Pos[0] + dist >= b.PC.Pos[0] && p.Pos[1] >= b.PC.Pos[1] - 24 && p.Pos[1] <= b.PC.Pos[1] + 24 {
                            return false
                        }
                    } else {
                        if p.Pos[0] + dist >= b.PC.Pos[0] && p.Pos[0] + dist < b.PC.Pos[0] + 24 && p.Pos[1] == b.PC.Pos[1] {
                            return false
                        }
                    }
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

