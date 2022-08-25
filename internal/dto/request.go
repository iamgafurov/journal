package dto

type TokenizeRequest struct {
	ServiceName string    `json:"serviceName"`
	ExternalRef string    `json:"externalRef"`
	LoginPass   LoginPass `json:"loginPass"`
}

type DeleteTokenRequest struct {
	Token       string `json:"token"`
	ServiceName string `json:"serviceName"`
	ExternalRef string `json:"externalRef"`
}

type Attendance struct {
}

type LoginPass struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CheckUserRequest struct {
	Login     string `json:"login"`
	UchprocId int64  `json:"uchprocId"`
}

type ServiceNameExternalRef struct {
	ServiceName string `json:"serviceName"`
	ExternalRef string `json:"externalRef"`
}
