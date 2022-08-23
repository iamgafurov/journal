package rest

import (
	"context"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/tools"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"time"
)

//http Response writes in json format
func (s *Server) reply(w http.ResponseWriter, r dto.Response) {
	b, err := jsoniter.Marshal(r)
	if err != nil {
		log.Println("error", "app/common.go, reply, jsoniter.Marshal", err)
		http.Error(w, "внутренняя ошибка, попробуйте позже", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(b)
}

//authorization
func (s *Server) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := r.Header.Get("Service")
		if service != enums.ServiceWeb && service != enums.ServiceMobi {
			s.reply(w, dto.Response{Code: enums.Unauthorized})
			return
		}

		token := r.Header.Get("Token")
		if tools.StrEmpty(token) {
			s.reply(w, dto.Response{Code: enums.Unauthorized})
			return
		}

		user, err := s.service.UserGetByToken(r.Context(), token)
		if err != nil {
			s.reply(w, dto.Response{Code: enums.Unauthorized})
			return
		}

		if service != user.Service {
			s.reply(w, dto.Response{Code: enums.Unauthorized})
			return
		}

		if user.ExpireAt.Before(time.Now()) {
			s.reply(w, dto.Response{Code: enums.TokenExpired, Status: enums.StatusFailed, Message: "token expired"})
			return
		}

		if user.Status != enums.StatusActive {
			s.reply(w, dto.Response{Code: enums.UserNotActive, Message: "user not active"})
			return
		}

		resp := s.service.CheckUser(r.Context(), dto.CheckUserRequest{Login: user.Login, UchprocId: user.UchprocId})
		if resp.Code != enums.Success {
			s.reply(w, dto.Response{Code: enums.Unauthorized, Status: enums.StatusFailed, Message: "user deleted"})
			return
		}

		ctx := context.WithValue(r.Context(), "user", &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
