package http

import (
	"github.com/cbr4yan/backend-template/pkg/logger"
	"github.com/cbr4yan/backend-template/pkg/server"
	"golang.org/x/sync/errgroup"
	"net/http"
)

var log = logger.Log

type httpServer struct {
	addr    string
	handler http.Handler
	stop    chan struct{}
}

func New() server.Server {
	return &httpServer{}
}

func (s httpServer) Start() error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.handler,
	}
	g := errgroup.Group{}

	g.Go(func() error {
		<-s.stop
		log.Info().Msg("[http] shutting down server")
		if err := srv.Close(); err != nil {
			log.Error().Msgf("[http] server shutdown failed: %+v", err)
			return err
		}
		return nil
	})
	g.Go(func() error {
		log.Info().Msgf("[http] shutting down server")
		if err := srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info().Msg("[http] server shutdown complete")
			} else {
				log.Error().Msgf("[http] server closed unexpect: %+v", err)
				return err
			}
		}
		return nil
	})
	return g.Wait()
}

func (s httpServer) Stop() error {
	s.stop <- struct{}{}
	return nil
}
