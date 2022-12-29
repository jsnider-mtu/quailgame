package items

import (
    "fmt"
    "log"
    "strconv"
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
        page := &Page{msg: msg}
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
