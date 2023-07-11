package impl

import (
	"context"
	domain "github.com/mikhailbolshakov/decision/domain/decision"
	"github.com/mikhailbolshakov/decision/kit"
)

type decisionServiceImpl struct {
}

func NewDecisionService() domain.DecisionService {
	return &decisionServiceImpl{}
}

func (p *decisionServiceImpl) MakeDecision(ctx context.Context, userId string, problem *domain.Problem) (*domain.Decision, error) {

	r := &domain.Decision{
		Id:        kit.NewRandString(),
		ProblemId: problem.Id,
		UserId:    userId,
	}

	for _, op := range problem.Options {
		//
		kCon, kPro := 0.0, 0.0
		for _, con := range op.Cons {
			kCon += con.Importance * con.Probability
		}
		//
		for _, pro := range op.Pros {
			kPro += pro.Importance * pro.Probability
		}
		r.Result.OptionsRating[op.Id] = kit.Round100(kPro / kCon)
	}

	return r, nil
}
