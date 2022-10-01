package spells

type ArmsOfHadar struct {}

func (a ArmsOfHadar) Cast(target string) bool {
    log.Println("The spell Arms of Hadar is not implemented yet")
}

func (a ArmsOfHadar) PrettyPrint() string {
    return "Arms of Hadar"
}
