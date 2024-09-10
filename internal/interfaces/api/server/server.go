package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Run()
	RegisterDefaultCORS()
	RegisterCORS(customConfig cors.Config)
	RegisterRouter(router Router)
	SetShutdownHandler(shutdown func())
}

type server struct {
	port                string
	log                 *logrus.Logger
	ginEngine           *gin.Engine
	connectionCloseChan chan func()
}

type Router interface {
	Register(engine *gin.Engine)
}

func NewServer(port string, log *logrus.Logger) *server {
	gin.SetMode(gin.ReleaseMode)
	ginEngine := gin.New()

	return &server{
		port:                port,
		log:                 log,
		ginEngine:           ginEngine,
		connectionCloseChan: make(chan func(), 1),
	}
}

func (s *server) RegisterDefaultCORS() {
	s.ginEngine.Use(cors.Default())
}

func (s *server) RegisterCORS(customConfig cors.Config) {
	s.ginEngine.Use(cors.New(customConfig))
}

func (s *server) RegisterRouter(router Router) {
	router.Register(s.ginEngine)
}

func (s *server) Run() {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.port),
		Handler: s.ginEngine,
	}

	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		s.log.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			s.log.Fatalf("http server shutdown error: %s", err)
		}
		close(s.connectionCloseChan)
	}()

	s.log.Infof("run http server %s address success", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.Errorf("run http server %s address error: %s", httpServer.Addr, err)
	}

	shutdown := <-s.connectionCloseChan
	if shutdown != nil {
		shutdown()
	}
}

func (s *server) SetShutdownHandler(shutdown func()) {
	s.connectionCloseChan <- shutdown
}
