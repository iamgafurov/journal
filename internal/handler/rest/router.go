package rest

import "net/http"

func (s *Server) routers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(s.cfg.ServerPrefix+"/ping", s.ping)
	mux.HandleFunc(s.cfg.ServerPrefix+"/tokenize", s.tokenize)
	return mux
}
