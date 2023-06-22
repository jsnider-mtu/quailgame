package quests

type Quest interface {
    Status() string
    TurnIn() bool
    GetRewards() (int, []items.Item)
    GetLevelReq() int
    GetXP() int
    GetID() (int, string)
    GetDescription() string
    GetObjective() (int, string)
    CompleteObjective(int) bool
}
