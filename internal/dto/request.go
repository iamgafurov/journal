package dto

type TokenizeRequest struct {
	ServiceName string    `json:"serviceName"`
	ExternalRef string    `json:"externalRef"`
	LoginPass   LoginPass `json:"loginPass"`
}

type Attendance struct {
}

type LoginPass struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
