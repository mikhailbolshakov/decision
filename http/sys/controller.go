package sys

import (
	"github.com/mikhailbolshakov/decision"
	kitHttp "github.com/mikhailbolshakov/decision/kit/http"
	"net/http"
)

type Controller interface {
	kitHttp.Controller
	Health(http.ResponseWriter, *http.Request)
}

type ctrlImpl struct {
	kitHttp.BaseController
}

func NewController() Controller {
	return &ctrlImpl{
		BaseController: kitHttp.BaseController{Logger: decision.LF()},
	}
}

func (c *ctrlImpl) Health(w http.ResponseWriter, r *http.Request) {
	c.RespondOK(w, kitHttp.EmptyOkResponse)
}
