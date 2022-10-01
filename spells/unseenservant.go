package spells

type UnseenServant struct {}

func (u UnseenServant) Cast(target string) bool {
    log.Println("The spell Unseen Servant is not implemented yet")
}

func (u UnseenServant) PrettyPrint() string {
    return "Unseen Servant"
}
