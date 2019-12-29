package handler

type userSignUpRequest struct {
	UserId string `json:"userId" validate:"required"`
	UserPw string `json:"userPw" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}
