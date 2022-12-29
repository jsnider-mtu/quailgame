package items

import (
    "fmt"
    "log"
    "strconv"
    "strings"
)

type Page struct {
    name string
    msg string
}

func (p *Page) GetName() string {
    return p.name
}

func (p *Page) Read() string {
    return p.msg
}

type Paper struct {
    Quantity int
    pages []*Page
}

func (p *Paper) Slot() string {
    return ""
}

func (p *Paper) Use() (string, []int) {
    return "read", []int{}
}

func (p *Paper) Save() string {
    return "Paper," + strconv.Itoa(p.Quantity)
}

func (p *Paper) PrettyPrint() string {
    return fmt.Sprintf("Paper (%d)", p.Quantity)
}

func (p *Paper) Function() string {
    return "writing"
}

func (p *Paper) Damage() (int, int, string) {
    return 0, 0, ""
}

func (p *Paper) Action() string {
    return "read"
}

func (p *Paper) GetQuantity() int {
    return p.Quantity
}

func (p *Paper) Write(msg string) {
    if p.GetQuantity() > 0 {
        log.Println(fmt.Sprintf("p.GetQuantity() == %d", p.GetQuantity()))
        var y int
        result := ""
        lines := strings.Split(msg, "\n")
        for ind, line := range lines {
            if len(line) > 55 {
                for x := 54; x < len(line); x = y + 56 {
                    for y = x; line[y] != ' '; y-- {
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
        page := &Page{msg: result}
        p.Quantity--
        p.pages = append(p.pages, page)
    } else {
        log.Println(fmt.Sprintf("p.GetQuantity() == %d", p.GetQuantity()))
    }
}

func (p *Paper) GetPages() []*Page {
    return p.pages
}

func (p *Paper) AddPaper(amount int) {
    p.Quantity += amount
    return
}
