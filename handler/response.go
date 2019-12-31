package handler

import "github.com/ScienceQuiz-Backend/utils"

type jwtResponse struct {
	AccessToken string `json:"access_token"`
}

func newJwtResponse(id uint) *jwtResponse {
	r := new(jwtResponse)

	r.AccessToken = utils.GenerateJWT(id)
	return r
}