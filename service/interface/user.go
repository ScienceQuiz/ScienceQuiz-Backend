package _interface

import "github.com/ScienceQuiz-Backend/model"

type User interface {
	// 새 사용자 생성 메서드
	Create(*model.User)
	// (userId 기준) 사용자 존재 여부 확인 메서드
	CheckUserExist(string) bool
}
