package bootstrap

import (
	"context"
	"github.com/mikhailbolshakov/decision"
	domain "github.com/mikhailbolshakov/decision/domain/decision"
	"github.com/mikhailbolshakov/decision/domain/decision/impl"
	"github.com/mikhailbolshakov/decision/http"
	decisionHttp "github.com/mikhailbolshakov/decision/http/decision"
	"github.com/mikhailbolshakov/decision/http/sys"
	"github.com/mikhailbolshakov/decision/kit"
	kitHttp "github.com/mikhailbolshakov/decision/kit/http"
)

// ServiceImpl implements a service bootstrapping
// all dependencies between layers must be specified here
type ServiceImpl struct {
	cfg             *decision.Config
	loadCfgFn       func() (*decision.Config, error)
	http            *kitHttp.Server
	decisionService domain.DecisionService
}

// New creates a new instance of the service
func New() kit.Service {
	s := &ServiceImpl{
		loadCfgFn: decision.LoadConfig,
	}
	s.decisionService = impl.NewDecisionService()
	return s
}

func (s *ServiceImpl) SetConfigLoadFn(fn func() (*decision.Config, error)) {
	s.loadCfgFn = fn
}

func (s *ServiceImpl) GetCode() string {
	return "decision"
}

func (s *ServiceImpl) initHttpServer(ctx context.Context) error {
	// create HTTP server
	s.http = kitHttp.NewHttpServer(s.cfg.Http, decision.LF())

	// create and set middlewares
	mdw := http.NewMiddleware()
	s.http.RootRouter.Use(mdw.SetContextMiddleware)

	// decision routing
	routeBuilder := http.NewRouteBuilder(s.http, mdw)
	routeBuilder.SetRoutes(sys.GetRoutes(sys.NewController()))
	routeBuilder.SetRoutes(decisionHttp.GetRoutes(decisionHttp.NewController(s.decisionService)))

	return routeBuilder.Build()
}

// Init does all initializations
func (s *ServiceImpl) Init(ctx context.Context) error {

	// load config
	var err error
	s.cfg, err = s.loadCfgFn()
	if err != nil {
		return err
	}

	// set log config
	decision.Logger.Init(s.cfg.Log)

	// init http server
	if err := s.initHttpServer(ctx); err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) Start(ctx context.Context) error {

	// listen HTTP connections
	s.http.Listen()

	return nil
}

func (s *ServiceImpl) Close(ctx context.Context) {
	s.http.Close()
}
