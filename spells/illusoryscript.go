package spells

type IllusoryScript struct {}

func (i IllusoryScript) Cast(target string) bool {
    log.Println("The spell Illusory Script is not implemented yet")
}

func (i IllusoryScript) PrettyPrint() string {
    return "Illusory Script"
}
