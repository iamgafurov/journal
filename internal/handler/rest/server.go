package rest

import (
	"context"
	"github.com/iamgafurov/journal/internal/config"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/logger"
	"github.com/iamgafurov/journal/internal/service"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//doc generating command
// swag init -g handlers/rest/server.go

//Server - http api server
// @title API
// @schemes http
// @BasePath /api/v1
type Server struct {
	httpServer *http.Server //http server
	cfg        *config.Config
	service    service.Service
}

//New - new http api server
func New(cfg *config.Config, svc service.Service) *Server {
	return &Server{cfg: cfg, service: svc}
}

//Run http api server
func (s *Server) Run() {

	s.httpServer = &http.Server{
		Addr:              s.cfg.ServerPort,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      40 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       2 * time.Minute,
		Handler:           s.routers(),
	}

	logger.Logger.Info("web api server starting", zap.String("Port", s.cfg.ServerPort), zap.String("", s.cfg.ServerPrefix))
	err := s.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

//Shutdown - graceful shutdown of http api server
func (s *Server) Shutdown(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}

//Ping handler
func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	s.reply(w, dto.Response{Code: enums.Success})
}

//parse body of http Request
func parseBody(r *http.Request, req interface{}) bool {
	if r.Body == nil {
		return false
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Logger.Error("internl/handler/rest, parsebody, readAll", zap.Error(err), zap.Any("Request", req))
		return false
	}
	err = jsoniter.Unmarshal(b, req)
	if err != nil {
		logger.Logger.Error("internl/handler/rest, parsebody, jsoniter.Unmarshal", zap.Error(err), zap.Any("Request", req))
		return false
	}
	return true
}
