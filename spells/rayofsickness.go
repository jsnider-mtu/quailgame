package spells

type RayOfSickness struct {}

func (r RayOfSickness) Cast(target string) bool {
    log.Println("The spell Ray of Sickness is not implemented yet")
}

func (r RayOfSickness) PrettyPrint() string {
    return "Ray of Sickness"
}
