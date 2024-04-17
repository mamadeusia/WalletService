package entity

type Vote struct {
	UserId     string
	Type       VoteType
	ProposalId string
}

type VoteType int

const (
	For VoteType = iota
	Against
	Abstain
)
