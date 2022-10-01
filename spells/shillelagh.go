package spells

type Shillelagh struct {}

func (s Shillelagh) Cast(target string) bool {
    log.Println("The spell Shillelagh is not implemented yet")
}

func (s Shillelagh) PrettyPrint() string {
    return "Shillelagh"
}
