// Code generated by goctl. DO NOT EDIT.
package types

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}

type LoginRequst struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	JwtToken
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	LabID    string `json:"labID"`
	LabPass  string `json:"labPass"`
}

type RegisterResponse struct {
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	Message   string `json:"message"`
	JwtToken
}