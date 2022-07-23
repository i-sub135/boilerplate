package model

import "time"

type (
	UsersModel struct {
		ID        int
		UserName  string
		Name      string
		Password  string
		CreatedAt time.Time
		CreatedBy int
		UpdatedAt time.Time
		UpdatedBy int
	}

	UserLoginRequest struct {
		UserName string
		Password string
	}
)
