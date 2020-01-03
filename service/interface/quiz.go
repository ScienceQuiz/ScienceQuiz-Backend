package _interface

import "github.com/ScienceQuiz-Backend/model"

type Quiz interface {
	// Quiz 테이블의 새로운 필드 생성 메서드
	Create(*model.Quiz)
	// Question을 기준으로 문제가 존재하는지 확인하는 메서드
	CheckQuizExist(string) bool
}