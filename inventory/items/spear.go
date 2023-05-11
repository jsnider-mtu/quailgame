package items

type Spear struct {
    Currslot string
}

func (s *Spear) GetCurrSlot() string {
    return s.Currslot
}

func (s *Spear) SwitchSlots() {
    if s.Currslot == "BothHands" {
        s.Currslot = "RightHand"
    } else if s.Currslot == "RightHand" {
        s.Currslot = "BothHands"
    }
    return
}

func (s *Spear) Slot() string {
    return "Versatile"
}

func (s *Spear) Use() (string, []int) {
    return s.Action(), []int{96, 288}
}

func (s *Spear) Save() string {
    return "Spear"
}

func (s *Spear) PrettyPrint() string {
    return "Spear"
}

func (s *Spear) Function() string {
    return "melee-thrown"
}

func (s *Spear) Damage() (int, int, string) {
    return 1, 6, "piercing"
}

func (s *Spear) Action() string {
    return "throw"
}

func (s *Spear) GetQuantity() int {
    return 1
}

func (s *Spear) GetRange() []float64 {
    return []float64{96.0, 288.0}
}
