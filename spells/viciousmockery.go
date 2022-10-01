package spells

type ViciousMockery struct {}

func (v ViciousMockery) Cast(target string) bool {
    log.Println("The spell Vicious Mockery is not implemented yet")
}

func (v ViciousMockery) PrettyPrint() string {
    return "Vicious Mockery"
}
