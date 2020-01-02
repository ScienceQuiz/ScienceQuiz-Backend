package _interface

import "github.com/ScienceQuiz-Backend/model"

type Quiz interface {
	// Quiz 테이블의 새로운 필드 생성 메서드
	Create(*model.Quiz)
}