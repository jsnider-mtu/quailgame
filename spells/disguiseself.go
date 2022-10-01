package spells

type DisguiseSelf struct {}

func (d DisguiseSelf) Cast(target string) bool {
    log.Println("The spell Disguise Self is not implemented yet")
}

func (d DisguiseSelf) PrettyPrint() string {
    return "Disguise Self"
}
