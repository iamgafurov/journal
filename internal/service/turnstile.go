package service

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/iamgafurov/journal/internal/dto"
	"github.com/iamgafurov/journal/internal/enums"
	"github.com/iamgafurov/journal/internal/tools"
	"log"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 60,
	Transport: &http.Transport{
		TLSHandshakeTimeout: 20 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func (s *service) CheckStudents(ctx context.Context, students []dto.CheckAttendanceItem) {
	var (
		resp dto.CheckAttendanceResponse
	)
	httpReq := dto.CheckAttendanceRequest{Students: students}
	code, rawResponse, err := s.post(ctx, s.cfg.TurnstileUrl, httpReq)
	if err != nil {
		s.log.Error("service/turnstile.go SendCheckStudent", zap.Error(err))
		return
	}
	log.Println(code, string(rawResponse), err)

	if code != enums.Success {
		s.log.Error("service/turnstile.go SendCheckStudent s.post", zap.Int("response code not ok, code:", code))
		return
	}

	err = jsoniter.Unmarshal(rawResponse, &resp)
	if err != nil {
		s.log.Error("service/turnstile.go SendCheckStudent jsoniter.Unmarshal", zap.Error(err))
		return
	}

	log.Println(resp)
}

func (s *service) StartTurnstileWorker(ctx context.Context) {
	tc := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-tc.C:
			students, err := s.mssqlDB.GetTurnstileStudents(ctx)
			if err != nil {
				s.log.Error("StartTurnstileWorker, s.mssqlDB.GetTurnstileStudents", zap.Error(err))
				return
			}
			students = append(students, dto.CheckAttendanceItem{RecordBook: "2022060147"})
			log.Println("work")
			if len(students) < 1 {
				continue
			}

			s.CheckStudents(ctx, students)
		case <-ctx.Done():
			s.log.Info("Worker stopped by context")
			return
		}
	}
}

// Post - http client request;
func (s *service) post(ctx context.Context, url string, body any) (code int, data []byte, err error) {
	//handling panics
	defer sentry.Recover()

	var (
		req     *http.Request
		reqBody []byte
	)
	reqBody, err = jsoniter.Marshal(body)
	if err != nil {
		code = enums.InternalError //request failed
		return
	}

	s.log.Info("POST http", zap.String("Request", string(reqBody)))

	// prepare request;
	req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(reqBody))
	if err != nil {
		code = enums.InternalError //request failed
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Token", s.cfg.TurnstileToken)
	xHash := tools.HmacHash(reqBody, s.cfg.TurnstileMasterKey)
	req.Header.Set("X-Hash", xHash)

	// send request;
	resp, err := netClient.Do(req)
	if err != nil {
		code = enums.GatewayTimeout // timeout or network error, need to retry request
		return
	}
	if resp != nil && resp.StatusCode != 200 {
		code = enums.GatewayTimeout
		err = errors.New(resp.Status)
		return
	}
	if resp.Body == nil {
		code = enums.GatewayTimeout
		err = errors.New("empty_http_response")
		return
	}
	defer resp.Body.Close()
	// read response;
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		code = enums.GatewayTimeout
		return
	}

	code = enums.Success
	return
}

func (s *service) UpdateAttendanceJournalTurnstile(ctx context.Context, req dto.UpdateAttendanceJournalRequest) (resp dto.Response) {
	if req.CourseId == 0 {
		resp.ErrCode(enums.BadRequest)
		resp.ErrStr = "null course id"
		return
	}

	if req.UserUchprocCode == 0 {
		resp.ErrCode(enums.Unauthorized)
		resp.ErrStr = "null user code"
		return
	}

	statement, err := s.mssqlDB.GetAttendanceStatement(ctx, req.CourseId)
	if err != nil {
		if err == dto.ErrNoRows {
			resp.ErrCode(enums.BadRequest)
			resp.Message = "course statement not exist"
			resp.ErrStr = resp.Message
			return
		}
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.attendance_journal.go, UpdateAttendanceJournal,  s.mssqlDB.GetAttendanceStatement", zap.Error(err), zap.Any("Request", req))
		return
	}

	topics, err := s.mssqlDB.GetTopics(ctx, req.UserUchprocCode, req.CourseId)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.attendance_journal.go, UpdateAttendanceJournal,  s.mssqlDB.GetTopics", zap.Error(err), zap.Any("Request", req))
		return
	}

	if statement.Kst != req.UserUchprocCode && statement.Kas != req.UserUchprocCode {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "course does not belong to this user"
		return
	}

	if len(req.Attendance) < 1 {
		resp.ErrCode(enums.BadRequest)
		resp.Message = "empty request"
		return
	}

	for i, at := range req.Attendance {
		if at.Id == 0 {
			resp.ErrCode(enums.BadRequest)
			resp.Message = fmt.Sprintf("empty student id, in item:%v", i)
			resp.ErrStr = resp.Message
			return
		}

		if len(at.Attendance) == 0 {
			resp.ErrCode(enums.BadRequest)
			resp.Message = fmt.Sprintf("empty student attendance, in item:%v", i)
			resp.ErrStr = resp.Message
			return
		}

		for i, v := range at.Attendance {
			if v.Number < 1 {
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("invalid topic_number:%v,  for item:%v ", v.Number, i)
				return
			}

			if v.Number > len(topics) {
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("topic with number:%v not exist", v.Number)
				return
			}
			if !topics[v.Number-1].Editable {
				log.Println(topics[v.Number])
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("topic with number:%v not editable", v.Number)
				return
			}

			if strings.TrimSpace(v.Value) != "" && v.Value != "н" {
				resp.ErrCode(enums.BadRequest)
				resp.Message = fmt.Sprintf("invalid attendance value, miss number:%v, value:%s", v.Number, v.Value)
				resp.ErrStr = resp.Message
				return
			}
		}
	}

	err, attErr := s.mssqlDB.UpdateAttendanceJournal(ctx, req.CourseId, req.Attendance)
	if err != nil {
		resp.ErrCode(enums.InternalError)
		resp.ErrStr = err.Error()
		s.log.Error("internal/service.attendance_journal.go, UpdateAttendanceJournal,  s.mssqlDB.UpdateAttendanceJournal", zap.Error(err), zap.Any("Request", req))
		return
	}

	if len(attErr) > 0 {
		resp.ErrCode(enums.SuccessPartially)
		resp.Payload = attErr
		resp.Message = "some records were not saved"
		return
	}

	resp.ErrCode(enums.Success)
	return
}
