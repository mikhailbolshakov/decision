package sys

import (
	"github.com/mikhailbolshakov/decision/http"
)

func GetRoutes(c Controller) []*http.Route {
	return []*http.Route{
		http.R("/health", c.Health).GET().NoAuth(),
	}
}
