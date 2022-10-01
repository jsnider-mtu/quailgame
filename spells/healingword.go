package spells

type HealingWord struct {}

func (h HealingWord) Cast(target string) bool {
    log.Println("The spell Healing Word is not implemented yet")
}

func (h HealingWord) PrettyPrint() string {
    return "Healing Word"
}
