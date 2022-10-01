package spells

type RayOfFrost struct {}

func (r RayOfFrost) Cast(target string) bool {
    log.Println("The spell Ray of Frost is not implemented yet")
}

func (r RayOfFrost) PrettyPrint() string {
    return "Ray of Frost"
}
