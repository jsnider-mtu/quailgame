package spells

type HuntersMark struct {}

func (h HuntersMark) Cast(target string) bool {
    log.Println("The spell Hunter's Mark is not implemented yet")
}

func (h HuntersMark) PrettyPrint() string {
    return "Hunter's Mark"
}
