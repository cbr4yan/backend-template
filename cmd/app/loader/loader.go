package loader

import (
	"context"
	"github.com/cbr4yan/backend-template/config"
	"github.com/cbr4yan/backend-template/pkg/logger"
	"github.com/cbr4yan/backend-template/pkg/server"
	"github.com/cbr4yan/backend-template/pkg/server/http"
	"github.com/cbr4yan/backend-template/pkg/signal"
	"golang.org/x/sync/errgroup"
)

var log = logger.Log

type Application struct {
	name   string
	server server.Server
}

func New(name string) *Application {
	return &Application{
		name: name,
	}
}

func (a *Application) Setup() {
	c, err := config.Setup(a.name)
	if err != nil {
		log.Fatal().Err(err).Msg("[loader] cannot setup config")
		return
	}

	logger.Setup(c)
	router := provideRouter()
	a.server = http.New(c.HttpServer, router)
}

func (a *Application) Run() {
	ctx := signal.WithContext(context.Background())
	g := errgroup.Group{}
	g.Go(func() error {
		return a.server.Start()
	})
	g.Go(func() error {
		<-ctx.Done()
		return a.server.Stop()
	})
	if err := g.Wait(); err != nil {
		log.Fatal().Err(err).Msg("[loader] terminated")
	}
}
