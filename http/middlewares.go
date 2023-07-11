package http

import (
	"fmt"
	"github.com/mikhailbolshakov/decision/kit"
	kitHttp "github.com/mikhailbolshakov/decision/kit/http"
	"net/http"
	"time"
)

const (
	HeaderXRealIp       = "x-real-ip"
	HeaderXForwarderFor = "x-forwarder-for"
	HeaderRequestId     = "x-request-id"
)

type Middleware struct {
	kitHttp.BaseController
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) AuthAccessTokenMiddleware(next http.HandlerFunc, tokenTypes ...string) http.HandlerFunc {

	f := func(w http.ResponseWriter, r *http.Request) {

		// check if context is a request context
		ctxRq, err := kit.MustRequest(r.Context())
		if err != nil {
			m.RespondError(w, err)
			return
		}
		ctx := r.Context()

		// extract token
		token, err := m.ExtractToken(ctx, r)
		if err != nil {
			m.RespondError(w, kitHttp.ErrAuthFailed(ctx))
			return
		}

		// TODO: authorization
		fmt.Println(token)

		// populate context
		r = r.WithContext(ctxRq.ToContext(r.Context()))

		next.ServeHTTP(w, r)
	}

	return f
}

func (m *Middleware) SetContextMiddleware(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {

		// init context
		ctxRq := kit.NewRequestCtx().Rest()

		// set request ID if specified
		requestId := r.Header.Get(HeaderRequestId)
		if requestId != "" {
			ctxRq = ctxRq.WithRequestId(requestId)
		} else {
			ctxRq = ctxRq.WithNewRequestId()
		}

		// set client ip header coming from client
		clientIP := r.Header.Get(HeaderXRealIp)
		// try to get ip from x-forwarder-for
		if clientIP == "" {
			clientIP = r.Header.Get(HeaderXForwarderFor)
		}
		if clientIP != "" {
			ctxRq = ctxRq.WithClientIp(clientIP)
		}

		ctx := ctxRq.ToContext(r.Context())

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}

func (m *Middleware) WithTimeoutMiddleware(next http.HandlerFunc, timeoutSec int) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		timeoutHandler := http.TimeoutHandler(next, time.Duration(timeoutSec)*time.Second, "")
		timeoutHandler.ServeHTTP(w, r)
	}

	return f
}
