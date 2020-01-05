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

type myInformResponse struct {
	Nickname string `json:"nickname"`
}

func newInformResponse(s string) *myInformResponse {
	r := new(myInformResponse)

	r.Nickname = s
	return r
}

type quizListResponse struct {
	Science []quizIndex `json:"science"`
}

func newQuizListResponse(qList []quizIndex) *quizListResponse {
	r := new(quizListResponse)

	r.Science = qList
	return r
}