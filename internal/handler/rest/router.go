package rest

import "net/http"

func (s *Server) routers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(s.cfg.ServerPrefix+"/ping", s.ping)
	mux.HandleFunc(s.cfg.ServerPrefix+"/tokenize", s.tokenize)
	mux.HandleFunc(s.cfg.ServerPrefix+"/untokenize", s.auth(s.tokenDelete))
	mux.HandleFunc(s.cfg.ServerPrefix+"/faculties", s.auth(s.userFaculties))
	return mux
}
