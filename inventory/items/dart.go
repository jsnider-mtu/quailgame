package items

type Dart struct {
}

func (d Dart) Slot() string {
    return "RightHand"
}

func (d Dart) Use() {
    // must be equipped to use
}

func (d Dart) Save() string {
    return "Dart"
}
