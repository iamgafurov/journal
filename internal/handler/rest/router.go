package rest

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func (s *Server) routers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(s.cfg.ServerPrefix+"/docs/", httpSwagger.Handler(
		httpSwagger.URL(s.cfg.ServerPrefix+"./docs/doc.json"),
	))

	mux.HandleFunc(s.cfg.ServerPrefix+"/ping", s.ping)
	mux.HandleFunc(s.cfg.ServerPrefix+"/tokenize", s.tokenize)
	mux.HandleFunc(s.cfg.ServerPrefix+"/untokenize", s.auth(s.tokenDelete))
	mux.HandleFunc(s.cfg.ServerPrefix+"/faculties", s.auth(s.userFaculties))
	mux.HandleFunc(s.cfg.ServerPrefix+"/courses/pt", s.auth(s.groupCoursesPt))
	mux.HandleFunc(s.cfg.ServerPrefix+"/courses/at", s.auth(s.groupCoursesAt))
	mux.HandleFunc(s.cfg.ServerPrefix+"/academic_years", s.auth(s.academicYears))
	mux.HandleFunc(s.cfg.ServerPrefix+"/topic/all", s.auth(s.topics))
	mux.HandleFunc(s.cfg.ServerPrefix+"/topic/create", s.auth(s.topicCreate))
	mux.HandleFunc(s.cfg.ServerPrefix+"/topic/delete", s.auth(s.topicDelete))
	mux.HandleFunc(s.cfg.ServerPrefix+"/topic/update", s.auth(s.topicUpdate))
	mux.HandleFunc(s.cfg.ServerPrefix+"/point_journal/get", s.auth(s.pointsJournal))
	mux.HandleFunc(s.cfg.ServerPrefix+"/point_journal/update", s.auth(s.pointsJournalUpdate))
	mux.HandleFunc(s.cfg.ServerPrefix+"/attendance_journal/get", s.auth(s.getAttendanceJournal))
	mux.HandleFunc(s.cfg.ServerPrefix+"/attendance_journal/update", s.auth(s.updateAttendanceJournal))

	/*
			point/getbycourseId[/weeks/maxPoints]
		    weeks = get empty weeks from up1&upl
		    maxPoint = 100 / emptyWeeks
			students = get from tblVdStKr



			point/update/byCourseId/weekNumber
			attendance/getTopic/bysomeId
			attendance/getJournal/bysomeId
			attendance/addTopic/bysomeId
			attendance/deleteTopic/bysomeId
			attendance/updateJournal/bysomeId
			user/info
	*/
	return mux
}
