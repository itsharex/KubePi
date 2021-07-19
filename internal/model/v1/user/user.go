package user

import v1 "github.com/KubeOperator/ekko/internal/model/v1"

type User struct {
	v1.BaseModel `storm:"inline"`
	v1.Metadata  `storm:"inline"`
	NickName     string       `json:"nickName" storm:"index"`
	Email        string       `json:"email" storm:"unique"`
	Language     string       `json:"language"`
	IsAdmin      bool         `json:"isAdmin"`
	Authenticate Authenticate `json:"authenticate"`
}

type Authenticate struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

