package kit

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
)

const (
	CallerTypeRest   = "rest"
	CallerTypeTest   = "test"
	CallerTypeJob    = "job"
	CallerTypeQueue  = "queue"
	CallerTypeWs     = "ws"
	CallerTypeWebRtc = "webrtc"

	AppTest = "test"
)

type requestContextKey struct{}

type RequestContext struct {
	// Rid request ID
	Rid string `json:"_ctx.rid,omitempty" mapstructure:"_ctx.rid"`
	// Sid session ID
	Sid string `json:"_ctx.sid,omitempty" mapstructure:"_ctx.sid"`
	// Uid user ID
	Uid string `json:"_ctx.uid,omitempty" mapstructure:"_ctx.uid"`
	// Un username
	Un string `json:"_ctx.un,omitempty" mapstructure:"_ctx.un"`
	// Prj project
	Prj string `json:"_ctx.prj,omitempty" mapstructure:"_ctx.prj"`
	// App application
	App string `json:"_ctx.app,omitempty" mapstructure:"_ctx.app"`
	// Caller who is calling
	Caller string `json:"_ctx.cl,omitempty" mapstructure:"_ctx.cl"`
	// ClId client ID
	ClId string `json:"_ctx.clId,omitempty" mapstructure:"_ctx.clId"`
	// ClIp client IP
	ClIp string `json:"_ctx.clIp,omitempty" mapstructure:"_ctx.clIp"`
	// Roles list of roles
	Roles []string `json:"_ctx.rl,omitempty" mapstructure:"_ctx.rl"`
	// PtId partner ID
	PtId string `json:"_ctx.ptId,omitempty" mapstructure:"_ctx.ptId"`
	// Lang client language
	Lang language.Tag `json:"_ctx.lang,omitempty" mapstructure:"_ctx.lang"`
	// Kv arbitrary key-value
	Kv KV `json:"_ctx.kv,omitempty" mapstructure:"_ctx.kv"`
}

func NewRequestCtx() *RequestContext {
	return &RequestContext{}
}

func (r *RequestContext) GetRequestId() string {
	return r.Rid
}

func (r *RequestContext) GetSessionId() string {
	return r.Sid
}

func (r *RequestContext) GetUserId() string {
	return r.Uid
}

func (r *RequestContext) GetCaller() string {
	return r.Caller
}

func (r *RequestContext) GetClientId() string {
	return r.ClId
}

func (r *RequestContext) GetPartnerId() string {
	return r.PtId
}

func (r *RequestContext) GetRoles() []string {
	return r.Roles
}

func (r *RequestContext) GetUsername() string {
	return r.Un
}

func (r *RequestContext) GetProject() string {
	return r.Prj
}

func (r *RequestContext) GetApp() string {
	return r.App
}

func (r *RequestContext) GetClientIp() string {
	return r.ClIp
}

func (r *RequestContext) GetLang() language.Tag {
	return r.Lang
}

func (r *RequestContext) GetKv() KV {
	return r.Kv
}

func (r *RequestContext) Empty() *RequestContext {
	return &RequestContext{}
}

func (r *RequestContext) WithRequestId(requestId string) *RequestContext {
	r.Rid = requestId
	return r
}

func (r *RequestContext) WithNewRequestId() *RequestContext {
	r.Rid = NewId()
	return r
}

func (r *RequestContext) WithSessionId(sessionId string) *RequestContext {
	r.Sid = sessionId
	return r
}

func (r *RequestContext) WithProject(project string) *RequestContext {
	r.Prj = project
	return r
}

func (r *RequestContext) WithApp(app string) *RequestContext {
	r.App = app
	return r
}

func (r *RequestContext) WithClientIp(ip string) *RequestContext {
	r.ClIp = ip
	return r
}

func (r *RequestContext) WithLang(lang language.Tag) *RequestContext {
	r.Lang = lang
	return r
}

func (r *RequestContext) WithKv(key string, val interface{}) *RequestContext {
	if r.Kv == nil {
		r.Kv = KV{}
	}
	r.Kv[key] = val
	return r
}

func (r *RequestContext) Rest() *RequestContext {
	return r.WithCaller(CallerTypeRest)
}

func (r *RequestContext) Webrtc() *RequestContext {
	return r.WithCaller(CallerTypeWebRtc)
}

func (r *RequestContext) Test() *RequestContext {
	return r.WithCaller(CallerTypeTest)
}

func (r *RequestContext) Job() *RequestContext {
	return r.WithCaller(CallerTypeJob)
}

func (r *RequestContext) Queue() *RequestContext {
	return r.WithCaller(CallerTypeQueue)
}

func (r *RequestContext) Ws() *RequestContext {
	return r.WithCaller(CallerTypeWs)
}

func (r *RequestContext) TestApp() *RequestContext {
	return r.WithApp(AppTest)
}

func (r *RequestContext) EN() *RequestContext {
	return r.WithLang(language.English)
}

func (r *RequestContext) WithCaller(caller string) *RequestContext {
	r.Caller = caller
	return r
}

func (r *RequestContext) WithUser(userId, username string) *RequestContext {
	r.Uid = userId
	r.Un = username
	return r
}

func (r *RequestContext) WithPartner(partnerId string) *RequestContext {
	r.PtId = partnerId
	return r
}

func (r *RequestContext) WithClient(clientId string) *RequestContext {
	r.ClId = clientId
	return r
}

func (r *RequestContext) WithRoles(roles ...string) *RequestContext {
	r.Roles = roles
	return r
}

func (r *RequestContext) ToContext(parent context.Context) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	return context.WithValue(parent, requestContextKey{}, r)
}

func (r *RequestContext) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"_ctx.rid":  r.Rid,
		"_ctx.sid":  r.Sid,
		"_ctx.uid":  r.Uid,
		"_ctx.un":   r.Un,
		"_ctx.prj":  r.Prj,
		"_ctx.app":  r.App,
		"_ctx.cl":   r.Caller,
		"_ctx.clId": r.ClId,
		"_ctx.ptId": r.PtId,
		"_ctx.clIp": r.ClIp,
		"_ctx.rl":   r.Roles,
		"_ctx.lang": r.Lang,
		"_ctx.kv":   r.Kv,
	}
}

func Request(context context.Context) (*RequestContext, bool) {
	if r, ok := context.Value(requestContextKey{}).(*RequestContext); ok {
		return r, true
	}
	return &RequestContext{}, false
}

func MustRequest(context context.Context) (*RequestContext, error) {
	if r, ok := context.Value(requestContextKey{}).(*RequestContext); ok {
		return r, nil
	}
	return &RequestContext{}, errors.New("context is invalid")
}

func ContextToGrpcMD(ctx context.Context) (metadata.MD, bool) {
	if r, ok := Request(ctx); ok {
		rm, _ := json.Marshal(*r)
		return metadata.Pairs("rq-bin", string(rm)), true
	}
	return metadata.Pairs(), false
}

func FromGrpcMD(ctx context.Context, md metadata.MD) context.Context {
	if rqb, ok := md["rq-bin"]; ok {
		if len(rqb) > 0 {
			rm := []byte(rqb[0])
			rq := &RequestContext{}
			_ = json.Unmarshal(rm, rq)
			return context.WithValue(ctx, requestContextKey{}, rq)
		}
	}
	return ctx
}

func FromMap(ctx context.Context, mp map[string]interface{}) (context.Context, error) {
	var r *RequestContext
	err := mapstructure.Decode(mp, &r)
	if err != nil {
		return nil, err
	}
	return r.ToContext(ctx), nil
}

func Copy(ctx context.Context) context.Context {
	if r, ok := Request(ctx); ok {
		ct, err := FromMap(context.TODO(), r.ToMap())
		if err != nil {
			return ctx
		}
		return ct
	}
	return ctx
}
