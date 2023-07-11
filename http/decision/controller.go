package decision

import (
	"github.com/mikhailbolshakov/decision"
	domain "github.com/mikhailbolshakov/decision/domain/decision"
	kitHttp "github.com/mikhailbolshakov/decision/kit/http"
	"net/http"
)

type Controller interface {
	kitHttp.Controller
	MakeDecision(http.ResponseWriter, *http.Request)
	MakeDecisionGuest(http.ResponseWriter, *http.Request)
}

type ctrlImpl struct {
	kitHttp.BaseController
	decisionService domain.DecisionService
}

func NewController(decisionService domain.DecisionService) Controller {
	return &ctrlImpl{
		decisionService: decisionService,
		BaseController:  kitHttp.BaseController{Logger: decision.LF()},
	}
}

func (c *ctrlImpl) MakeDecision(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, err := c.UserIdVar(ctx, r, "userId")
	if err != nil {
		c.RespondError(w, err)
		return
	}

	rq := &Problem{}
	if err = c.DecodeRequest(ctx, r, rq); err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.decisionService.MakeDecision(ctx, userId, c.toProblemDomain(rq))
	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toDecisionResultApi(res))
}

func (c *ctrlImpl) MakeDecisionGuest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rq := &Problem{}
	if err := c.DecodeRequest(ctx, r, rq); err != nil {
		c.RespondError(w, err)
		return
	}

	res, err := c.decisionService.MakeDecision(ctx, "", nil)
	if err != nil {
		c.RespondError(w, err)
		return
	}

	c.RespondOK(w, c.toDecisionResultApi(res))
}
