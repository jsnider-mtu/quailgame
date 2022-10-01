package spells

type Thunderwave struct {}

func (t Thunderwave) Cast(target string) bool {
    log.Println("The spell Thunderwave is not implemented yet")
}

func (t Thunderwave) PrettyPrint() string {
    return "Thunderwave"
}
