package spells

type Friends struct {}

func (f Friends) Cast(target string) bool {
    log.Println("The spell Friends is not implemented yet")
}

func (f Friends) PrettyPrint() string {
    return "Friends"
}
