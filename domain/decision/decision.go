package domain

import "context"

type Quality struct {
	Id          string
	Name        string
	Importance  float64
	Probability float64
}

type Option struct {
	Id   string
	Name string
	Pros []*Quality
	Cons []*Quality
}

type Problem struct {
	Id      string
	Name    string
	Options []*Option
}

type DecisionResult struct {
	OptionsRating map[string]float64
}

type Decision struct {
	Id        string
	ProblemId string
	UserId    string
	Result    DecisionResult
}

type DecisionService interface {
	// MakeDecision makes decision for the problem
	MakeDecision(ctx context.Context, userId string, problem *Problem) (*Decision, error)
}
